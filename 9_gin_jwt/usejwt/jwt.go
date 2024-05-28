package main

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"log"
	"time"
)

// 自定义声明类型
type MyClaims struct {
	username string
	password string
	jwt.RegisteredClaims
}

// 定义用于签名的密钥
var keySecret = []byte("我一定行！")

// 生成JWT
func GenToken(username, password string) (string, error) {
	//创建一个声明
	claims := MyClaims{
		username,
		password,
		jwt.RegisteredClaims{
			//签发时间
			IssuedAt: jwt.NewNumericDate(time.Now()),
			//生效时间
			//NotBefore: jwt.NewNumericDate(time.Now()),
			//失败时间
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			//签发人
			Issuer: "my-project",
			//主题
			Subject: "test",
			//ID
			ID: "1",
		},
	}
	//使用指定的签名方法创建签名对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	ss, err := token.SignedString(keySecret)
	if err != nil {
		log.Fatalf("签名失败！")
	}
	return ss, nil
}

// 解析JWT
func ParseToken(tokenString string) (*MyClaims, error) {
	//如果自定义claims结构体需要用ParseWithClaims
	token, err := jwt.ParseWithClaims(tokenString, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		return keySecret, nil
	})
	if err != nil {
		log.Fatalf("解析失败！")
		return nil, err
	}
	//对token对象中的claims进行类型断言
	if claims, ok := token.Claims.(*MyClaims); ok && token.Valid {
		fmt.Println(claims)
		return claims, nil
	}
	return nil, errors.New("invaild token")
}
