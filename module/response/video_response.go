package response

import (
	"biu-x.org/TikTok/dal/model"
	"biu-x.org/TikTok/dao"
	"biu-x.org/TikTok/module/log"
	"errors"
	"gorm.io/gorm"
	"strconv"
	"time"
)

// GetVideoResponseByVideoAndUserID 根据 video 和 userID 来获取视频
func GetVideoResponseByVideoAndUserID(video *model.Video, userID int64) (*VideoResponse, error) {
	// GetUserResponseByID 中返回的 error 不会涉及 ErrRecordNotFound
	userResponse, err := GetUserResponseByID(video.AuthorID, userID)
	if err != nil {
		userResponse = &UserResponse{}
		return nil, err
	}

	favoriteCount, err := dao.GetFavoriteCountByVideoID(video.ID)
	if err != nil {
		log.Logger.Error("favoriteCount query failed")
		return nil, err
	}

	isFavorite := dao.GetUserIsFavoriteVideo(userID, video.ID)
	log.Logger.Info("isFavorite: ", isFavorite)

	count, err := dao.GetCommentCountByVideoID(video.ID)
	if err != nil {
		log.Logger.Error(err)
		return nil, err
	}

	return &VideoResponse{
		VideoID:       video.ID,
		Author:        *userResponse,
		PlayURL:       video.PlayURL,
		CoverURL:      video.CoverURL,
		FavoriteCount: favoriteCount,
		CommentCount:  count,
		IsFavorite:    isFavorite,
		Title:         video.Title,
	}, nil
}

// GetVideoListResponseByUserIDAndLatestTime 根据 当前登录用户 id 和 时间戳 来获取视频
func GetVideoListResponseByUserIDAndLatestTime(userID int64, latestTime string) ([]VideoResponse, error) {
	return GetVideoListResponseByIDAndLatestTime(0, userID, latestTime)
}

// GetVideoListResponseByID 根据 当前登录用户 id 和 时间戳 来获取视频
func GetVideoListResponseByID(targetID, ownerID int64) ([]VideoResponse, error) {
	return GetVideoListResponseByIDAndLatestTime(targetID, ownerID, "")
}

// GetVideoListResponseByIDAndLatestTime 根据两个用户 id 获取视频列表，如果 targetID 或者 ownerID 为 0，则通过时间戳来获取视频, 即未登录状态
func GetVideoListResponseByIDAndLatestTime(targetID, ownerID int64, latestTime string) ([]VideoResponse, error) {
	var videoRespList []VideoResponse
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
		videoRes, err := GetVideoResponseByVideoAndUserID(video, ownerID)
		if err != nil {
			log.Logger.Error(err)
			continue
		}
		videoRespList = append(videoRespList, *videoRes)
	}
	return videoRespList, nil
}

// GetFavoriteVideoListResponseByUserID 根据用户 id 获取喜欢列表
func GetFavoriteVideoListResponseByUserID(userID int64) ([]VideoResponse, error) {
	// 根据 user_id 查询喜欢列表
	favorites, err := dao.GetFavoriteByUserID(userID)
	if err != nil {
		return nil, err
	}

	var videoList []VideoResponse

	for _, favorite := range favorites {
		video, err := dao.GetVideoByID(favorite.VideoID) // 根据 video_id 查询视频信息
		if errors.Is(err, gorm.ErrRecordNotFound) {      // 如果没查询到该 videoId 的记录，那么跳过这个即可
			continue
		} else if err != nil {
			log.Logger.Error(err)
			return nil, err
		}

		videoRes, err := GetVideoResponseByVideoAndUserID(video, userID)
		if err != nil {
			log.Logger.Error(err)
			continue
		}

		videoList = append(videoList, *videoRes)
	}

	return videoList, nil
}
