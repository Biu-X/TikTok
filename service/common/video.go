package common

import (
	"strconv"
	"time"

	"biu-x.org/TikTok/dal/model"
	"biu-x.org/TikTok/dao"
	"biu-x.org/TikTok/module/log"
	"biu-x.org/TikTok/module/response"
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
		// GetUserResponseByID 中返回的 error 不会涉及 ErrRecordNotFound
		userResponse, err := response.GetUserResponseByID(video.AuthorID, ownerID)
		if err != nil {
			userResponse = &response.UserResponse{}
			continue
		}

		favoriteCount, err := dao.GetFavoriteCountByVideoID(video.ID)
		log.Logger.Info("favoriteCount", favoriteCount)
		if err != nil {
			log.Logger.Error("favoriteCount query failed")
			favoriteCount = 0
			continue
		}

		isFavorite := dao.GetUserIsFavoriteVideo(ownerID, video.ID)
		log.Logger.Info("isFavorite: ", isFavorite)
		count, err := dao.GetCommentCountByVideoID(video.ID)
		if err != nil {
			log.Logger.Error(err)
			count = 0
			continue
		}

		videoRespList = append(videoRespList, response.VideoResponse{
			VideoID:       video.ID,
			Author:        *userResponse,
			PlayURL:       video.PlayURL,
			CoverURL:      video.CoverURL,
			FavoriteCount: favoriteCount,
			CommentCount:  count,
			IsFavorite:    isFavorite,
			Title:         video.Title,
		})
	}
	return videoRespList, nil
}
