package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {
	//创建一个路由引擎
	router := gin.Default()

	//绑定路由规则
	//get：请求方式，/hello请求路径
	//客户端以GET方法请求/hello函数时，会执行后面的匿名函数handlefunc
	router.GET("/hello", func(c *gin.Context) {
		//c.Json：返回JSON格式的数据
		c.JSON(http.StatusOK, gin.H{
			"message": "hello world!",
		})
	})

	//启动HTTP服务监听端口，默认8080
	err := router.Run("8080")
	if err != nil {
		log.Fatalf("router run failed: %v", err)
	}
}
