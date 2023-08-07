package dao

import (
	"biu-x.org/TikTok/dal/query"
	"biu-x.org/TikTok/model"
)

func CreateFavorite(favorite *model.Favorite) (err error) {
	f := query.Favorite
	err = f.Create(favorite)
	return err
}

func GetFavoriteByBoth(userID int64, videoID int64) (favorite *model.Favorite, err error) {
	f := query.Favorite
	favorite, err = f.Where(f.UserID.Eq(userID), f.VideoID.Eq(videoID)).First()
	return favorite, err
}

func GetFavoriteByUserID(userID int64) (favorites []*model.Favorite, err error) {
	f := query.Favorite
	favorites, err = f.Where(f.UserID.Eq(userID)).Find()
	return favorites, err
}

func GetFavoriteByID(id int64) (favorite *model.Favorite, err error) {
	f := query.Favorite
	favorite, err = f.Where(f.ID.Eq(id)).First()
	return favorite, err
}
func SetFavoriteCancelByID(id int64, cancel bool) (err error) {
	f := query.Favorite
	_, err = f.Where(f.ID.Eq(id)).Update(f.Cancel, cancel)
	return err
}
