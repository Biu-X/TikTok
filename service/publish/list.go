package publish

import (
	"biu-x.org/TikTok/dao"
	"biu-x.org/TikTok/model"
	"biu-x.org/TikTok/module/log"
	"biu-x.org/TikTok/module/response"
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"strconv"
)

// List 获取发布列表 /douyin/publish/list/
func List(c *gin.Context) {
	// 获取某个用户的 id
	targetID, _ := strconv.ParseInt(c.Query("user_id"), 10, 64)
	log.Logger.Infof("target user id is: %v", targetID)
	// 自己的 id
	ownerID, _ := strconv.ParseInt(c.GetString("user_id"), 10, 64)
	log.Logger.Infof("owner id is: %v", ownerID)
	userResponse, err := response.GetUserResponseByID(targetID, ownerID)
	if err != nil {
		log.Logger.Error(err)
		if errors.As(err, &gorm.ErrRecordNotFound) {
			userResponse = &response.UserResponse{}
		} else {
			response.ErrRespWithMsg(c, err.Error())
			return
		}
	}
	var videoRespList []response.VideoResponse
	videoList, err := dao.GetVideoByAuthorID(targetID)
	if err != nil {
		log.Logger.Error(err)
		response.ErrRespWithMsg(c, err.Error())
		return
	}

	for _, video := range videoList {
		favoriteCount, err := dao.GetFavoriteCountByVideoID(video.ID)
		if err != nil {
			log.Logger.Error(err)
			if errors.As(err, &gorm.ErrRecordNotFound) {
				favoriteCount = 0
			} else {
				response.ErrRespWithMsg(c, err.Error())
				continue
			}
		}
		favorite, err := dao.GetFavoriteByBoth(ownerID, video.ID)
		if err != nil {
			log.Logger.Error(err)
			if errors.As(err, &gorm.ErrRecordNotFound) {
				favorite = &model.Favorite{}
			} else {
				response.ErrRespWithMsg(c, err.Error())
				continue
			}
		}
		count, err := dao.GetCommentCountByVideoID(video.ID)
		if err != nil {
			log.Logger.Error(err)
			if errors.As(err, &gorm.ErrRecordNotFound) {
				count = 0
			} else {
				response.ErrRespWithMsg(c, err.Error())
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
	response.OKRespWithData(c, map[string]interface{}{
		"video_list": videoRespList,
	})
}
