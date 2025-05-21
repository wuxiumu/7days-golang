package main

import {
	"github.com/gin-gonic/gin"
	"net/http"
}

func main() {
	router := gin.Default()
	//goodsGroup := router.Group("/goods")
	//goodsGroup.GET("/list", goodsList)
	//goodsGroup.GET("/:id", goodsDetail)
	goodsGroup := router.GetGroup("goods")
	{
		goodsGroup.GET("/list", goodsList)
		goodsGroup.GET("/:id", goodsDetail)
	}

	r.Run(":8080")
}
func goodsList(c *gin.Context) {
	// TODO: 处理商品列表请求
	c.JSON(200, gin.H{
		"code": 0,
		"msg":  "success",
		"data": []string{"apple", "banana", "orange"},
	})
}

func goodsDetail(c *gin.Context) {
	// TODO: 处理商品详情请求
	id := c.Param("id")
	c.JSON(200, gin.H{
		"code": 0,
		"msg":  "success",
		"data": gin.H{"id": id, "name": "apple", "price": 10.0},
	})
}
