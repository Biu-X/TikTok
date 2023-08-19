package relation

import (
	"biu-x.org/TikTok/dao"
	"biu-x.org/TikTok/module/log"
	"biu-x.org/TikTok/module/response"
	"github.com/gin-gonic/gin"
	"strconv"
)

// FollowList /douyin/relatioin/follow/list/ - 用户关注列表
func FollowList(c *gin.Context) {
	// 从 RequireAuth 处读取 user_id
	userId, _ := strconv.ParseInt(c.GetString("user_id"), 10, 64)

	var userList []response.UserResponse

	followIDs, err := dao.GetFollowUserIDsByUserID(userId)
	if err != nil {
		response.ErrRespWithMsg(c, err.Error())
		return
	}

	for _, followID := range followIDs {
		userRes, err := response.GetUserResponseByID(followID, userId)
		if err != nil {
			log.Logger.Error(err)
			continue
		}
		userList = append(userList, *userRes)
	}

	response.OKRespWithData(c, map[string]interface{}{
		"user_list": userList,
	})
}
