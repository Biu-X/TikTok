package dao

import (
	"biu-x.org/TikTok/dal/query"
	"biu-x.org/TikTok/model"
)

func CreateFollow(follow *model.Follow) (err error) {
	f := query.Follow
	err = f.Create(follow)
	return err
}

func GetFollowByID(id int64) (follow *model.Follow, err error) {
	f := query.Follow
	follow, err = f.Where(f.ID.Eq(id)).First()
	return follow, err
}

// 查询粉丝的follow表记录 cancel=0
func GetFollowByUserID(userID int64) (follows []*model.Follow, err error) {
	f := query.Follow
	follows, err = f.Where(f.UserID.Eq(userID), f.Cancel.Eq(0)).Find()
	return follows, err
}

// 查询关注的人的follow表记录 cancel=0
func GetFollowByFollowerID(userID int64) (follows []*model.Follow, err error) {
	f := query.Follow
	follows, err = f.Where(f.FollowerID.Eq(userID), f.Cancel.Eq(0)).Find()
	return follows, err
}

// 查询关注的人的ID
func GetFollowingIDByFollowerID(userID int64) (id []int64, err error) {
	follows, err := GetFollowByFollowerID(userID)
	for _, follow := range follows {
		id = append(id, follow.UserID)
	}
	return id, err
}

// 查询关注的人的数量
func GetFollowingCountByFollowerID(userID int64) (count int64, err error) {
	f := query.Follow
	count, err = f.Where(f.FollowerID.Eq(int64(userID)), f.Cancel.Eq(0)).Count()
	return count, err
}

// 查询用户的粉丝的ID
func GetFollowerIDByUserID(userID int64) (id []int64, err error) {
	follows, err := GetFollowByFollowerID(userID)
	for _, follow := range follows {
		id = append(id, follow.FollowerID)
	}
	return id, err
}

// 查询用户粉丝数量
func GetFollowerCountByUserID(userID int64) (count int64, err error) {
	f := query.Follow
	count, err = f.Where(f.UserID.Eq(int64(userID)), f.Cancel.Eq(0)).Count()
	return count, err
}

func GetFollowByBoth(userID int64, followID int64) (follow *model.Follow, err error) {
	f := query.Follow
	follow, err = f.Where(f.UserID.Eq(userID), f.FollowerID.Eq(followID)).First()
	return follow, err
}

func SetFollowCancelByID(id int64, cancel bool) (err error) {
	f := query.Follow
	_, err = f.Where(f.ID.Eq(id)).Update(f.Cancel, cancel)
	return err
}
