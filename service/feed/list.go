package feed

import (
	"strconv"
	"time"

	"github.com/Biu-X/TikTok/dao"
	"github.com/Biu-X/TikTok/module/log"
	"github.com/Biu-X/TikTok/module/response"
	"github.com/Biu-X/TikTok/module/util"
	"github.com/gin-gonic/gin"
)

// List /douyin/feed/ - 视频流接口
func List(c *gin.Context) {
	ownerID := util.GetUserIDFromGinContext(c)
	// 获取latest_time
	latest_time := c.Query("latest_time")
	// 如果latest_time为 0 或者为空，则设置为当前时间的UnixMilli()
	if latest_time == "0" || len(latest_time) == 0 {
		latest_time = strconv.FormatInt(time.Now().UnixMilli(), 10)
	}

	log.Logger.Infof("latest time: %v", latest_time)

	videoList, err := response.GetVideoListResponseByOwnerIDAndLatestTime(ownerID, latest_time)
	if err != nil {
		response.ErrRespWithMsg(c, err.Error())
		return
	}

	log.Logger.Infof("video list: %v", videoList)

	nextTime := ""
	length := len(videoList) - 1
	if length < 0 {
		length = 0
		nextTime = "0"
	}

	if length > 0 {
		log.Logger.Debugf("length: %v", length+1)
		t, err := dao.GetVideoCreateTimeByID(videoList[length].VideoID)
		if err != nil {
			log.Logger.Error(err)
			nextTime = "0"
		}
		nextTime = strconv.FormatInt(t.UnixMilli(), 10)
	}
	log.Logger.Debugf("nextTime: %v", nextTime)

	response.OKRespWithData(c, map[string]interface{}{
		"next_time":  nextTime,
		"video_list": videoList,
	})
}
