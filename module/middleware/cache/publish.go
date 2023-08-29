package cache

import (
	"github.com/Biu-X/TikTok/module/cache"

	"github.com/gin-gonic/gin"
)

func PublishMiddleware(rc *cache.Client, service gin.HandlerFunc, empty interface{}) gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}
