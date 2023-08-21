package util

import (
	"biu-x.org/TikTok/module/log"
	"github.com/gin-gonic/gin"
	"strconv"
)

// GetUserIDFromGinContext 从 RequireAuth 处读取 user_id
func GetUserIDFromGinContext(c *gin.Context) int64 {
	// 判断是否登陆
	if _, ok := c.Get("is_login"); !ok {
		log.Logger.Infof("current login status: %v", ok)
		return 0
	}

	// 已登录
	userID, err := strconv.ParseInt(c.GetString("user_id"), 10, 64)
	if err != nil {
		log.Logger.Errorf("strconv.ParseInt failed, err: %v", err)
		return 0
	}

	return userID
}
