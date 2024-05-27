package main

import "github.com/gin-gonic/gin"

func main() {
	//创建一个路由引擎
	router := gin.Default()

	//绑定路由规则
	router.GET("/path/:id", func(c *gin.Context) {
		//返回URL的参数值
		id := c.Param("id")
		//查询对应的参数值，如果存在返回对应的值
		//否则返回对应的空字符串
		user := c.Query("user")
		pwd := c.Query("pwd")
		if user == "" {
			user = c.DefaultQuery("user", "hahaha")
		}
		//返回JSON格式的数据
		c.JSON(200, gin.H{
			"success": true,
			"id":      id,
			"user":    user,
			"pwd":     pwd,
		})
	})

	//启动服务
	router.Run(":8080")
}
