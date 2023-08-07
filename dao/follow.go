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

func GetFollowUserByUserID(userID int64) (follows []*model.Follow, err error) {
	f := query.Follow
	follows, err = f.Where(f.UserID.Eq(userID)).Find()
	return follows, err
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
