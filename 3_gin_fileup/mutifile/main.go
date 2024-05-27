package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	//创建路由引擎
	router := gin.Default()
	//绑定路由规则
	router.POST("/multifile", func(c *gin.Context) {
		form, err := c.MultipartForm()
		if err != nil {
			log.Fatalf("failed: %v", err)
		}
		//通过字段名映射
		f := form.File["file"]
		//for range遍历文件
		for _, file := range f {
			fmt.Println(file.Filename)
			c.SaveUploadedFile(file, "./"+file.Filename)
			c.Writer.Header().Add("Content-Disposition",
				fmt.Sprintf("attachment; filename=%s"+file.Filename))
			c.File("./" + file.Filename)
		}
	})
	//启动服务
	router.Run(":8080")
}
