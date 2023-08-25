package dao

import (
	"github.com/Biu-X/TikTok/dal/model"
	"github.com/Biu-X/TikTok/dal/query"
	"github.com/Biu-X/TikTok/module/log"
	"gorm.io/gorm"
)

// CreateFavorite 创建用户和视频之间的点赞记录
func CreateFavorite(userID, videoID int64) error {
	f := query.Favorite

	newFavorite := &model.Favorite{
		UserID:  userID,
		VideoID: videoID,
	}

	err := f.Create(newFavorite)
	if err != nil {
		log.Logger.Error(err.Error())
		return err
	}

	return nil
}

// GetUserIsFavoriteVideo 获取用户是否对某视频点赞
func GetUserIsFavoriteVideo(userID int64, videoID int64) bool {
	f := query.Favorite
	count, _ := f.Where(f.UserID.Eq(userID), f.VideoID.Eq(videoID), f.Cancel.Eq(0)).Count()
	return count == 1
}

// GetFavoriteIsExistByUserIDAndVideoID 获取用户是否曾经对某视频点赞
func GetFavoriteIsExistByUserIDAndVideoID(userID int64, videoID int64) bool {
	f := query.Favorite
	count, _ := f.Where(f.UserID.Eq(userID), f.VideoID.Eq(videoID)).Count()
	return count == 1
}

// GetFavoriteByUserID 通过用户ID获取用户的所有点赞记录 cancel=0
func GetFavoriteByUserID(userID int64) (favorites []*model.Favorite, err error) {
	f := query.Favorite
	favorites, err = f.Where(f.UserID.Eq(userID), f.Cancel.Eq(0)).Find()
	return favorites, err
}

// GetFavoriteByID 通过 favorite id 获取对应点赞记录信息
func GetFavoriteByID(id int64) (favorite *model.Favorite, err error) {
	f := query.Favorite

	// 因为接下来使用 First() 调用，避免报错先用 Count 检查
	count, _ := f.Where(f.ID.Eq(id)).Count()
	if count == 0 {
		return &model.Favorite{}, gorm.ErrRecordNotFound
	}

	favorite, err = f.Where(f.ID.Eq(id)).First()
	return favorite, err
}

// GetFavoriteCountByVideoID 通过视频ID获取视频点赞数量
func GetFavoriteCountByVideoID(videoID int64) (count int64, err error) {
	f := query.Favorite
	count, err = f.Where(f.VideoID.Eq(videoID), f.Cancel.Eq(0)).Count()
	return count, err
}

// GetFavoriteCountByUserID 通过用户ID获取用户点赞的视频数量
func GetFavoriteCountByUserID(userID int64) (count int64, err error) {
	f := query.Favorite
	count, err = f.Where(f.UserID.Eq(userID), f.Cancel.Eq(0)).Count()
	return count, err
}

// SetFavoriteCancelByID 通过点赞ID设置是否取消点赞
func SetFavoriteCancelByID(id int64, cancel int32) (err error) {
	f := query.Favorite
	_, err = f.Where(f.ID.Eq(id)).Update(f.Cancel, cancel)
	return err
}

// SetFavoriteCancelByUserIDAndVideoID 通过 用户ID 和 视频ID 取消点赞
func SetFavoriteCancelByUserIDAndVideoID(userID, videoID int64) (err error) {
	f := query.Favorite
	_, err = f.Where(f.VideoID.Eq(videoID), f.UserID.Eq(userID)).Update(f.Cancel, 1)
	return err
}

// SetFavoriteByUserIDAndVideoID 通过 用户ID 和 视频ID 点赞
func SetFavoriteByUserIDAndVideoID(userID, videoID int64) (err error) {
	f := query.Favorite
	// 如果用户曾经点赞过该视频，那么就直接修改脏位以重新点赞；否则就创建点赞记录
	if GetFavoriteIsExistByUserIDAndVideoID(userID, videoID) {
		_, err = f.Where(f.VideoID.Eq(videoID), f.UserID.Eq(userID)).Update(f.Cancel, 0)
		if err != nil {
			log.Logger.Error(err.Error())
			return err
		}
	} else {
		err = CreateFavorite(userID, videoID)
		if err != nil {
			log.Logger.Error(err.Error())
			return err
		}
	}
	return nil
}
