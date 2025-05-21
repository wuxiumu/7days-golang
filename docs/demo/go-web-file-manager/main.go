package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"
)

// 文件信息结构体
type FileInfo struct {
	Name        string    `json:"name"`
	Path        string    `json:"path"`
	Size        int64     `json:"size"`
	IsDir       bool      `json:"isDir"`
	ModTime     time.Time `json:"modTime"`
	Permissions string    `json:"permissions"`
}

// 配置信息
type Config struct {
	RootPath string `json:"rootPath"`
	Port     int    `json:"port"`
}

var config Config

func main() {
	// 加载配置
	err := loadConfig("config.json")
	if err != nil {
		log.Fatalf("加载配置失败: %v", err)
	}

	// 确保根目录存在
	if _, err := os.Stat(config.RootPath); os.IsNotExist(err) {
		if err := os.MkdirAll(config.RootPath, 0755); err != nil {
			log.Fatalf("创建根目录失败: %v", err)
		}
	}

	// 注册路由
	http.HandleFunc("/api/list/", listFiles)
	http.HandleFunc("/api/mkdir/", createDirectory)
	http.HandleFunc("/api/delete/", deleteFile)
	http.HandleFunc("/api/upload/", uploadFile)
	http.HandleFunc("/download/", downloadFile)
	http.HandleFunc("/view/", viewFile)
	http.HandleFunc("/", serveStatic)

	// 启动服务器
	address := fmt.Sprintf(":%d", config.Port)
	log.Printf("服务器启动在 %s，根目录: %s", address, config.RootPath)
	log.Fatal(http.ListenAndServe(address, nil))
}

// 加载配置文件
func loadConfig(path string) error {
	file, err := os.Open(path)
	if err != nil {
		// 使用默认配置
		config = Config{
			RootPath: "files",
			Port:     8080,
		}
		return nil
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	return decoder.Decode(&config)
}

// 处理文件列表请求
func listFiles(w http.ResponseWriter, r *http.Request) {
	relativePath := strings.TrimPrefix(r.URL.Path, "/api/list")
	absolutePath := filepath.Join(config.RootPath, filepath.Clean(relativePath))

	// 安全检查
	if !strings.HasPrefix(absolutePath, config.RootPath) {
		http.Error(w, "访问被拒绝", http.StatusForbidden)
		return
	}

	// 检查目录是否存在
	stat, err := os.Stat(absolutePath)
	if err != nil {
		http.Error(w, "目录不存在", http.StatusNotFound)
		return
	}

	// 确保是目录
	if !stat.IsDir() {
		http.Error(w, "不是目录", http.StatusBadRequest)
		return
	}

	// 读取目录内容
	files, err := os.ReadDir(absolutePath)
	if err != nil {
		http.Error(w, "无法读取目录", http.StatusInternalServerError)
		return
	}

	// 准备文件信息列表
	fileInfoList := make([]FileInfo, 0, len(files))
	for _, file := range files {
		info, err := file.Info()
		if err != nil {
			continue
		}

		fileInfo := FileInfo{
			Name:        info.Name(),
			Path:        filepath.Join(relativePath, info.Name()),
			Size:        info.Size(),
			IsDir:       info.IsDir(),
			ModTime:     info.ModTime(),
			Permissions: info.Mode().String(),
		}
		fileInfoList = append(fileInfoList, fileInfo)
	}

	// 返回JSON响应
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(fileInfoList)
}

// 创建目录
func createDirectory(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "方法不允许", http.StatusMethodNotAllowed)
		return
	}

	relativePath := strings.TrimPrefix(r.URL.Path, "/api/mkdir")
	absolutePath := filepath.Join(config.RootPath, filepath.Clean(relativePath))

	// 安全检查
	if !strings.HasPrefix(absolutePath, config.RootPath) {
		http.Error(w, "访问被拒绝", http.StatusForbidden)
		return
	}

	// 创建目录
	err := os.MkdirAll(absolutePath, 0755)
	if err != nil {
		http.Error(w, fmt.Sprintf("创建目录失败: %v", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "目录创建成功: %s", relativePath)
}

// 删除文件或目录
func deleteFile(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "方法不允许", http.StatusMethodNotAllowed)
		return
	}

	relativePath := strings.TrimPrefix(r.URL.Path, "/api/delete")
	absolutePath := filepath.Join(config.RootPath, filepath.Clean(relativePath))

	// 安全检查
	if !strings.HasPrefix(absolutePath, config.RootPath) {
		http.Error(w, "访问被拒绝", http.StatusForbidden)
		return
	}

	// 删除文件或目录
	err := os.RemoveAll(absolutePath)
	if err != nil {
		http.Error(w, fmt.Sprintf("删除失败: %v", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "删除成功: %s", relativePath)
}

// 上传文件
func uploadFile(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "方法不允许", http.StatusMethodNotAllowed)
		return
	}

	// 解析表单数据
	err := r.ParseMultipartForm(32 << 20) // 32MB
	if err != nil {
		http.Error(w, fmt.Sprintf("解析表单失败: %v", err), http.StatusBadRequest)
		return
	}

	// 获取文件
	file, handler, err := r.FormFile("file")
	if err != nil {
		http.Error(w, fmt.Sprintf("获取文件失败: %v", err), http.StatusBadRequest)
		return
	}
	defer file.Close()

	// 获取目标路径
	relativePath := r.FormValue("path")
	if relativePath == "" {
		relativePath = "/"
	}
	absoluteDir := filepath.Join(config.RootPath, filepath.Clean(relativePath))

	// 安全检查
	if !strings.HasPrefix(absoluteDir, config.RootPath) {
		http.Error(w, "访问被拒绝", http.StatusForbidden)
		return
	}

	// 确保目录存在
	if _, err := os.Stat(absoluteDir); os.IsNotExist(err) {
		if err := os.MkdirAll(absoluteDir, 0755); err != nil {
			http.Error(w, fmt.Sprintf("创建目录失败: %v", err), http.StatusInternalServerError)
			return
		}
	}

	// 创建文件
	absolutePath := filepath.Join(absoluteDir, handler.Filename)
	dst, err := os.Create(absolutePath)
	if err != nil {
		http.Error(w, fmt.Sprintf("创建文件失败: %v", err), http.StatusInternalServerError)
		return
	}
	defer dst.Close()

	// 复制文件内容
	if _, err := io.Copy(dst, file); err != nil {
		http.Error(w, fmt.Sprintf("保存文件失败: %v", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "文件上传成功: %s", filepath.Join(relativePath, handler.Filename))
}

// 下载文件
func downloadFile(w http.ResponseWriter, r *http.Request) {
	relativePath := strings.TrimPrefix(r.URL.Path, "/download")
	absolutePath := filepath.Join(config.RootPath, filepath.Clean(relativePath))

	// 安全检查
	if !strings.HasPrefix(absolutePath, config.RootPath) {
		http.Error(w, "访问被拒绝", http.StatusForbidden)
		return
	}

	// 检查文件是否存在
	if _, err := os.Stat(absolutePath); os.IsNotExist(err) {
		http.Error(w, "文件不存在", http.StatusNotFound)
		return
	}

	// 下载文件
	w.Header().Set("Content-Disposition", "attachment; filename="+filepath.Base(absolutePath))
	http.ServeFile(w, r, absolutePath)
}

// 查看文件内容
func viewFile(w http.ResponseWriter, r *http.Request) {
	relativePath := strings.TrimPrefix(r.URL.Path, "/view")
	absolutePath := filepath.Join(config.RootPath, filepath.Clean(relativePath))

	// 安全检查
	if !strings.HasPrefix(absolutePath, config.RootPath) {
		http.Error(w, "访问被拒绝", http.StatusForbidden)
		return
	}

	// 检查文件是否存在
	stat, err := os.Stat(absolutePath)
	if err != nil {
		http.Error(w, "文件不存在", http.StatusNotFound)
		return
	}

	// 确保不是目录
	if stat.IsDir() {
		http.Error(w, "无法查看目录内容", http.StatusBadRequest)
		return
	}

	// 限制文件大小
	if stat.Size() > 1024*1024 { // 1MB
		http.Error(w, "文件太大，无法在线查看", http.StatusBadRequest)
		return
	}

	// 读取文件内容
	content, err := os.ReadFile(absolutePath)
	if err != nil {
		http.Error(w, "无法读取文件内容", http.StatusInternalServerError)
		return
	}

	// 设置内容类型
	contentType := getContentType(absolutePath)
	w.Header().Set("Content-Type", contentType)
	w.Write(content)
}

// 获取文件内容类型
func getContentType(path string) string {
	ext := strings.ToLower(filepath.Ext(path))
	switch ext {
	case ".txt", ".md", ".log", ".json", ".xml", ".html", ".css", ".js":
		return "text/plain; charset=utf-8"
	case ".jpg", ".jpeg":
		return "image/jpeg"
	case ".png":
		return "image/png"
	case ".gif":
		return "image/gif"
	case ".pdf":
		return "application/pdf"
	default:
		return "application/octet-stream"
	}
}

// 提供静态文件
func serveStatic(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./static/index.html")
}