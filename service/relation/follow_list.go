package relation

import (
	"github.com/Biu-X/TikTok/dao"
	"github.com/Biu-X/TikTok/module/log"
	"github.com/Biu-X/TikTok/module/response"
	"github.com/Biu-X/TikTok/module/util"
	"github.com/gin-gonic/gin"
)

// FollowList /douyin/relatioin/follow/list/ - 用户关注列表
func FollowList(c *gin.Context) {
	userID := util.GetUserIDFromGinContext(c)

	var userList []response.UserResponse

	followIDs, err := dao.GetFollowingIdsByUserID(userID)
	if err != nil {
		response.ErrRespWithMsg(c, err.Error())
		return
	}

	for _, followID := range followIDs {
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
