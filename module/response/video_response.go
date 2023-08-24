package response

import (
	"errors"
	"github.com/Biu-X/TikTok/dal/model"
	"github.com/Biu-X/TikTok/dao"
	"github.com/Biu-X/TikTok/module/config"
	"github.com/Biu-X/TikTok/module/log"
	"gorm.io/gorm"
	"strconv"
	"time"
)

// GetVideoResponseByVideoAndOwnerID 根据 video 和 ownerID 来获取视频
func GetVideoResponseByVideoAndOwnerID(video *model.Video, ownerID int64) (*VideoResponse, error) {
	// GetUserResponseByID 中返回的 error 不会涉及 ErrRecordNotFound
	userResponse, err := GetUserResponseByID(video.AuthorID, ownerID)
	if err != nil {
		userResponse = &UserResponse{}
		return nil, err
	}

	favoriteCount, err := dao.GetFavoriteCountByVideoID(video.ID)
	if err != nil {
		log.Logger.Error("favoriteCount query failed")
		return nil, err
	}

	isFavorite := dao.GetUserIsFavoriteVideo(ownerID, video.ID)
	log.Logger.Info("isFavorite: ", isFavorite)

	count, err := dao.GetCommentCountByVideoID(video.ID)
	if err != nil {
		log.Logger.Error(err)
		return nil, err
	}

	return &VideoResponse{
		VideoID:       video.ID,
		Author:        *userResponse,
		PlayURL:       config.OSS_PREFIX + video.PlayURL,
		CoverURL:      config.OSS_PREFIX + video.CoverURL,
		FavoriteCount: favoriteCount,
		CommentCount:  count,
		IsFavorite:    isFavorite,
		Title:         video.Title,
	}, nil
}

// GetVideoListResponseByOwnerIDAndLatestTime 根据 当前登录用户 id 和 时间戳 来获取视频列表
func GetVideoListResponseByOwnerIDAndLatestTime(ownerID int64, latestTime string) ([]VideoResponse, error) {
	return GetVideoListResponseByIDAndLatestTime(0, ownerID, latestTime)
}

// GetVideoListResponseByID 根据两个用户 id 来获取视频列表
func GetVideoListResponseByID(targetID, ownerID int64) ([]VideoResponse, error) {
	return GetVideoListResponseByIDAndLatestTime(targetID, ownerID, "")
}

// GetVideoListResponseByIDAndLatestTime 根据两个用户 id 来获取视频列表，如果 targetID 或者 ownerID 为 0，则通过时间戳来获取视频, 即未登录状态
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
		videoRes, err := GetVideoResponseByVideoAndOwnerID(video, ownerID)
		if err != nil {
			log.Logger.Error(err)
			continue
		}
		videoRespList = append(videoRespList, *videoRes)
	}
	return videoRespList, nil
}

// GetFavoriteVideoListResponseByOwnerID 根据用户 id 获取喜欢列表
func GetFavoriteVideoListResponseByOwnerID(ownerID int64) ([]VideoResponse, error) {
	// 根据 user_id 查询喜欢列表
	favorites, err := dao.GetFavoriteByUserID(ownerID)
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

		videoRes, err := GetVideoResponseByVideoAndOwnerID(video, ownerID)
		if err != nil {
			log.Logger.Error(err)
			continue
		}
		videoList = append(videoList, *videoRes)
	}

	return videoList, nil
}
