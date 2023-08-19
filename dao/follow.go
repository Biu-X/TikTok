package dao

import (
	"biu-x.org/TikTok/dal/query"
	"biu-x.org/TikTok/model"
)

// CreateFollow 创建关注记录
// FollowerID ->关注 UserID，则写成
//
//	&model.Follow{
//		 UserID:     UserID, // 用户id
//		 FollowerID: FollowerID, // 粉丝id
//	}
func CreateFollow(follow *model.Follow) (err error) {
	f := query.Follow
	err = f.Create(follow)
	return err
}

// 通过ID获取关注记录
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

// 返回 userID 的所有 粉丝，Follow.FollowerID 是粉丝的ID
func GetFollowFollowerIDsByUserID(userID int64) (followerIDs []int64, err error) {
	follows, err := GetFollowByUserID(userID)
	for _, follow := range follows {
		followerIDs = append(followerIDs, follow.FollowerID)
	}
	return followerIDs, err
}

// 查询关注的人的follow表记录 cancel=0
func GetFollowByFollowerID(userID int64) (follows []*model.Follow, err error) {
	f := query.Follow
	follows, err = f.Where(f.FollowerID.Eq(userID), f.Cancel.Eq(0)).Find()
	return follows, err
}

// 返回 userID 关注的所有人，Follow.UserID 是关注的人的ID
func GetFollowUserIDsByUserID(userID int64) (userIDs []int64, err error) {
	follows, err := GetFollowByFollowerID(userID)
	for _, follow := range follows {
		userIDs = append(userIDs, follow.UserID)
	}
	return userIDs, err
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

// 查询两人的关注信息, 可获取 第二个用户 是否关注了 第一个用户，返回 follow表记录
func GetFollowByBoth(userID int64, followerID int64) (follow *model.Follow, err error) {
	f := query.Follow
	follow, err = f.Where(f.UserID.Eq(userID), f.FollowerID.Eq(followerID)).First()
	return follow, err
}

// 查询两人的关注信息, 可获取 第二个用户 是否关注了 第一个用户，返回值为 bool
func GetIsFollowByBoth(userID int64, followerID int64) bool {
	follow, err := GetFollowByBoth(userID, followerID)
	if err != nil {
		return false
	}
	return follow.Cancel == 0
}

// 通过记录ID设置是否取关
func SetFollowCancelByID(id int64, cancel bool) (err error) {
	f := query.Follow
	_, err = f.Where(f.ID.Eq(id)).Update(f.Cancel, cancel)
	return err
}

// SetFollowCancelByBoth 取消关注，第二个用户取消关注第一个用户
// FollowerID ->取关 userId，则应该是
//
// err := dao.SetFollowCancelByBoth(userId, FollowerID) // 粉丝取关用户
func SetFollowCancelByBoth(userID int64, followerID int64) (err error) {
	f := query.Follow
	_, err = f.Where(f.UserID.Eq(userID), f.FollowerID.Eq(followerID)).Update(f.Cancel, true)
	return err
}
