package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func middlewOne() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("我是middlewOne：方法调用前！")
		//调用下一个中间件或者路由方法
		c.Next()
		fmt.Println("我是middlewOne：方法调用后！")
	}
}

func middlewTwo() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("我是middlewTwo：方法调用前！")
		//调用下一个中间件或者路由方法
		c.Next()
		fmt.Println("我是middlewTwo：方法调用后！")
	}
}

func middlewThree() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("我是middlewThree：方法调用前！")
		//调用下一个中间件或者路由方法
		c.Next()
		fmt.Println("我是middlewThree：方法调用后！")
	}
}

func main() {
	//创建一个路由引擎
	router := gin.Default()
	//绑定路由规则，使用多个中间件
	router.GET("/test", middlewOne(), middlewTwo(), middlewThree(), func(c *gin.Context) {
		fmt.Println("我是路由test的方法！")
		c.JSON(200, gin.H{
			"message": "success",
		})
	})
	router.Run(":8080")
}
