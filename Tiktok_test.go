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
	user := &model.User{Name: "smy", ID: 1}
	errStr := query.User.Create(user).Error()

	if len(errStr) != 0 {
		t.Fatalf("err: %v", errStr)
	}
	// 模拟登录成功
	token := jwt.GenerateToken("smy")
	fmt.Printf("token: %#v", token)

	useClaims, err := jwt.ParseToken(token)
	if err != nil {
		t.Fatalf("jwt token generate success, but vaild failed....")
	}

	log.Println(*useClaims)
}
