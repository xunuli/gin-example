package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func main() {
	//创建一个路由引擎
	router := gin.Default()
	//绑定路由规则
	//绑定一条获取Token的路由
	router.POST("/auth", authHandle)
	//用户通过上面的路由接口获取token，后续会携带token再来请求其他接口，需要对token进行校验
	//通过校验中间件实现
	router.GET("/home", JWTAuthMiddleware(), func(c *gin.Context) {
		username := c.MustGet("username").(string)
		c.JSON(http.StatusOK, gin.H{
			"code":     2000,
			"messsage": "success",
			"data": gin.H{
				"username": username,
			},
		})
	})

	router.Run(":8080")
}

// 获取用户的token
type UserInfo struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func authHandle(c *gin.Context) {
	//用户发送用户名和密码过来
	var user UserInfo
	//反序列化发送过来的user
	err := c.ShouldBind(&user)
	if err != nil {
		fmt.Println("无效参数！")
		c.JSON(http.StatusOK, gin.H{
			"code":    2001,
			"message": "无效的参数！",
		})
		return
	}
	//校验用户和密码是否正确
	if user.Username == "xuji" && user.Password == "xuji123456" {
		//生成token
		tokenString, _ := GenToken(user.Username, user.Password)
		//将token字符串加入返回响应
		c.JSON(http.StatusOK, gin.H{
			"code":    2000,
			"message": "success",
			"data": gin.H{
				"token": tokenString,
			},
		})
		return
	}
	//如果验证未成功
	c.JSON(http.StatusOK, gin.H{
		"code":    2002,
		"message": "鉴权失败",
	})
}

// 基于JWT鉴权的中间件
func JWTAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 客户端携带Token有三种方式 1.放在请求头 2.放在请求体 3.放在URI
		// 这里假设Token放在Header的Authorization中，并使用Bearer开头
		// 这里的具体实现方式要依据你的实际业务情况决定
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusOK, gin.H{
				"code":    2003,
				"message": "请求头中auth为空",
			})
			//只执行当前中间件，然后返回上一级
			c.Abort()
			return
		}
		//按空格分割
		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			c.JSON(http.StatusOK, gin.H{
				"code":    2004,
				"message": "请求头中auth格式有误",
			})
			c.Abort()
			return
		}
		// parts[1]是获取到的tokenString，我们使用之前定义好的解析JWT的函数来解析它
		fmt.Println(parts)
		mc, err := ParseToken(parts[1])
		fmt.Println(mc)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code": 2005,
				"msg":  "无效的Token",
			})
			c.Abort()
			return
		}
		// 将当前请求的username信息保存到请求的上下文c上
		c.Set("username", mc.username)

		c.Next() // 后续的处理函数可以用过c.Get("username")来获取当前请求的用户信息
	}
}
