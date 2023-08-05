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

func GetVideoByID(id int64) (video *model.Video, err error) {
	video, err = v.Where(v.ID.Eq(id)).First()
	return video, err
}

func GetVideoByAuthorID(authorID int64) (videos []*model.Video, err error) {
	videos, err = v.Where(v.AuthorID.Eq(authorID)).Find()
	return videos, err
}
