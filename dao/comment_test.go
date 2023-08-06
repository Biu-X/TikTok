package dao

import (
	"testing"

	"biu-x.org/TikTok/model"
	"biu-x.org/TikTok/module/config"
	"biu-x.org/TikTok/module/db"
)

func init() {
	config.Init()
	db.Init()
}

func Test_CommentDAO(t *testing.T) {
	c := &model.Comment{
		UserID:  0,
		VideoID: 0,
		Content: "test_content",
	}
	err := CreateComment(c)
	if err != nil {
		t.Error("CreateComment fail", err)
		return
	}
	t.Log(c)
}
