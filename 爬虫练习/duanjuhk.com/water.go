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
