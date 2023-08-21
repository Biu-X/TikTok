package dao

import (
	"errors"

	"biu-x.org/TikTok/dal/query"
	"biu-x.org/TikTok/model"
)

// 创建点赞记录
func CreateFavorite(favorite *model.Favorite) (err error) {
	f := query.Favorite
	err = f.Create(favorite)
	return err
}

// 获取用户对某视频的点赞记录
// 可获取：用户是否曾点赞视频
func GetFavoriteByBoth(userID int64, videoID int64) (favorite *model.Favorite, err error) {
	f := query.Favorite
	// 查询是否存在这条记录，不存在直接返回
	count, _ := f.Where(f.UserID.Eq(userID), f.VideoID.Eq(videoID)).Count()
	if count == 0 {
		return &model.Favorite{}, errors.New("record not found")
	}

	favorite, err = f.Where(f.UserID.Eq(userID), f.VideoID.Eq(videoID)).First()
	return favorite, err
}

// 通过用户ID获取用户的所有点赞记录 cancel=0
func GetFavoriteByUserID(userID int64) (favorites []*model.Favorite, err error) {
	f := query.Favorite
	favorites, err = f.Where(f.UserID.Eq(userID), f.Cancel.Eq(0)).Find()
	return favorites, err
}

// 通过 favorite id 获取对应点赞记录信息
func GetFavoriteByID(id int64) (favorite *model.Favorite, err error) {
	f := query.Favorite

	// 因为接下来使用 First() 调用，避免报错先用 Count 检查
	count, _ := f.Where(f.ID.Eq(id)).Count()
	if count == 0 {
		return &model.Favorite{}, errors.New("record not found")
	}

	favorite, err = f.Where(f.ID.Eq(id)).First()
	return favorite, err
}

// 通过视频ID获取视频点赞数量
func GetFavoriteCountByVideoID(videoID int64) (count int64, err error) {
	f := query.Favorite
	count, err = f.Where(f.VideoID.Eq(videoID), f.Cancel.Eq(0)).Count()
	return count, err
}

// 通过用户ID获取用户点赞的视频数量
func GetFavoriteCountByUserID(userID int64) (count int64, err error) {
	f := query.Favorite
	count, err = f.Where(f.UserID.Eq(userID), f.Cancel.Eq(0)).Count()
	return count, err
}

// 通过点赞ID设置是否取消点赞
func SetFavoriteCancelByID(id int64, cancel int32) (err error) {
	f := query.Favorite
	_, err = f.Where(f.ID.Eq(id)).Update(f.Cancel, cancel)
	return err
}

// 通过视频 ID 设置是否取消点赞
func SetFavoriteCancelByVideoID(videoID int64, cancel int32) (err error) {
	f := query.Favorite
	_, err = f.Where(f.VideoID.Eq(videoID)).Update(f.Cancel, cancel)
	return err
}
