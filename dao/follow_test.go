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
	// ----------------------------
	// Test for CreateFollow 每次测试的时候，需要增加参数的值，下次测试是 2, 2（不允许重复关注）
	// ----------------------------
	err := CreateFollow(10, 10)
	if err != nil {
		t.Error("Create Follow fail", err)
		return
	}

	// ----------------------------
	// Test for GetFollowByBoth
	// ----------------------------
	follow, err := GetFollowRelation(10, 10)
	if err != nil {
		t.Error("GetFollowByBoth fail", err)
	}

	expect := &model.Follow{
		UserID:     10,
		FollowerID: 10,
	}

	if follow.UserID != expect.UserID || follow.FollowerID != expect.FollowerID {
		t.Errorf("we expect userId = %v, followId = %v, but got userId = %v, followId = %v", expect.UserID, expect.FollowerID, follow.UserID, follow.FollowerID)
	}

	// ----------------------------
	// Test for GetFollowByID
	// ----------------------------
	followRecord, err := GetFollowRecordByID(follow.ID)
	if err != nil {
		t.Error("GetFollowByID fail", err)
		return
	}

	if !reflect.DeepEqual(follow, followRecord) {
		t.Error("GetFollowByID result error")
		t.Error(follow)
	}

	// ----------------------------
	// Test for SetFollowCancelByID
	// ----------------------------
	err = SetFollowRelationByID(follow.ID, true)
	if err != nil {
		t.Error("SetFollowCancelByID fail", err)
		return
	}
	test_cancel_f, _ := GetFollowRecordByID(follow.ID)
	if test_cancel_f.Cancel == 0 {
		t.Error("SetFollowCancelByID result wrong")
		t.Error(test_cancel_f)
	}

}
