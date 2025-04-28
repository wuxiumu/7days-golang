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
