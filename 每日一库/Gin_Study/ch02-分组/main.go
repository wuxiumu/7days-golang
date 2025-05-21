package ch02_分组

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// 带参数的路由
func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.GET("/user/:name/:action", func(c *gin.Context) {
		name := c.Param("name")
		action := c.Param("action")
		message := name + " is " + action
		c.JSON(http.StatusOK, gin.H{
			"message": message,
		})
	})

	r.GET("/user/:name/*action", func(c *gin.Context) {
		name := c.Param("name")
		action := c.Param("action")
		message := name + " is " + action
		c.JSON(http.StatusOK, gin.H{
			"message": message,
		})
	})

	r.Run()
}
