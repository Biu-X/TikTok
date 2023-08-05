package main

import (
	"fmt"
	"log"
	"testing"

	"biu-x.org/TikTok/dal/query"
	"biu-x.org/TikTok/model"
	"biu-x.org/TikTok/module/config"
	"biu-x.org/TikTok/module/db"
	"biu-x.org/TikTok/module/middleware/jwt"
)

func TestSingToken(t *testing.T) {
	// 初始 mysql 连接
	config.Init()
	db.Init()
	// 模拟新建一个用户
	fmt.Printf("err: %v\n", 1)
	u := query.Use(db.DB).User
	user := &model.User{Name: "Test9", Password: "abc", Signature: "newtess", Avatar: "avatar", BackgroundImage: "background"}
	_ = u.Create(user)
	fmt.Printf("err: %v\n", 2)

	// 模拟登录成功
	token := jwt.GenerateToken("Test9")
	fmt.Printf("err: %v\n", 3)
	fmt.Printf("token: %#v", token)

	useClaims, err := jwt.ParseToken(token)
	if err != nil {
		t.Fatalf("jwt token generate success, but vaild failed....")
	}

	log.Println(*useClaims)
}
