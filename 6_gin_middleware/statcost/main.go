package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

// StatCost 是一个统计耗时请求的中间件
func StatCost() gin.HandlerFunc {
	return func(c *gin.Context) {
		//请求开始时间
		start := time.Now()
		//可以通过c.Set在请求上下文中设置值，后续的处理函数能够取到该值
		c.Set("name", "小王子")
		//调用该请求的剩余处理程序
		c.Next()
		//不调用该请求的剩余处理程序
		//c.Abort()
		cost := time.Since(start)
		log.Println(cost)
	}
}

func main() {
	//创建一个没有任何默认中间件的路由
	router := gin.New()
	//注册一个全局中间件
	router.Use(StatCost())
	//绑定路由规则
	router.GET("/test", func(c *gin.Context) {
		//从上下文取值
		name := c.MustGet("name").(string)
		log.Println(name, "你好！")
		c.JSON(200, gin.H{
			"message": "hello world",
		})
	})
	//启动服务
	_ = router.Run()
}
