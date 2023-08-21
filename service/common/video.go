package common

import (
	"biu-x.org/TikTok/dao"
	"biu-x.org/TikTok/dal/model"
	"biu-x.org/TikTok/module/log"
	"biu-x.org/TikTok/module/response"
	"errors"
	"gorm.io/gorm"
	"strconv"
	"time"
)

func GetVideoList(targetID, ownerID int64, latestTime string) ([]response.VideoResponse, error) {
	var videoRespList []response.VideoResponse
	var videoList []*model.Video
	var err error
	var latest_time int64

	if len(latestTime) == 0 {
		latest_time = 0
	} else {
		latest_time, _ = strconv.ParseInt(latestTime, 10, 64)
	}
	latest_time_ms := time.UnixMilli(latest_time)

	if targetID == 0 || ownerID == 0 { // 如果 targetID 或者 ownerID 为 0，则通过时间戳来获取视频, 即未登录状态
		videoList, err = dao.GetVideosByLatestTimeOrderByDESC(latest_time_ms)
		if err != nil {
			log.Logger.Error(err)
			return nil, err
		}
	} else if targetID > 0 && targetID != ownerID { // 查别人的视频列表
		videoList, err = dao.GetVideosByAuthorID(targetID)
		if err != nil {
			return nil, err
		}
	} else { // 查自己的视频列表
		videoList, err = dao.GetVideosByAuthorID(ownerID)
		if err != nil {
			log.Logger.Error(err)
			return nil, err
		}
	}

	for _, video := range videoList {
		userResponse, err := response.GetUserResponseByID(video.AuthorID, ownerID)
		if err != nil {
			if errors.As(err, &gorm.ErrRecordNotFound) {
				log.Logger.Info(gorm.ErrRecordNotFound)
				userResponse = &response.UserResponse{}
			} else {
				log.Logger.Error(err)
				continue
			}
		}
		favoriteCount, err := dao.GetFavoriteCountByVideoID(video.ID)
		if err != nil {
			if errors.As(err, &gorm.ErrRecordNotFound) {
				log.Logger.Info(gorm.ErrRecordNotFound)
				favoriteCount = 0
			} else {
				log.Logger.Error(err)
				continue
			}
		}
		favorite, err := dao.GetFavoriteByBoth(ownerID, video.ID)
		if err != nil {
			if errors.As(err, &gorm.ErrRecordNotFound) {
				log.Logger.Info(gorm.ErrRecordNotFound)
				favorite = &model.Favorite{}
			} else {
				log.Logger.Error(err)
				continue
			}
		}
		count, err := dao.GetCommentCountByVideoID(video.ID)
		if err != nil {
			log.Logger.Error(err)
			if errors.As(err, &gorm.ErrRecordNotFound) {
				log.Logger.Info(gorm.ErrRecordNotFound)
				count = 0
			} else {
				log.Logger.Error(err)
				continue
			}
		}
		videoRespList = append(videoRespList, response.VideoResponse{
			VideoID:       video.ID,
			Author:        *userResponse,
			PlayURL:       video.PlayURL,
			CoverURL:      video.CoverURL,
			FavoriteCount: favoriteCount,
			CommentCount:  count,
			IsFavorite:    favorite.Cancel == 0,
			Title:         video.Title,
		})
	}
	return videoRespList, nil
}
