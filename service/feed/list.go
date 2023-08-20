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
	id, ok := c.Get("user_id")
	log.Logger.Infof("user id: %v,login: %v", id, ok)
	if _, ok := c.Get("user_id"); !ok {
		videoListDefault, err := common.GetVideoListDefault()
		log.Logger.Debugf("videoListDefault: %v", videoListDefault)
		if err != nil {
			log.Logger.Error(err)
			response.ErrRespWithMsg(c, err.Error())
			return
		}

		length := len(videoListDefault) - 1
		var ms = time.Now().UnixMilli()
		if length > 0 {
			t, err := dao.GetVideoCreateTimeByID(videoListDefault[length].VideoID)
			if err != nil {
				ms = time.Now().UnixMilli()
			}
			ms = t.UnixMilli()
		}

		log.Logger.Debugf("return!")
		response.OKRespWithData(c, map[string]interface{}{
			"next_time":  ms,
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

	ms := time.Now().UnixMilli()
	length := len(videoList) - 1
	if length > 0 {
		lt, err := dao.GetVideoCreateTimeByID(videoList[length].VideoID)
		if err != nil {
			log.Logger.Error(err)
		}
		ms = lt.UnixMilli()
	}
	log.Logger.Infof("time: %v", ms)
	if err != nil {
		log.Logger.Error(err)
	}
	response.OKRespWithData(c, map[string]interface{}{
		"next_time":  strconv.FormatInt(ms, 10),
		"video_list": videoList,
	})
}
