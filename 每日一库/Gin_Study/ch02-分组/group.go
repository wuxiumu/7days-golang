package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()

	// Group 1
	v1 := r.Group("/api/v1")
	{
		v1.GET("/users", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "get users",
			})
		})
		v1.POST("/users", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "create user",
			})
		})
	}

	// Group 2
	v2 := r.Group("/api/v2")
	{
		v2.GET("/users", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "get users",
			})
		})
		v2.POST("/users", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "create user",
			})
		})
	}

	r.Run(":8080")
}
