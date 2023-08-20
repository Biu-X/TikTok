package feed

import (
	"biu-x.org/TikTok/dao"
	"biu-x.org/TikTok/module/log"
	"biu-x.org/TikTok/module/response"
	"biu-x.org/TikTok/service/common"
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
)

func List(c *gin.Context) {
	if _, ok := c.Get("user_id"); !ok {
		videoListDefault, err := common.GetVideoListDefault()
		if err != nil {
			log.Logger.Error(err)
			response.ErrRespWithMsg(c, err.Error())
			return
		}

		length := len(videoListDefault) - 1
		var t = time.Now()
		if length > 0 {
			t, err = dao.GetVideoCreateTimeByID(videoListDefault[length].VideoID)
			if err != nil {
				t = time.Now()
			}
		}

		response.OKRespWithData(c, map[string]interface{}{
			"next_time":  t.Format(time.DateTime),
			"video_list": videoListDefault,
		})
		return
	}
	targetID, _ := strconv.ParseInt(c.Query("user_id"), 10, 64)
	ownerID, _ := strconv.ParseInt(c.GetString("user_id"), 10, 64)
	latestTime := c.Query("latest_time")
	videoList, err := common.GetVideoList(targetID, ownerID, latestTime)
	if err != nil {
		log.Logger.Error(err)
		response.ErrRespWithMsg(c, err.Error())
		return
	}

	var t string
	length := len(videoList) - 1
	if length < 0 {
		t = time.Now().Format(time.DateTime)
	} else {
		lt, err := dao.GetVideoCreateTimeByID(videoList[length].VideoID)
		if err != nil {
			log.Logger.Error(err)
			t = time.Now().Format(time.DateTime)
		}
		t = lt.Format(time.DateTime)
	}
	log.Logger.Infof("time: %v", t)
	if err != nil {
		log.Logger.Error(err)
	}
	response.OKRespWithData(c, map[string]interface{}{
		"next_time":  t,
		"video_list": videoList,
	})
}
