package publish

import (
	"github.com/Biu-X/TikTok/module/log"
	"github.com/Biu-X/TikTok/module/response"
	"github.com/Biu-X/TikTok/module/util"
	"github.com/gin-gonic/gin"
	"strconv"
)

// List /douyin/publish/list/ - 发布列表
func List(c *gin.Context) {
	ownerID := util.GetUserIDFromGinContext(c)
	// 获取某个用户的 id
	targetID, _ := strconv.ParseInt(c.Query("user_id"), 10, 64)
	log.Logger.Infof("target user id is: %v", targetID)

	videoRespList, err := response.GetVideoListResponseByID(targetID, ownerID)
	if err != nil {
		log.Logger.Error(err)
		response.ErrRespWithMsg(c, err.Error())
		return
	}
	response.OKRespWithData(c, map[string]interface{}{
		"video_list": videoRespList,
	})
}
