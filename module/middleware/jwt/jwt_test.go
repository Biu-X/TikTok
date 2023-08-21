package jwt_test

import (
	"biu-x.org/TikTok/module/log"
	"fmt"
	"testing"

	"biu-x.org/TikTok/dal/query"
	"biu-x.org/TikTok/dal/model"
	"biu-x.org/TikTok/module/config"
	"biu-x.org/TikTok/module/db"
	"biu-x.org/TikTok/module/middleware/jwt"
)

func TestSingToken1(t *testing.T) {
	// 配置初始化
	config.Init()
	log.Init()
	db.Init()

	// 模拟注册（新建一个用户）
	u := query.User
	user := &model.User{Name: "newuser", Password: "abc", Signature: "newtess", Avatar: "avatar", BackgroundImage: "background"}
	err := u.Create(user)
	if err != nil {
		t.Fatalf("create user failed, err: %v", err)
	}

	// 模拟登录（生成 JWT token）
	token := jwt.GenerateToken("TableNewUser")
	fmt.Printf("token: %#v\n", token)

	useClaims, err := jwt.ParseToken(token)
	if err != nil {
		t.Fatalf("jwt token generate success, but vaild failed....")
	}

	fmt.Printf("token 解析成功，获取到的用户信息：%#v\n", *useClaims)
}
