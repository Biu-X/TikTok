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
