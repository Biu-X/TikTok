package common

import (
	"biu-x.org/TikTok/dao"
	"biu-x.org/TikTok/model"
	"biu-x.org/TikTok/module/log"
	"biu-x.org/TikTok/module/response"
	"errors"
	"gorm.io/gorm"
	"time"
)

func GetVideoList(targetID, ownerID int64, latestTime ...string) ([]response.VideoResponse, error) {
	var videoRespList []response.VideoResponse
	var videoList []*model.Video
	var err error
	if latestTime == nil {
		videoList, err = dao.GetVideoByAuthorID(targetID)
		if err != nil {
			log.Logger.Error(err)
			return nil, err
		}
	} else if latestTime[0] == "0" {
		now := time.Now()
		log.Logger.Debugf("now: %v", now)
		videoList, err = dao.GetVideoListByLatestTimeOrderByDESC(now)
		if err != nil {
			return nil, err
		}
	} else {
		latesttime, err := time.Parse(time.DateTime, latestTime[0])
		videoList, err = dao.GetVideoListByLatestTimeOrderByDESC(latesttime)
		if err != nil {
			return nil, err
		}
	}

	for _, video := range videoList {
		userResponse, err := response.GetUserResponseByID(video.AuthorID, ownerID)
		if err != nil {
			log.Logger.Error(err)
			if errors.As(err, &gorm.ErrRecordNotFound) {
				userResponse = &response.UserResponse{}
			} else {
				continue
			}
		}
		favoriteCount, err := dao.GetFavoriteCountByVideoID(video.ID)
		if err != nil {
			log.Logger.Error(err)
			if errors.As(err, &gorm.ErrRecordNotFound) {
				favoriteCount = 0
			} else {
				continue
			}
		}
		favorite, err := dao.GetFavoriteByBoth(ownerID, video.ID)
		if err != nil {
			log.Logger.Error(err)
			if errors.As(err, &gorm.ErrRecordNotFound) {
				favorite = &model.Favorite{}
			} else {
				continue
			}
		}
		count, err := dao.GetCommentCountByVideoID(video.ID)
		if err != nil {
			log.Logger.Error(err)
			if errors.As(err, &gorm.ErrRecordNotFound) {
				count = 0
			} else {
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

// 在没有登陆的情况下调用
func GetVideoListDefault() ([]response.VideoResponse, error) {
	now := time.Now()
	videos, err := dao.GetVideoListByLatestTimeOrderByDESC(now)
	if err != nil {
		log.Logger.Error(err)
		return nil, err
	}

	var videoRespList []response.VideoResponse
	for _, video := range videos {
		userResponse, err := response.GetUserResponseByUserId(video.AuthorID)
		if err != nil {
			log.Logger.Error(err)
			if errors.As(err, &gorm.ErrRecordNotFound) {
				userResponse = &response.UserResponse{}
			} else {
				continue
			}
		}
		favoriteCount, err := dao.GetFavoriteCountByVideoID(video.ID)
		if err != nil {
			log.Logger.Error(err)
			if errors.As(err, &gorm.ErrRecordNotFound) {
				favoriteCount = 0
			} else {
				continue
			}
		}
		count, err := dao.GetCommentCountByVideoID(video.ID)
		if err != nil {
			log.Logger.Error(err)
			if errors.As(err, &gorm.ErrRecordNotFound) {
				count = 0
			} else {
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
			IsFavorite:    false,
			Title:         video.Title,
		})
	}

	return videoRespList, nil
}
