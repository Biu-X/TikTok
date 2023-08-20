package feed

import (
	"biu-x.org/TikTok/dao"
	"biu-x.org/TikTok/module/log"
	"biu-x.org/TikTok/module/response"
	"biu-x.org/TikTok/service/common"
	"github.com/gin-gonic/gin"
	"strconv"
)

func List(c *gin.Context) {
	targetID, _ := strconv.ParseInt(c.Query("user_id"), 10, 64)
	ownerID, _ := strconv.ParseInt(c.GetString("user_id"), 10, 64)
	latestTime := c.Query("latest_time")
	videoList, err := common.GetVideoList(targetID, ownerID, latestTime)
	if err != nil {
		log.Logger.Error(err)
		response.ErrRespWithMsg(c, err.Error())
		return
	}

	time, err := dao.GetVideoCreateTimeByID(videoList[len(videoList)-1].VideoID)
	if err != nil {
		log.Logger.Error(err)
	}
	response.OKRespWithData(c, map[string]interface{}{
		"next_time":  time,
		"video_list": videoList,
	})
}
