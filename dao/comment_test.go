package dao

import (
	"testing"

	"github.com/Biu-X/TikTok/dal/model"
	"github.com/Biu-X/TikTok/module/config"
	"github.com/Biu-X/TikTok/module/db"
	"github.com/Biu-X/TikTok/module/log"
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
	// ----------------------------
	// Test for CreateComment
	// ----------------------------
	err := CreateComment(c)
	if err != nil {
		t.Error("CreateComment fail", err)
		return
	}
	t.Log(c)

	// ----------------------------
	// Test for GetCommentByVideoID
	// ----------------------------
	comments, err := GetCommentByVideoID(0)
	if err != nil {
		t.Error("GetCommentByVideoID fail", err)
		return
	}
	t.Log(comments)

	// ----------------------------
	// Test for DeleteCommentByID
	// ----------------------------
	err = DeleteCommentByID(c.ID)
	if err != nil {
		t.Error("DeleteCommentByID fail", err)
		return
	}
}
