package relation

import (
	"biu-x.org/TikTok/dao"
	"biu-x.org/TikTok/module/log"
	"biu-x.org/TikTok/module/response"
	"github.com/gin-gonic/gin"
	"strconv"
)

// FollowerList /douyin/relation/follower/list/ - 用户粉丝列表
func FollowerList(c *gin.Context) {
	// 从 RequireAuth 处读取 user_id
	userId, _ := strconv.ParseInt(c.GetString("user_id"), 10, 64)

	var userList []response.UserResponse

	followerIDs, err := dao.GetFollowFollowerIDsByUserID(userId)
	if err != nil {
		response.ErrRespWithMsg(c, err.Error())
		return
	}

	for _, followerID := range followerIDs {
		userRes, err := response.GetUserResponseByID(followerID, userId)
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
