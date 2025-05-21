# 01一个简单的 Go Web 程序

## 1. 新建项目

```bash
mkdir 01-simple-web-app
cd 01-simple-web-app
```

## 2. 初始化项目

```bash
go mod init github.com/yourusername/01-simple-web-app
```

## 3. 编写代码

```go
package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, world!")
}
```

## 4. 运行程序

```bash
go run main.go
```

## 5. 浏览器访问

打开浏览器，访问 `http://localhost:8080` 即可看到输出的 `Hello, world!` 字样。    