package dao

import (
	"reflect"
	"testing"

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

func Test_FollowDAO(t *testing.T) {
	f := &model.Follow{
		UserID:     0,
		FollowerID: 0,
	}

	// ----------------------------
	// Test for CreateFollow
	// ----------------------------
	err := CreateFollow(f)
	if err != nil {
		t.Error("CreateFollow fail", err)
		return
	}

	// ----------------------------
	// Test for GetFollowByBoth
	// ----------------------------
	follow, err := GetFollowByBoth(f.UserID, f.FollowerID)
	if err != nil {
		t.Error("GetFollowByBoth fail", err)
		return
	}
	if !reflect.DeepEqual(follow, f) {
		t.Error("GetFollowByBoth result error")
		t.Error(follow)
	}

	// ----------------------------
	// Test for GetFollowByID
	// ----------------------------
	follow, err = GetFollowByID(f.ID)
	if err != nil {
		t.Error("GetFollowByID fail", err)
		return
	}
	if !reflect.DeepEqual(follow, f) {
		t.Error("GetFollowByID result error")
		t.Error(follow)
	}

	// ----------------------------
	// Test for SetFollowCancelByID
	// ----------------------------
	err = SetFollowCancelByID(f.ID, true)
	if err != nil {
		t.Error("SetFollowCancelByID fail", err)
		return
	}
	test_cancel_f, _ := GetFollowByID(f.ID)
	if test_cancel_f.Cancel == 0 {
		t.Error("SetFollowCancelByID result wrong")
		t.Error(test_cancel_f)
	}

}
