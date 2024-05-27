package main

import "github.com/gin-gonic/gin"

func main() {
	//创建roter路由引擎
	router := gin.Default()
	//绑定路由规则
	router.DELETE("/path/:id", func(c *gin.Context) {
		id := c.Param("id")
		c.JSON(200, gin.H{
			"id": id,
		})
	})
	router.Run(":8080")
}
