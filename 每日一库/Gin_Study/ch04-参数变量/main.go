package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.GET("/good/:id", goodDetail)
	r.Run(":8080")
}

func goodDetail(c *gin.Context) {
	id := c.Param("id")
	name := c.Query("name")
	age := c.DefaultQuery("age", "18")
	gender := c.PostForm("gender")
	c.JSON(http.StatusOK, gin.H{
		"id":     id,
		"name":   name,
		"age":    age,
		"gender": gender,
	})
}
