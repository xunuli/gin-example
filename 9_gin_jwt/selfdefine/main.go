package main

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"log"
	"time"
)

// 通过自定义JWT中保存那些数据，比如要保存username
type MyClaims struct {
	//可以自行添加需要的字段
	Username string `json:"username"`
	//内存标准声明
	jwt.RegisteredClaims
}

func main() {
	//自定义声明，生成签名字符串
	fmt.Println("==============生成签名字符串======================")
	//自定义签名私钥
	var mySignedKey = []byte("一等能行！")
	//创建声明
	claims := MyClaims{
		"username",
		jwt.RegisteredClaims{
			//签发时间
			IssuedAt: jwt.NewNumericDate(time.Now()),
			//生效时间
			NotBefore: jwt.NewNumericDate(time.Now()),
			//过期时间
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			//签发人
			Issuer: "my-project",
			//主题
			Subject: "test",
			//ID
			ID: "1",
		},
	}
	//使用指定的签名方法生成token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	//使用指定的签名私钥签名，并获得完整的编码字符串token
	ss, err := token.SignedString(mySignedKey)
	if err != nil {
		log.Fatalf("签名失败！")
	}
	fmt.Println(ss)

	//解析签名字符串
	fmt.Println("==============解析签名字符串======================")
	//如果是自定义的需要用ParseWithClaims方法
	token, err = jwt.ParseWithClaims("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6InVzZXJuYW1lIiwiaXNzIjoibXktcHJvamVjdCIsInN1YiI6InRlc3QiLCJleHAiOjE3MTY5NjM2NDIsIm5iZiI6MTcxNjg3NzI0MiwiaWF0IjoxNzE2ODc3MjQyLCJqdGkiOiIxIn0.20HCAcWY1C9kF-P72ac7MzxW4_acbiY1bxXzpxMnsXw", &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		//返回解析出的密钥
		return mySignedKey, nil
	})
	if err != nil {
		log.Fatalf("解析失败！")
		return
	}
	//对token对象中的Claims进行类型断言
	if claims, ok := token.Claims.(*MyClaims); ok && token.Valid {
		fmt.Println(claims)
	} else {
		log.Fatalf("invaild token!")
	}

}
