package main

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"log"
	"time"
)

//使用JWT默认的字段，没有其他定制化的需求，则可以快速生成和解析token

func main() {
	fmt.Println("================生成签名字符串===================")
	//用于签名的字符串，密钥
	var mySigningKey = []byte("woshixuji")
	//创建claims，标准声明
	claims := &jwt.RegisteredClaims{
		//过期时间，设置24小时
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
		//签发人
		Issuer: "xuji",
	}
	//生成token对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	//生成签名字符串
	ss, _ := token.SignedString(mySigningKey)
	//eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJ4dWppIiwiZXhwIjoxNzE2OTU3Mjc0fQ.belmmtz56Gr0ait2tUvei6MkCvkurWkkqEU8gS4E3mE
	fmt.Println(ss)
	fmt.Println("==============解析token=====================")

	token, err := jwt.Parse(ss, func(token *jwt.Token) (interface{}, error) {
		return mySigningKey, nil
	})
	if err != nil {
		log.Fatalf("解析失败！")
		return
	}
	fmt.Println(token.Valid)
	fmt.Println(token.Claims)
}
