package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"sync"
	"time"
)

// 第一个路由的handler处理函数
func router1() http.Handler {
	//创建一个路由引擎
	r := gin.New()
	//注册中间件
	r.Use(gin.Recovery())
	r.GET("/test", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"code":    http.StatusOK,
			"message": "Welcome to server1!",
		})
	})
	return r
}

// 第二个路由的handler处理函数
func router2() http.Handler {
	//创建一个路由引擎
	r := gin.New()
	//注册中间件
	r.Use(gin.Recovery())
	r.GET("/test", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"code":    http.StatusOK,
			"message": "Welcome to server2!",
		})
	})
	return r
}

func main() {
	//初始化server1
	server1 := &http.Server{
		Addr:         ":8080",
		Handler:      router1(),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	//初始化server2
	server2 := &http.Server{
		Addr:         ":8081",
		Handler:      router2(),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	//开启两个服务
	var wg sync.WaitGroup
	wg.Add(2)
	go func() error {
		defer wg.Done()
		err := server1.ListenAndServe()
		if err != nil {
			log.Fatalf("server1 failed: %v", err)
			return err
		}
		return nil
	}()
	go func() error {
		defer wg.Done()
		err := server2.ListenAndServe()
		if err != nil {
			log.Fatalf("server1 failed: %v", err)
			return err
		}
		return nil
	}()
	wg.Wait()
}
