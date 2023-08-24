package dao

import (
	"errors"

	"github.com/Biu-X/TikTok/dal/model"
	"github.com/Biu-X/TikTok/dal/query"
	"github.com/Biu-X/TikTok/module/log"
	"gorm.io/gorm"
)

// 使用须知
// 一、所有以 Find()、Count() Finisher 函数结尾的调用都不会触发 "record not found" 错误
// 二、所有以 First() Finisher 函数结尾的调用在查询不到指定记录时会报错，因此在调用 First() 函数前，我们先使用查询空记录也不会报错的 Find() 或者 Count() （推荐），
// 判断 Count() 返回结果是否为 0，如果是说明查询不到指定记录，这时我们直接向上层返回 gorm.ErrRecordNotFound 错误，上层逻辑只要再对返回的 err 使用 errors.Is 进行判断即可知道调用的三种返回状态：
// 1. 未查询到 errors.Is(err, gorm.ErrRecordNotFound)
// 2. 查询错误 err != nil
// 3. 查询成功 err == nil

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

// GetFollowRecordByID 通过 follow 表的主键 ID 获取对应的关注记录
func GetFollowRecordByID(id int64) (*model.Follow, error) {
	f := query.Follow

	count, _ := f.Where(f.ID.Eq(id)).Count()
	if count == 0 {
		return &model.Follow{}, gorm.ErrRecordNotFound
	}

	// 到这里就一定可以查询到记录了
	followingRecord, err := f.Where(f.ID.Eq(id)).First()
	if err != nil {
		log.Logger.Error(err.Error())
		return followingRecord, err
	}

	return followingRecord, nil
}

// GetFollowerRecordsByUserID 查询该用户被哪些粉丝关注，找出那些记录
func GetFollowerRecordsByUserID(userID int64) ([]*model.Follow, error) {
	f := query.Follow

	followerRecords, err := f.Where(f.UserID.Eq(userID), f.Cancel.Eq(0)).Find()
	if err != nil {
		log.Logger.Error(err.Error())
		return followerRecords, err
	}

	return followerRecords, nil
}

// GetFollowerIDsByUserID 返回指定用户的所有粉丝 id 并以切片形式返回
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

// GetFollowerCountByUserID 查询用户粉丝数量
func GetFollowerCountByUserID(userID int64) (int64, error) {
	f := query.Follow

	count, err := f.Where(f.UserID.Eq(int64(userID)), f.Cancel.Eq(0)).Count()
	if err != nil {
		log.Logger.Error(err.Error())
		return count, err
	}

	return count, nil
}

// 2. Following
// GetFollowingRecordsByUserID 查询用户所有关注的人的记录
func GetFollowingRecordsByUserID(userID int64) ([]*model.Follow, error) {
	f := query.Follow

	followingRecords, err := f.Where(f.FollowerID.Eq(userID), f.Cancel.Eq(0)).Find()
	if err != nil {
		log.Logger.Error(err.Error())
		return followingRecords, err
	}

	return followingRecords, nil
}

// GetFollowingIdsByUserID 查询用户所有关注的人的用户 id，并以切片形式返回
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

// GetFollowingCountByUserID 查询指定用户关注的人的数量
func GetFollowingCountByUserID(userID int64) (int64, error) {
	f := query.Follow

	count, err := f.Where(f.FollowerID.Eq(int64(userID)), f.Cancel.Eq(0)).Count()
	if err != nil {
		log.Logger.Error(err.Error())
		return 0, err
	}

	return count, nil
}

// GetFollowRelation
// 查询两个用户之间的关注信息
// 第一个参数是指定用户 ID，第二个参数是粉丝 ID
// 查询 UserID = userID 并且 FollowID = FollowId 的那条记录并返回
func GetFollowRelation(userID int64, followerID int64) (*model.Follow, error) {
	f := query.Follow

	count, _ := f.Where(f.UserID.Eq(userID), f.FollowerID.Eq(followerID)).Count()
	if count == 0 {
		return &model.Follow{}, gorm.ErrRecordNotFound
	}

	follow, err := f.Where(f.UserID.Eq(userID), f.FollowerID.Eq(followerID)).First()
	if err != nil {
		log.Logger.Error(err.Error())
		return follow, err
	}

	return follow, nil
}

// GetIsFollowByBothID
// 查询两人的关注信息, 第二个 Id 表示粉丝的 ID，第一个 Id 代表用户的 ID
// 我们将要判断粉丝 ID 对应的用户是否关注了指定的用户
// 我们可以通过判断这条记录的 Cancel 字段是否为 0 得知
func GetIsFollowByBothID(userID int64, followerID int64) (bool, error) {
	if followerID == 0 || userID == followerID {
		return false, nil
	}

	follow, err := GetFollowRelation(userID, followerID)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		log.Logger.Info(err.Error())
		return false, nil // 查询不到直接返回 false 即可，这不是错误
	}

	return follow.Cancel == 0, nil
}

// SetFollowRelationByID 设置某条记录的关注关系
func SetFollowRelationByID(id int64, cancel bool) error {
	var flag int32
	if cancel {
		flag = 1
	} else {
		flag = 0
	}
	f := query.Follow

	_, err := f.Where(f.ID.Eq(id)).Update(f.Cancel, flag)
	if err != nil {
		log.Logger.Error(err.Error())
		return err
	}

	return nil
}

// SetFollowCancelByBoth 粉丝取关用户
func SetFollowCancelByBoth(userID int64, followerID int64) error {
	f := query.Follow
	_, err := f.Where(f.UserID.Eq(userID), f.FollowerID.Eq(followerID)).Update(f.Cancel, 1)
	if err != nil {
		log.Logger.Error(err.Error())
		return err
	}

	return nil
}

// SetFollowingByBoth 粉丝关注用户，若之前已经关注又取关，则修改脏位
func SetFollowingByBoth(userID int64, followerID int64) error {
	// 查询两人之间的关注关系，若不存在则创建
	_, err := GetFollowRelation(userID, followerID)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		err = CreateFollow(userID, followerID)
		if err != nil {
			log.Logger.Debug(err.Error())
			return err
		}
		return nil
	} else if err != nil {
		log.Logger.Debug(err.Error())
		return err
	}
	log.Logger.Info("GetFollowRelation Success")

	// 若存在则修改脏位
	f := query.Follow

	_, err = f.Where(f.UserID.Eq(userID), f.FollowerID.Eq(followerID)).Update(f.Cancel, 0)
	if err != nil {
		log.Logger.Error(err.Error())
		return err
	}
	return nil
}
