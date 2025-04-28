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
