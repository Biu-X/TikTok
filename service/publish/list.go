package publish

import (
	"biu-x.org/TikTok/module/log"
	"biu-x.org/TikTok/module/response"
	"github.com/gin-gonic/gin"
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
	videoRespList, err := response.GetVideoListResponseByIDAndLatestTime(targetID, ownerID, "")
	if err != nil {
		log.Logger.Error(err)
		response.ErrRespWithMsg(c, err.Error())
		return
	}
	response.OKRespWithData(c, map[string]interface{}{
		"video_list": videoRespList,
	})
}
