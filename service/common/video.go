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
	if len(latestTime[0]) == 0 {
		videoList, err = dao.GetVideoByAuthorID(targetID)
		if err != nil {
			log.Logger.Error(err)
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
