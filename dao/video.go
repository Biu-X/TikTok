package dao

import (
	"biu-x.org/TikTok/dal/query"
	"biu-x.org/TikTok/model"
)

// 创建视频记录
func CreateVideo(video *model.Video) (err error) {
	v := query.Video
	err = v.Create(video)
	return err
}

// 通过视频ID获取对应记录
func GetVideoByID(id int64) (video *model.Video, err error) {
	v := query.Video
	video, err = v.Where(v.ID.Eq(id)).First()
	return video, err
}

// 通过作者ID获取记录
func GetVideoByAuthorID(authorID int64) (videos []*model.Video, err error) {
	v := query.Video
	videos, err = v.Where(v.AuthorID.Eq(authorID)).Find()
	return videos, err
}

// 通过作者 ID 获取作品数量
func GetVideoCountByAuthorID(authorID int64) (count int64, err error) {
	v := query.Video
	count, err = v.Where(v.AuthorID.Eq(authorID)).Count()
	return count, err
}

// 通过作者ID获取视频ID表
func GetVideoIDByAuthorID(authorID int64) (id []int64, err error) {
	videos, err := GetVideoByAuthorID(authorID)
	for _, video := range videos {
		id = append(id, video.AuthorID)
	}
	return id, err
}

// 通过记录ID删除对应记录
func DeleteVideoByID(id int64) (err error) {
	v := query.Video
	_, err = v.Where(v.ID.Eq(id)).Delete()
	return err
}
