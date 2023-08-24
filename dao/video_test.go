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

func Test_VideoDAO(t *testing.T) {
	v := &model.Video{
		AuthorID: 0,
		PlayURL:  "playURL",
		CoverURL: "coverURL",
		Title:    "title",
	}

	// ----------------------------
	// Test for CreateVideo
	// ----------------------------
	err := CreateVideo(v)
	if err != nil {
		t.Error("CreateVideo fail", err)
		return
	}

	// ----------------------------
	// Test for GetVideoByID
	// ----------------------------
	acc, err := GetVideoByID(v.ID)
	if err != nil {
		t.Error("GetVideoByID fail", err)
		return
	}
	if acc.ID != v.ID {
		t.Error("GetVideoByID result wrong")
	}

	// ----------------------------
	// Test for GetVideoByAuthorID
	// ----------------------------
	videos, err := GetVideoByAuthorID(v.AuthorID)
	if err != nil {
		t.Error("GetVideoByAuthorID fail", err)
		return
	}
	if videos[0].ID != v.ID {
		t.Error("GetVideoByAuthorID result wrong")
	}

	// ----------------------------
	// Test for DeleteVideoByID
	// ----------------------------
	err = DeleteVideoByID(v.ID)
	if err != nil {
		t.Error("DeleteVideoByID fail", err)
		return
	}
}
