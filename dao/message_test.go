package dao

import (
	"reflect"
	"testing"
	"time"

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

func Test_MessageDAO(t *testing.T) {
	f := &model.Message{
		FromUserID: 0,
		ToUserID:   0,
		Content:    "test_content",
	}
	// ----------------------------
	// Test for CreateMessage
	// ----------------------------
	err := CreateMessage(f)
	if err != nil {
		t.Error("CreateMessage fail", err)
		return
	}

	// ----------------------------
	// Test for GetMessageByID
	// ----------------------------
	message, err := GetMessageByID(f.ID)
	if err != nil {
		t.Error("GetMessageByID fail", err)
		return
	}
	if !reflect.DeepEqual(message, f) {
		t.Error("GetMessageByID result wrong")
		t.Error(message)
	}

	// ----------------------------
	// Test for GetMessageByBoth
	// ----------------------------
	messages, err := GetMessageByBoth(f.FromUserID, f.ToUserID, time.Unix(0, 0))
	if err != nil {
		t.Error("DeleteMessageByID fail", err)
		return
	}
	for _, message := range messages {
		t.Log(message)
	}
}
