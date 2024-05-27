package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	//创建路由引擎
	router := gin.Default()
	//绑定路由规则
	router.POST("/filetest", func(c *gin.Context) {
		//FormFile返回所提供的表单键的第一个文件
		f, _ := c.FormFile("file")
		//SaveUploadFile上传表单文件到指定的路径
		_ = c.SaveUploadedFile(f, "./"+f.Filename)
		//给前端返回文件
		c.Writer.Header().Add("Content-Disposition", fmt.Sprintf("attachment; filename=%s", f.Filename))
		//File以有效的方式将指定文件写入主体流。
		c.File("./" + f.Filename)
		c.JSON(200, gin.H{
			"msg": f,
		})
	})
	//启动服务
	router.Run(":8080")
}
