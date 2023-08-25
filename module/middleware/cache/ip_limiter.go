package cache

import (
	"fmt"
	"net/http"
	"time"

	"github.com/Biu-X/TikTok/module/cache"
	"github.com/Biu-X/TikTok/module/log"
	"github.com/Biu-X/TikTok/module/response"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

func NewRateLimiter(redisClient *cache.Client, key string, limit int, slidingWindow time.Duration) gin.HandlerFunc {
	_, err := redisClient.Ping().Result()
	if err != nil {
		panic(fmt.Sprint("error init redis", err.Error()))
	}

	return func(c *gin.Context) {
		now := time.Now().UnixNano()
		log.Logger.Infof("-------------------> path: %v", c.Request.URL.Path)
		userCntKey := fmt.Sprint(c.ClientIP(), ":", key, ":", c.Request.URL.Path)

		_, err := redisClient.ZRemRangeByScore(userCntKey,
			"0",
			fmt.Sprint(now-(slidingWindow.Nanoseconds()))).Result()
		if err != nil {
			log.Logger.Error(err)
			return
		}

		reqs, _ := redisClient.ZRange(userCntKey, 0, -1).Result()

		if len(reqs) >= limit {
			c.AbortWithStatusJSON(http.StatusOK, gin.H{
				"status_code": response.OK,
				"status_msg":  "too many request",
			})
			log.Logger.Warnf("------------------> too many request, key: %v", userCntKey)
			return
		}

		c.Next()
		redisClient.ZAddNX(userCntKey, redis.Z{Score: float64(now), Member: float64(now)})
		redisClient.Expire(userCntKey, slidingWindow)
	}
}
