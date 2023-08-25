package dao

import (
	"reflect"
	"testing"

	"github.com/Biu-X/TikTok/dal/model"
	"github.com/Biu-X/TikTok/dal/query"
	"github.com/Biu-X/TikTok/module/config"
	"github.com/Biu-X/TikTok/module/db"
	"github.com/Biu-X/TikTok/module/log"
)

func init() {
	config.Init()
	log.Init()
	db.Init()
}

func Test_UserDAO(t *testing.T) {
	u := &model.User{
		Name:            "Zaire",
		Password:        "root",
		Signature:       "signature",
		Avatar:          "avatar",
		BackgroundImage: "backgroundimage",
	}

	t.Log(u)

	// ----------------------------
	// Test for CreateUser
	// ----------------------------
	err := CreateUser(u)
	if err != nil {
		t.Error("CreateUser fail", err)
		return
	}

	// ----------------------------
	// Test for GetUserByID
	// ----------------------------
	acc, err := GetUserByID(u.ID)
	if err != nil {
		t.Error("GetUserByID fail", err)
		return
	}
	if !reflect.DeepEqual(acc, u) {
		t.Error("GetUserByID result wrong")
	}

	// ----------------------------
	// Test for GetUserByName
	// ----------------------------
	acc, err = GetUserByName(u.Name)
	if err != nil {
		t.Error("GetUserByName fail", err)
		return
	}
	if !reflect.DeepEqual(acc, u) {
		t.Error("GetUserByName result wrong")
	}

	// ----------------------------
	// Test for SetAvatarByID
	// ----------------------------
	err = SetAvatarByID(u.ID, "test_avatar")
	if err != nil {
		t.Error("SetAvatarByID fail", err)
		return
	}
	acc, err = GetUserByID(u.ID)
	if err != nil {
		t.Error("GetUserByID fail when test for SetAvatarByID", err)
	}
	if acc.Avatar != "test_avatar" {
		t.Error("SetAvatarByID result wrong")
	}

	// ----------------------------
	// Test for SetSignatureByID
	// ----------------------------
	err = SetSignatureByID(u.ID, "test_signature")
	if err != nil {
		t.Error("SetSignatureByID fail", err)
		return
	}
	acc, err = GetUserByID(u.ID)
	if err != nil {
		t.Error("GetUserByID fail when test for SetSignatureByID", err)
	}
	if acc.Signature != "test_signature" {
		t.Error("SetSignatureByID result wrong")
	}

	// ----------------------------
	// Test for SetPasswordByID
	// ----------------------------
	err = SetPasswordByID(u.ID, "test_password")
	if err != nil {
		t.Error("SetPasswordByID fail", err)
		return
	}
	acc, err = GetUserByID(u.ID)
	if err != nil {
		t.Error("GetUserByID fail when test for SetPasswordByID", err)
	}
	if acc.Password != "test_password" {
		t.Error("SetPasswordByID result wrong")
	}

	// ----------------------------
	// Test for SetBackgroundImageByID
	// ----------------------------
	err = SetBackgroundImageByID(u.ID, "test_background_image")
	if err != nil {
		t.Error("SetBackgroundImageByID fail", err)
		return
	}
	acc, err = GetUserByID(u.ID)
	if err != nil {
		t.Error("GetUserByID fail when test for SetBackgroundImageByID", err)
	}
	if acc.BackgroundImage != "test_background_image" {
		t.Error("SetBackgroundImageByID result wrong")
	}

	// ----------------------------
	// Test for SetNameByID
	// ----------------------------
	err = SetNameByID(u.ID, "test_name")
	if err != nil {
		t.Error("SetNameByID fail", err)
		return
	}
	acc, err = GetUserByID(u.ID)
	if err != nil {
		t.Error("GetUserByID fail when test for SetNameByID", err)
	}
	if acc.Name != "test_name" {
		t.Error("SetNameByID result wrong")
	}

	// ----------------------------
	// Test for DeleteUserByID
	// ----------------------------
	err = DeleteUserByID(u.ID)
	if err != nil {
		t.Error("DeleteUserByID fail", err)
		return
	}
}

func TestSaveUser(t *testing.T) {
	config.Init()
	log.Init()
	db.Init()
	err := query.Comment.Create(&model.Comment{Content: "hello"})
	if err != nil {
		log.Logger.Info(err)
	}
	info, err := query.Comment.Where(query.Comment.ID.Eq(1)).Delete()
	if err != nil {
		log.Logger.Info(err)
	}
	log.Logger.Infof("info: %v", info)
}
