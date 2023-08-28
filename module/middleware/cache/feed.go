package cache

import (
	"github.com/Biu-X/TikTok/module/cache"
	"github.com/Biu-X/TikTok/module/log"
	"github.com/Biu-X/TikTok/module/response"
	"github.com/gin-gonic/gin"
	"github.com/pquerna/ffjson/ffjson"
	"time"
)

// FeedMiddleware feed 缓存装饰器
func FeedMiddleware(rc *cache.Client, service gin.HandlerFunc, empty interface{}) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 从连接池获取一个连接
		conn := rc.C.Conn()
		defer conn.Close()

		// latest time 作为 key
		latestTime := c.Query("latest_time")
		if latestTime == "" {
			latestTime = "0"
		}
		log.Logger.Infof("----------------> feed latest time: %v", latestTime)
		result, err := conn.Get(rc.Ctx, latestTime).Result()
		if err != nil {
			// 缓存没有存储该key
			log.Logger.Infof("get feed data from MySQL database, due to %v", err)
			service(c)
			ret, exists := c.Get("feed")
			if !exists {
				return
			}
			retData, _ := ffjson.Marshal(ret)
			conn.SetNX(rc.Ctx, latestTime, retData, 1*time.Hour)
			return
		}
		log.Logger.Info("get feed data from redis database")
		err = ffjson.Unmarshal([]byte(result), &empty)
		if err != nil {
			log.Logger.Errorf("unmarshal result error: %v", err)
			response.ErrRespWithMsg(c, "unmarshal result error")
			return
		}

		response.OKRespWithData(c, map[string]interface{}{
			"next_time":  time.Now().UnixMilli(),
			"video_list": empty,
		})
	}
}
