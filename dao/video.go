package dao

import (
	"biu-x.org/TikTok/dal/query"
	"biu-x.org/TikTok/model"
)

var v = query.Video

func CreateVideo(video *model.Video) (err error) {
	err = v.Create(video)
	return err
}
