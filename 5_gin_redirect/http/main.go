package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	//创建路由引擎
	router := gin.Default()
	//绑定路由规则
	router.GET("/test", func(c *gin.Context) {
		//重定向
		c.Redirect(http.StatusMovedPermanently, "http://www.baidu.com")
	})
	//启动服务
	router.Run(":8080")
}
