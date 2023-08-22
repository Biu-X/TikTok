package relation

import (
	"biu-x.org/TikTok/dao"
	"biu-x.org/TikTok/module/log"
	"biu-x.org/TikTok/module/response"
	"biu-x.org/TikTok/module/util"
	"github.com/gin-gonic/gin"
)

// FriendList /douyin/relation/friend/list/ - 用户好友列表
func FriendList(c *gin.Context) {
	userID := util.GetUserIDFromGinContext(c)

	var userList []response.FriendUserResponse

	followerIDs, err := dao.GetFollowerIDsByUserID(userID)
	if err != nil {
		log.Logger.Error(err)
		response.ErrRespWithMsg(c, err.Error())
		return
	}

	for _, followerID := range followerIDs {
		userRes, err := response.GetUserResponseByID(followerID, userID)
		if err != nil {
			log.Logger.Error(err)
			continue
		}

		message, err := dao.GetLatestBidirectionalMessage(userID, followerID)
		if err != nil {
			// 第一次加好友时，没有消息可以获取，这里忽略错误
			log.Logger.Error(err)
		}

		var msgType int64
		if message.FromUserID == userID {
			msgType = 1
		} else {
			msgType = 0
		}

		userList = append(userList, response.FriendUserResponse{
			UserResponse: *userRes,
			Message:      message.Content,
			MsgType:      msgType,
		})
	}

	response.OKRespWithData(c, map[string]interface{}{
		"user_list": userList,
	})
}
