package main

import "github.com/gin-gonic/gin"

func main() {
	//创建路由引擎
	router := gin.Default()
	//绑定路由规则
	router.PUT("/path", func(c *gin.Context) {
		user := c.DefaultPostForm("user", "xuji")
		pwd := c.PostForm("pwd")
		c.JSON(200, gin.H{
			"success": true,
			"user":    user,
			"pwd":     pwd,
		})
	})
	//启动服务
	router.Run(":8080")
}
