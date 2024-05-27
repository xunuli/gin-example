package main

import "github.com/gin-gonic/gin"

func main() {
	//创建路由引擎
	router := gin.Default()
	//绑定路由规则
	router.POST("/path", func(c *gin.Context) {
		//PostForm从POST url编码表或多部份表单返回指定的键值
		user := c.PostForm("user")
		pwd := c.PostForm("pwd")
		//返回JSON格式的数据
		c.JSON(200, gin.H{
			"success": true,
			"user":    user,
			"pwd":     pwd,
		})
	})
	//启动服务
	router.Run(":8080")
}
