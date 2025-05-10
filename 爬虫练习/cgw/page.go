package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"sync"

	"github.com/PuerkitoBio/goquery"
)

const (
	cacheDir = "./cache/page_temp/"
)

// 处理单个页面的函数
func processPage(page int, wg *sync.WaitGroup) error {
	defer wg.Done()

	// 构建当前页的URL
	baseURL := "https://because.gezawbla.xyz/archives/page/"
	url := baseURL + strconv.Itoa(page)
	i := page - 1

	// 发送HTTP GET请求
	resp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("请求 %s 出错: %v\n", url, err)
	}
	defer resp.Body.Close()

	// 使用goquery解析HTML文档
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return fmt.Errorf("解析 %s 的HTML文档出错: %v\n", url, err)
	}

	// 获取 HTML 代码
	html, err := doc.Html()
	if err != nil {
		return fmt.Errorf("获取 %s 的HTML代码出错: %v\n", url, err)
	}

	// 创建 cacheDir 目录（如果不存在）
	err = os.MkdirAll(cacheDir, os.ModePerm)
	if err != nil {
		return fmt.Errorf("创建目录 %s 出错: %v\n", cacheDir, err)
	}

	// html 内容存入 cacheDir/{i}.html 文件中
	filePath := fmt.Sprintf(cacheDir+"%d.html", i)
	f, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("创建文件 %s 出错: %v\n", filePath, err)
	}
	defer f.Close()
	_, err = f.WriteString(html)
	if err != nil {
		return fmt.Errorf("写入文件 %s 出错: %v\n", filePath, err)
	}
	fmt.Println("正在处理第", page, "页")
	return nil
}
func main() {
	var wg sync.WaitGroup
	totalPages := 138
	var errors []error

	// 为每一页启动一个 goroutine
	for page := 1; page <= totalPages; page++ {
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
