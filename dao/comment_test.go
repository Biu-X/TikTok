package dao

import (
	"testing"

	"biu-x.org/TikTok/dal/query"
	"biu-x.org/TikTok/model"
	"biu-x.org/TikTok/module/config"
	"biu-x.org/TikTok/module/db"
	"biu-x.org/TikTok/module/log"
)

func init() {
	config.Init()
	log.Init()
	db.Init()
}

func Test_CommentDAO(t *testing.T) {
	c := &model.Comment{
		UserID:  0,
		VideoID: 0,
		Content: "test_content",
	}
	err := query.Comment.Create(c)
	if err != nil {
		t.Error("CreateComment fail", err)
		return
	}
	t.Log(c)
}
