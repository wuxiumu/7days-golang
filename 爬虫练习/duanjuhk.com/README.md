# 爬虫练习

## 目标

爬取 [duanjuhk.com](https://www.duanjuhk.com/) 网站的新闻内容，并保存到本地文件。

## 步骤   

1. go get github.com/PuerkitoBio/goquery   
2. 编写爬虫代码，爬取 [duanjuhk.com](https://www.duanjuhk.com/) 网站的新闻内容，并保存到本地文件。
3. 运行爬虫代码，保存到本地文件。
4. 分析爬取到的内容，提取有价值的信息。
5. 编写报表程序，生成本地数据。

### 1.文件：baidu.go
> 这是爬取网站的新闻内容并保存到本地文件的代码。
> 
> https://duanjuhk.com/vodtype/jingxuanduanju-{i}.html
> 
> html 内容存入 cache/page/{i}.html 文件中


```go
package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	baseURL := "https://duanjuhk.com/vodtype/jingxuanduanju-"
	for i := 1; i <= 577; i++ {
		// 构建当前页的URL
		url := baseURL + strconv.Itoa(i) + ".html"

		// 发送HTTP GET请求
		resp, err := http.Get(url)
		if err != nil {
			log.Printf("请求 %s 出错: %v\n", url, err)
			continue
		}
		defer resp.Body.Close()

		// 使用goquery解析HTML文档
		if resp.StatusCode != http.StatusOK {
			fmt.Printf("请求 %s 时状态码错误: %d %s\n", url, resp.StatusCode, resp.Status)
			continue
		}
		doc, err := goquery.NewDocumentFromReader(resp.Body)
		if err != nil {
			log.Printf("解析 %s 的HTML文档出错: %v\n", url, err)
			continue
		}

		// 获取 HTML 代码
		html, err := doc.Html()
		if err != nil {
			fmt.Printf("获取 %s 的HTML代码出错: %v\n", url, err)
			continue
		}

		// 创建 cache/page 目录（如果不存在）
		err = os.MkdirAll("cache/page", os.ModePerm)
		if err != nil {
			fmt.Printf("创建目录 cache/page 出错: %v\n", err)
			continue
		}

		// html 内容存入 cache/page/{i}.html 文件中
		filePath := fmt.Sprintf("cache/page/%d.html", i)
		f, err := os.Create(filePath)
		if err != nil {
			fmt.Printf("创建文件 %s 出错: %v\n", filePath, err)
			continue
		}
		defer f.Close()
		_, err = f.WriteString(html)
		if err != nil {
			fmt.Printf("写入文件 %s 出错: %v\n", filePath, err)
			continue
		}
		fmt.Printf("网页 %s 的内容已保存到 %s 文件中\n", url, filePath)
	}
}

```

### 2. 程序sp.go

> 解析 ./cache/page/{i}.html 文件
> 
> 获取标题，链接，内容，图片链接，评分等信息
>
> 存入 ./cache/data/{i}.json 文件中
>
```go
package main

import (
	"encoding/json"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"os"
	"strconv"
	"strings"
	"sync"
)

// 常量定义
const (
	baseCachePageURL = "./cache/page/"
	baseCacheDataURL = "./cache/data/"
)

// 处理单个页面的函数
func processPage(page int, wg *sync.WaitGroup) error {
	defer wg.Done()
	localPath := baseCachePageURL + strconv.Itoa(page) + ".html"

	// 获取 localPath 文件内容
	content, err := os.ReadFile(localPath)
	if err != nil {
		return fmt.Errorf("读取文件 %s 失败: %w", localPath, err)
	}

	// 解析 HTML 内容
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(string(content)))
	if err != nil {
		return fmt.Errorf("解析 HTML 文件 %s 失败: %w", localPath, err)
	}

	// 解析列表标题，id，链接，内容,图片链接
	list := parseList(doc)

	// list 转成 json 存入文件 ./cache/data/1.json
	localJsonPath := baseCacheDataURL + strconv.Itoa(page) + ".json"
	fmt.Println("正在处理第", page, "页")
	err = saveListToJSON(list, localJsonPath)
	if err != nil {
		return fmt.Errorf("保存 JSON 文件 %s 失败: %w", localJsonPath, err)
	}
	return nil
}

// 解析 HTML 文档中的列表信息
func parseList(doc *goquery.Document) []map[string]string {
	var list []map[string]string
	doc.Find(".hl-list-item.hl-col-xs-4.hl-col-sm-3.hl-col-md-20w.hl-col-lg-2").Each(func(i int, selection *goquery.Selection) {
		title := selection.Find(".hl-item-thumb.hl-lazy").AttrOr("title", "")       // 标题
		link := selection.Find(".hl-item-thumb.hl-lazy").AttrOr("href", "")         // 链接
		img := selection.Find(".hl-item-thumb.hl-lazy").AttrOr("data-original", "") // 图片链接
		content := selection.Find(".hl-lc-1.remarks").Text()                        // 集合内容
		score := selection.Find(".hl-item-sub.hl-text-muted.hl-lc-1").Text()        // 评分

		if title != "" && link != "" && img != "" {
			// 去掉评分中的空格 和 '/'
			score = strings.Replace(score, " ", "", -1)
			score = strings.Replace(score, "/", "", -1)
			// 组装 map
			item := make(map[string]string)
			item["title"] = title
			item["link"] = link
			item["img"] = img
			item["content"] = content
			item["score"] = score
			list = append(list, item)
		}
	})
	return list
}

// 将列表信息保存为 JSON 文件
func saveListToJSON(list []map[string]string, filePath string) error {
	jsonStr, err := json.Marshal(list)
	if err != nil {
		return fmt.Errorf("将列表转换为 JSON 失败: %w", err)
	}

	jsonFile, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("创建文件 %s 失败: %w", filePath, err)
	}
	defer jsonFile.Close()

	_, err = jsonFile.Write(jsonStr)
	if err != nil {
		return fmt.Errorf("写入文件 %s 失败: %w", filePath, err)
	}
	return nil
}

func main() {
	var wg sync.WaitGroup
	totalPages := 577
	var errors []error

	// 为每一页启动一个 goroutine
	for page := 10; page <= totalPages; page++ {
		wg.Add(1)
		go func(p int) {
			err := processPage(p, &wg)
			if err != nil {
				errors = append(errors, err)
			}
		}(page)
	}

	// 等待所有 goroutine 完成
	wg.Wait()

	// 处理错误
	if len(errors) > 0 {
		fmt.Println("处理过程中出现以下错误:")
		for _, err := range errors {
			fmt.Println(err)
		}
	} else {
		fmt.Println("所有页面处理完成")
	}
}

```

### 3. 程序 water.go

> ./cache/data/{i}.json 文件解析出图片
> 
> 下载图片并保存到本地
> 
> 保存到 ./cache/water/ 目录下
> 

```go
package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"sync"
)

// 常量定义
const (
	baseDomainrUrl   = "https://duanjuhk.com"
	baseCacheDataURL = "./cache/data/"
	baseWaterPageUrl = "./cache/water"
	// 并发数量限制
	maxConcurrency = 5
)

// 从 JSON 文件中读取数据并转换为数组
func readJSONFile(filePath string) ([]map[string]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open JSON file: %w", err)
	}
	defer file.Close()

	var data []map[string]string
	err = json.NewDecoder(file).Decode(&data)
	if err != nil {
		return nil, fmt.Errorf("failed to decode JSON file: %w", err)
	}
	return data, nil
}

// 下载图片并保存到本地
func downloadImage(url string, savePath string) error {
	resp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("failed to download image from %s: %w", url, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to download image from %s: status code %d", url, resp.StatusCode)
	}

	out, err := os.Create(savePath)
	if err != nil {
		return fmt.Errorf("failed to create file %s: %w", savePath, err)
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return fmt.Errorf("failed to copy image data to %s: %w", savePath, err)
	}
	return nil
}

// 创建保存图片的目录
func createImageDirectory(dir string) error {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err := os.MkdirAll(dir, 0755)
		if err != nil {
			return fmt.Errorf("failed to create image directory %s: %w", dir, err)
		}
	}
	return nil
}

// 处理单个页面的函数
func processPage(page int) error {
	jsonFilePath := baseCacheDataURL + strconv.Itoa(page) + ".json"

	// 读取 JSON 文件内容
	data, err := readJSONFile(jsonFilePath)
	if err != nil {
		return err
	}

	// 创建保存图片的目录
	if err := createImageDirectory(baseWaterPageUrl); err != nil {
		return err
	}

	// 遍历数组，获取图片地址并下载图片
	for i, item := range data {
		imgURL := baseDomainrUrl + item["img"]
		if imgURL != "" {
			// 生成保存图片的文件名
			imagename := fmt.Sprintf("img_%d-%d-%d.jpg", page, i+1, i+1)
			savePath := fmt.Sprintf("%s/%s", baseWaterPageUrl, imagename)
			log.Printf("Downloading image %d from %s to %s\n", i+1, imgURL, savePath)
			err := downloadImage(imgURL, savePath)
			if err != nil {
				log.Printf("Failed to download image %d: %v\n", i+1, err)
			}
		}
	}

	log.Printf("Finished processing page %d", page)
	return nil
}

// 处理所有页面
func processAllPages(totalPages int) []error {
	var wg sync.WaitGroup
	var errors []error
	sem := make(chan struct{}, maxConcurrency)

	// 为每一页启动一个 goroutine
	for page := 1; page <= totalPages; page++ {
		wg.Add(1)
		sem <- struct{}{} // 获取信号量
		go func(p int) {
			defer wg.Done()
			defer func() { <-sem }() // 释放信号量
			err := processPage(p)
			if err != nil {
				errors = append(errors, err)
			}
		}(page)
	}

	// 等待所有 goroutine 完成
	wg.Wait()
	return errors
}

func main() {
	totalPages := 577 // 总共 577 页，这里只处理前 2 页

	// 处理所有页面
	errors := processAllPages(totalPages)

	// 处理错误
	if len(errors) > 0 {
		log.Println("以下错误发生在处理过程中:")
		for _, err := range errors {
			log.Println(err)
		}
	} else {
		log.Println("所有页面处理完成")
	}
}

```