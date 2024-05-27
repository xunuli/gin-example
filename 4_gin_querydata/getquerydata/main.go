package main

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
)

func main() {
	//创建一个路由引擎
	router := gin.Default()
	//绑定路由规则
	//获取querystring参数
	router.GET("/user/search", func(c *gin.Context) {
		username := c.DefaultQuery("username", "小王子")
		address := c.Query("address")
		//输出结果给调用方
		c.JSON(200, gin.H{
			"message":  "ok",
			"username": username,
			"address":  address,
		})
	})
	//获取form参数
	router.POST("/user/search", func(c *gin.Context) {
		//获取form中的值
		username := c.DefaultPostForm("username", "小王子")
		address := c.PostForm("address")
		//输出json给调用方
		c.JSON(200, gin.H{
			"message":  "ok",
			"username": username,
			"address":  address,
		})
	})
	//获取json参数
	router.POST("/json", func(c *gin.Context) {
		b, _ := c.GetRawData() //从c.request.body读取请求数据
		//定义map或结构体
		var m map[string]interface{}
		//反序列化
		_ = json.Unmarshal(b, &m)
		c.JSON(200, m)
	})

	router.Run(":8080")
}
