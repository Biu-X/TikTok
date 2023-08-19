package dao

import (
	"biu-x.org/TikTok/dal/query"
	"biu-x.org/TikTok/model"
)

func CreateVideo(video *model.Video) (err error) {
	v := query.Video
	err = v.Create(video)
	return err
}

func GetVideoByID(id int64) (video *model.Video, err error) {
	v := query.Video
	video, err = v.Where(v.ID.Eq(id)).First()
	return video, err
}

func GetVideoByAuthorID(authorID int64) (videos []*model.Video, err error) {
	v := query.Video
	videos, err = v.Where(v.AuthorID.Eq(authorID)).Find()
	return videos, err
}

func GetVideoIDByAuthorID(authorID int64) (id []int64, err error) {
	videos, err := GetVideoByAuthorID(authorID)
	for _, video := range videos {
		id = append(id, video.AuthorID)
	}
	return id, err
}

func DeleteVideoByID(id int64) (err error) {
	v := query.Video
	_, err = v.Where(v.ID.Eq(id)).Delete()
	return err
}
