package dao

import (
	"testing"

	"biu-x.org/TikTok/model"
)

func Test_UserDAO(t *testing.T) {
	u := &model.User{
		Name:     "Zaire",
		Password: "root",
	}
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
	if acc != u {
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
	if acc != u {
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

}
