package dao

import (
	"errors"

	"biu-x.org/TikTok/dal/query"
	"biu-x.org/TikTok/model"
	"biu-x.org/TikTok/module/log"
	"gorm.io/gorm"
)

// CreateFollow 创建用户和粉丝之间的关系记录
func CreateFollow(userId, followerId int64) error {
	f := query.Follow

	newFollow := &model.Follow{
		UserID:     userId,
		FollowerID: followerId,
	}

	err := f.Create(newFollow)
	if err != nil {
		log.Logger.Error(err.Error())
		return err
	}

	return nil
}

// 通过 follow 表的主键 ID 获取对应的关注记录
func GetFollowRecordByID(id int64) (*model.Follow, error) {
	f := query.Follow

	if count, _ := f.Where(f.ID.Eq(id)).Count(); count == 0 {
		log.Logger.Debug("record not found")
		return &model.Follow{}, errors.New("record not found")
	}

	return f.Where(f.ID.Eq(id)).First()
}

// 1 Follower 部分
// 查询该用户被哪些粉丝关注，找出那些记录
func GetFollowerRecordsByUserID(userID int64) ([]*model.Follow, error) {
	f := query.Follow

	followerRecords, err := f.Where(f.UserID.Eq(userID), f.Cancel.Eq(0)).Find()
	if err != nil {
		log.Logger.Error(err.Error())
		return followerRecords, err
	}

	return followerRecords, nil
}

// 返回指定用户的所有粉丝 id 并以切片形式返回
func GetFollowerIDsByUserID(userID int64) ([]int64, error) {
	// 1. 找到所有粉丝的记录
	follows, err := GetFollowerRecordsByUserID(userID)
	if err != nil {
		log.Logger.Error(err.Error())
		return []int64{}, err
	}

	followerIDs := []int64{}
	// 2. 获取粉丝记录中的粉丝 id 以切片形式返回
	for _, follow := range follows {
		followerIDs = append(followerIDs, follow.FollowerID)
	}

	return followerIDs, nil
}

// 查询用户粉丝数量
func GetFollowerCountByUserID(userID int64) (int64, error) {
	f := query.Follow
	return f.Where(f.UserID.Eq(int64(userID)), f.Cancel.Eq(0)).Count()
}

// 2. Following
// 查询用户所有关注的人的记录
func GetFollowingRecordsByUserID(userID int64) ([]*model.Follow, error) {
	f := query.Follow

	followingRecords, err := f.Where(f.FollowerID.Eq(userID), f.Cancel.Eq(0)).Find()
	if err != nil {
		log.Logger.Error(err.Error())
		return followingRecords, err
	}

	return followingRecords, nil
}

// 查询用户所有关注的人的用户 id，并以切片形式返回
func GetFollowingIdsByUserID(userID int64) ([]int64, error) {
	followingRecords, err := GetFollowingRecordsByUserID(userID)
	if err != nil {
		return []int64{}, err
	}

	followingIds := []int64{}
	for _, following := range followingRecords {
		followingIds = append(followingIds, following.UserID)
	}

	return followingIds, nil
}

// 查询指定用户关注的人的数量
func GetFollowingCountByUserID(userID int64) (int64, error) {
	f := query.Follow
	return f.Where(f.FollowerID.Eq(int64(userID)), f.Cancel.Eq(0)).Count()
}

// 查询两个用户之间的关注信息
// 第一个参数是指定用户 ID，第二个参数是粉丝 ID
// 查询 UserID = userID 并且 FollowID = FollowId 的那条记录并返回
func GetFollowRelation(userID int64, followerID int64) (*model.Follow, error) {
	f := query.Follow

	if count, _ := f.Where(f.UserID.Eq(userID), f.FollowerID.Eq(followerID)).Count(); count == 0 {
		log.Logger.Debug("record not found")
		return &model.Follow{}, errors.New("record not found")
	}

	follow, err := f.Where(f.UserID.Eq(userID), f.FollowerID.Eq(followerID)).First()
	if err != nil {
		log.Logger.Error(err.Error())
		return follow, err
	}

	return follow, nil
}

// 查询两人的关注信息, 第二个 Id 表示粉丝的 ID，第一个 Id 代表用户的 ID
// 我们将要判断粉丝 ID 对应的用户是否关注了指定的用户
// 我们可以通过判断这条记录的 Cancel 字段是否为 0 得知
func GetIsFollowByBothID(userID int64, followerID int64) (bool, error) {
	if userID == followerID {
		return false, nil
	}

	follow, err := GetFollowRelation(userID, followerID)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return false, nil
	}

	if err != nil {
		log.Logger.Error(err.Error())
		return false, err
	}

	return follow.Cancel == 0, nil
}

// 设置某条记录的关注关系
func SetFollowRelationByID(id int64, cancel bool) error {
	f := query.Follow

	_, err := f.Where(f.ID.Eq(id)).Update(f.Cancel, cancel)
	if err != nil {
		log.Logger.Error(err.Error())
		return err
	}

	return nil
}

// 粉丝取关用户
func SetFollowCancelByBoth(userID int64, followerID int64) error {
	f := query.Follow

	_, err := f.Where(f.UserID.Eq(userID), f.FollowerID.Eq(followerID)).Update(f.Cancel, true)
	if err != nil {
		log.Logger.Error(err.Error())
		return err
	}

	return nil
}

// 粉丝关注用户，若之前已经关注又取关，则修改脏位
func SetFollowFollowByBoth(userID int64, followerID int64) error {
	// 查询两人之间的关注关系，若不存在则创建
	_, err := GetFollowRelation(userID, followerID)
	// 1. 非常规错误
	if err != nil {
		log.Logger.Error(err.Error())
		return err
	}
	// 2. 关系不存在时的处理
	if errors.Is(err, gorm.ErrRecordNotFound) {
		err = CreateFollow(userID, followerID)
		if err != nil {
			log.Logger.Error(err.Error())
			return err
		}
		return nil
	}

	// 3. 若关系存在则修改脏位
	f := query.Follow
	_, err = f.Where(f.UserID.Eq(userID), f.FollowerID.Eq(followerID)).Update(f.Cancel, false)
	if err != nil {
		log.Logger.Error(err.Error())
		return err
	}

	return nil
}
