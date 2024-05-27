package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	//创建路由引擎
	router := gin.Default()
	//绑定路由规则
	router.GET("test", func(c *gin.Context) {
		//指定重定向的URL
		c.Request.URL.Path = "/test2"
		//重写一个引擎
		router.HandleContext(c)
	})
	router.GET("/test2", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"hello": "world",
		})
	})

	//启动服务
	router.Run(":8080")
}
