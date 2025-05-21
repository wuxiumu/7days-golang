package main

import (
	"github.com/gin-gonic/gin"
)

type Person struct {
	Id   string `uri:"id" binding:"required,uuid"`
	Name string `uri:"name" binding:"required"`
}

// GroupRouter 路由组
func main() {
	router := gin.Default()
	router.GET("/:name/:id", func(c *gin.Context) {
		var person Person
		if err := c.ShouldBindUri(&person); err == nil {
			c.Json(400, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, gin.H{"name": person.Name, "id": person.Id})

	})
	router.Run(":8080")
}
