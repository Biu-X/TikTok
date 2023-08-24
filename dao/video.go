package dao

import (
	"time"

	"github.com/Biu-X/TikTok/dal/model"
	"github.com/Biu-X/TikTok/dal/query"
	"github.com/Biu-X/TikTok/module/log"
	"gorm.io/gorm"
)

// CreateVideo 创建视频记录
func CreateVideo(video *model.Video) (err error) {
	v := query.Video
	err = v.Create(video)
	return err
}

// GetVideoByID 通过视频ID获取对应记录
func GetVideoByID(id int64) (video *model.Video, err error) {
	v := query.Video

	count, _ := v.Where(v.ID.Eq(id)).Count()
	if count == 0 {
		return &model.Video{}, gorm.ErrRecordNotFound
	}

	video, err = v.Where(v.ID.Eq(id)).First()
	return video, err
}

// GetVideoByAuthorID 通过作者ID获取记录
func GetVideoByAuthorID(authorID int64) (videos []*model.Video, err error) {
	v := query.Video
	videos, err = v.Where(v.AuthorID.Eq(authorID)).Find()
	return videos, err
}

// GetVideoCountByAuthorID 通过作者 ID 获取作品数量
func GetVideoCountByAuthorID(authorID int64) (count int64, err error) {
	v := query.Video
	count, err = v.Where(v.AuthorID.Eq(authorID)).Count()
	return count, err
}

// GetVideoIDByAuthorID 通过作者ID获取视频ID表
func GetVideoIDByAuthorID(authorID int64) (id []int64, err error) {
	videos, err := GetVideoByAuthorID(authorID)
	for _, video := range videos {
		id = append(id, video.AuthorID)
	}
	return id, err
}

// DeleteVideoByID 通过记录ID删除对应记录
func DeleteVideoByID(id int64) (err error) {
	v := query.Video
	_, err = v.Where(v.ID.Eq(id)).Delete()
	return err
}

// GetVideosByLatestTimeOrderByDESC 通过时间点来获取比该时间点早的十个视频
func GetVideosByLatestTimeOrderByDESC(latestTime time.Time) ([]*model.Video, error) {
	v := query.Video
	videos, err := v.Where(v.CreatedAt.Lt(latestTime)).Order(v.CreatedAt.Desc()).Limit(10).Find()
	if err != nil {
		log.Logger.Error(err)
		return nil, err
	}
	return videos, nil
}

// GetVideosByAuthorIDAnTimeOrderByDESC 通过 AuthorID 和时间点来获取比该时间点早的十个视频
func GetVideosByAuthorIDAnTimeOrderByDESC(authorID int64, latestTime time.Time) ([]*model.Video, error) {
	v := query.Video
	videos, err := v.Where(v.AuthorID.Eq(authorID), v.CreatedAt.Lt(latestTime)).Limit(10).Find()
	if err != nil {
		log.Logger.Error(err)
		return nil, err
	}
	return videos, nil
}

// GetVideosByAuthorID 通过作者 ID 获取视频列表
func GetVideosByAuthorID(authorID int64) ([]*model.Video, error) {
	v := query.Video
	videos, err := v.Where(v.AuthorID.Eq(authorID)).Find()
	if err != nil {
		log.Logger.Error(err)
		return nil, err
	}
	return videos, nil
}

// GetVideoCreateTimeByID 通过视频 ID 获取视频的创建时间
func GetVideoCreateTimeByID(id int64) (time.Time, error) {
	v := query.Video
	video, err := v.Where(v.ID.Eq(id)).First()
	if err != nil {
		return time.Time{}, err
	}
	return video.CreatedAt, nil
}
