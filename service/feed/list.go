package feed

import (
	"biu-x.org/TikTok/dao"
	"biu-x.org/TikTok/module/log"
	"biu-x.org/TikTok/module/response"
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
)

func List(c *gin.Context) {
	// 获取latest_time
	latest_time := c.Query("latest_time")
	// 如果latest_time为 0 或者为空，则设置为当前时间的UnixMilli()
	if latest_time == "0" || len(latest_time) == 0 {
		latest_time = strconv.FormatInt(time.Now().UnixMilli(), 10)
	}
	// 判断是否登陆
	_, ok := c.Get("is_login")
	log.Logger.Infof("current login status: %v", ok)

	var ownerID int64
	if !ok {
		// 未登录
		ownerID = 0
	} else {
		// 已登录
		// 从RequireAuth获取当前登陆的用户 id
		ownerID, _ = strconv.ParseInt(c.GetString("user_id"), 10, 64)
	}
	log.Logger.Infof("owner id: %v", ownerID)

	videoList, err := response.GetVideoListResponseByUserIDAndLatestTime(ownerID, latest_time)
	if err != nil {
		response.ErrRespWithMsg(c, err.Error())
		return
	}

	log.Logger.Infof("video list: %v", *videoList)

	nextTime := ""
	length := len(*videoList) - 1
	if length < 0 {
		length = 0
		nextTime = "0"
	}

	if length > 0 {
		log.Logger.Debugf("length: %v", length)
		t, err := dao.GetVideoCreateTimeByID((*videoList)[length].VideoID)
		if err != nil {
			log.Logger.Error(err)
			nextTime = "0"
		}
		nextTime = strconv.FormatInt(t.UnixMilli(), 10)
	}
	log.Logger.Debugf("nextTime: %v", nextTime)

	response.OKRespWithData(c, map[string]interface{}{
		"next_time":  nextTime,
		"video_list": *videoList,
	})
}
