package dao

import (
	"errors"
	"testing"

	"github.com/Biu-X/TikTok/dal/model"
	"github.com/Biu-X/TikTok/module/config"
	"github.com/Biu-X/TikTok/module/db"
	"github.com/Biu-X/TikTok/module/log"
	"gorm.io/gorm"
)

func init() {
	config.Init()
	log.Init()
	db.Init()
}

// 测试正常插入关注关系和查询插入关系的逻辑
// ----------------------------
// Test for CreateFollow 每次测试的时候，需要修改参数的值（不允许重复关注）
// ----------------------------
func Test_FollowDAO(t *testing.T) {
	userId, followId := 11, 12
	err := CreateFollow(int64(userId), int64(followId))
	if err != nil {
		t.Error("Create Follow fail", err)
		return
	}

	follow, err := GetFollowRelation(int64(userId), int64(followId))
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		t.Error("GetFollowByBoth fail", err)
		return
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		log.Logger.Info("record not found")
		return
	}

	expect := &model.Follow{
		UserID:     11,
		FollowerID: 12,
	}

	if follow.UserID != expect.UserID || follow.FollowerID != expect.FollowerID {
		t.Errorf("we expect userId = %v, followId = %v, but got userId = %v, followId = %v", expect.UserID, expect.FollowerID, follow.UserID, follow.FollowerID)
	}
}

// 测试获取不到记录时的逻辑
// ----------------------------
// Test for CreateFollow 每次测试的时候，需要修改参数的值（不允许重复关注）
// ----------------------------
func Test_FollowRelationNotFound(t *testing.T) {
	_, err := GetFollowRelation(0, 0) // 故意触发 record not found，不允许报错，应该输出的是 Info 信息
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		t.Error("GetFollowByBoth fail", err)
		return
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		log.Logger.Info("record not found")
		return
	}
}

// 测试获取两个用户 id 之间的关注关系功能
// ----------------------------
// Test for GetFollowByBoth
// ----------------------------
func TestIsFollowByBothID(t *testing.T) {
	isFollow, err := GetIsFollowByBothID(11, 12) // 执行这个测试前确保 11, 12 这对关系已经插入
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		t.Error("GetFollowByBoth fail", err)
		return
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		log.Logger.Info("record not found")
		return
	}

	expect := 1

	if !isFollow {
		t.Fatalf("Expect GetIsFollowByBothID result is %v, but got %v", expect, isFollow)
	}
}

// 测试取消关注功能
// ----------------------------
// Test for SetFollowCancelByBoth
// ----------------------------
func TestSetFollowRelationByID(t *testing.T) {
	err := SetFollowCancelByBoth(11, 12)
	if err != nil {
		t.Fatal(err.Error())
	}
	expect := false
	follow, err := GetIsFollowByBothID(11, 12)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		t.Error("GetFollowByBoth fail", err)
		return
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		log.Logger.Info("record not found")
		return
	}

	if follow != expect {
		t.Fatalf("call SetFollowCancelByBoth, expect follow relation is %v, but got %v", expect, follow)
	}
}

// 测试关注功能
// ----------------------------
// Test for SetFollowingByBoth
// ----------------------------
func TestSetFollowingByBoth(t *testing.T) {
	// 以关注关系 userID=11 followerID=12
	// 我们先获取这个关注记录，然后判断是否已经关注，如果没有关注则修改 Cancel 位
	// 然后再次获取关注记录，判断是否关注成功
	err := SetFollowingByBoth(11, 12)
	if err != nil {
		t.Fatal(err.Error())
	}

	follow, err := GetIsFollowByBothID(11, 12)
	log.Logger.Infof("%v", follow)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		t.Error("GetFollowByBoth fail", err)
		return
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		t.Fatal("SetFollowingByBoth failed")
	}

	expect := true
	if follow != expect {
		t.Fatalf("call SetFollowingByBoth, expect follow relation is %v, but got %v", expect, follow)
	}
}
