package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

// 全局中间件
func middle() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("全局中间件：方法调用前！")
		//调用方法
		c.Next()
		fmt.Println("全局中间件：方法调用后！")
	}
}

// test3的中间件
func middle3() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("test3中间件：方法调用前！")
		//调用方法
		c.Next()
		fmt.Println("test3中间件：方法调用后！")
	}
}

// v1路由组的中间件
func middlev1() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("v1路由组中间件：方法调用前！")
		//调用方法
		c.Next()
		fmt.Println("v1路由组中间件：方法调用后！")
	}
}

func main() {
	//创建一个新的路由引擎
	router := gin.Default()
	//注册一个全局中间件
	router.Use(middle())
	//绑定路由规则
	router.GET("/test1", func(c *gin.Context) {
		fmt.Println("我是路由test1的方法！")
		c.JSON(200, gin.H{
			"message": "success",
		})
	})

	router.GET("/test2", func(c *gin.Context) {
		fmt.Println("我是路由test2的方法！")
		c.JSON(200, gin.H{
			"message": "success",
		})
	})

	router.GET("/test3", middle3(), func(c *gin.Context) {
		fmt.Println("我是路由test3的方法！")
		c.JSON(200, gin.H{
			"message": "success",
		})
	})

	//定义一个路由组，并注册路由组中间件
	v1 := router.Group("v1").Use(middlev1())
	{
		v1.GET("/test1", func(c *gin.Context) {
			fmt.Println("我是路由组v1中test1的方法！")
			c.JSON(200, gin.H{
				"message": "success",
			})
		})
		v1.GET("/test2", func(c *gin.Context) {
			fmt.Println("我是路由组v1中test2的方法！")
			c.JSON(200, gin.H{
				"message": "success",
			})
		})
	}

	//启动服务
	router.Run(":8080")
}
