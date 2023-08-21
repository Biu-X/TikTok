package user

import (
	"biu-x.org/TikTok/module/response"
	"biu-x.org/TikTok/module/util"
	"github.com/gin-gonic/gin"
)

// Info /douyin/user/ - 用户信息
func Info(c *gin.Context) {
	userID := util.GetUserIDFromGinContext(c)

	userinfo, err := response.GetUserResponseByOwnerId(userID)
	if err != nil {
		response.ErrRespWithMsg(c, "User not found")
		return
	}

	response.OKRespWithData(c, map[string]interface{}{
		"user": *userinfo,
	})
}
