package user

import (
	"biu-x.org/TikTok/dal/query"
	"biu-x.org/TikTok/model"
	"biu-x.org/TikTok/module/db"
	"context"
	"fmt"
)

func SaveUser() {
	q := query.Use(db.DB)
	if !q.Available() {
		fmt.Println("query Use(db) fail: query.Available() == false")
	}
	do := q.WithContext(context.Background())

	u := &model.User{
		Name:     "i@hiif.ong",
		Password: "123456",
	}

	err := do.User.Create(u)
	if err != nil {
		fmt.Println(err)
		return
	}
}
