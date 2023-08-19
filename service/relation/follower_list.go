package relation

import (
	"strconv"

	"biu-x.org/TikTok/dao"
	"biu-x.org/TikTok/module/log"
	"biu-x.org/TikTok/module/response"
	"github.com/gin-gonic/gin"
)

// FollowerList /douyin/relation/follower/list/ - 用户粉丝列表
func FollowerList(c *gin.Context) {
	// 从 RequireAuth 处读取 user_id
	userId, _ := strconv.ParseInt(c.GetString("user_id"), 10, 64)

	var userList []response.UserResponse

	followerIDs, err := dao.GetFollowerIDsByUserID(userId)
	if err != nil {
		response.ErrRespWithMsg(c, err.Error())
		return
	}

	for _, followID := range followerIDs {
		userRes, err := response.GetUserResponseByID(followID, userId)
		if err != nil {
			log.Logger.Error(err)
			response.ErrRespWithMsg(c, err.Error())
			return
		}
		log.Logger.Error(err)
		userList = append(userList, *userRes)
	}

	response.OKRespWithData(c, map[string]interface{}{
		"user_list": userList,
	})
}
