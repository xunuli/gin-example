package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 绑定参数
type Login struct {
	User     string `form:"user" json:"user" binding:"required"`
	Password string `form:"password" json:"password" binding:"required" `
}

func main() {
	router := gin.Default()

	//绑定JSON实例 {"user":"xuji", "password":"123456"}
	router.POST("/loginJson", func(c *gin.Context) {
		var login Login
		//反序列化：将序列化数据转化为结构体
		if err := c.ShouldBind(&login); err == nil {
			fmt.Printf("login info:%#v\n", login)
			c.JSON(200, gin.H{
				"user":     login.User,
				"password": login.Password,
			})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		}
	})

	//绑定form表单示例（user=xuji&password=123456）
	router.POST("/loginForm", func(c *gin.Context) {
		var login Login
		//shoulfBind会根据请求的content-type自行选择绑定器
		if err := c.ShouldBind(&login); err == nil {
			c.JSON(200, gin.H{
				"user":     login.User,
				"password": login.Password,
			})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		}
	})

	//启动服务
	router.Run(":8080")
}
