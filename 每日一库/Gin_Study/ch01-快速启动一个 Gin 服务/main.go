package main

import "github.com/gin-gonic/gin"

// 快速启动一个 Gin 服务
func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run(:8080) // 监听并在 0.0.0.0:8080 上启动服务
}
