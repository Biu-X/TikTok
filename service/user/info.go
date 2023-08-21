package user

import (
	"strconv"

	"biu-x.org/TikTok/module/response"
	"github.com/gin-gonic/gin"
)

// token 验证通过后，可以根据用户 id 查询用户的信息
func UserInfo(c *gin.Context) {
	idStr := c.GetString("user_id")
	id, _ := strconv.Atoi(idStr)
	userinfo, err := response.GetUserResponseByOwnerId(int64(id))
	if err != nil {
		response.ErrRespWithMsg(c, "User not found")
		return
	}

	response.OKRespWithData(c, map[string]interface{}{
		"user": *userinfo,
	})
}
