package relation

import (
	"biu-x.org/TikTok/dao"
	"biu-x.org/TikTok/module/log"
	"biu-x.org/TikTok/module/response"
	"biu-x.org/TikTok/module/util"
	"github.com/gin-gonic/gin"
)

// FollowerList /douyin/relation/follower/list/ - 用户粉丝列表
func FollowerList(c *gin.Context) {
	userID := util.GetUserIDFromGinContext(c)

	var userList []response.UserResponse

	followerIDs, err := dao.GetFollowerIDsByUserID(userID)
	if err != nil {
		response.ErrRespWithMsg(c, err.Error())
		return
	}

	for _, followID := range followerIDs {
		userRes, err := response.GetUserResponseByID(followID, userID)
		if err != nil {
			log.Logger.Error(err)
			response.ErrRespWithMsg(c, err.Error())
			return
		}
		userList = append(userList, *userRes)
	}

	response.OKRespWithData(c, map[string]interface{}{
		"user_list": userList,
	})
}
