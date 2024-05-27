package main

import "github.com/gin-gonic/gin"

func main() {
	//创建路由引擎
	router := gin.Default()
	//初始化路由组
	v1 := router.Group("/v1")
	{
		//绑定路由规则
		v1.GET("/test1", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "this is test1!",
			})
		})
		v1.GET("/test2", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "this is test2!",
			})
		})
		//嵌套路由组
		v2 := v1.Group("v2")
		{
			v2.GET("testv2", func(c *gin.Context) {
				c.JSON(200, gin.H{
					"message": "this is testv2!",
				})
			})
		}
	}
	//启动服务
	router.Run(":8080")
}
