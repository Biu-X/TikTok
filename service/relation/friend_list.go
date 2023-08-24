package relation

import (
	"github.com/Biu-X/TikTok/dao"
	"github.com/Biu-X/TikTok/module/log"
	"github.com/Biu-X/TikTok/module/response"
	"github.com/Biu-X/TikTok/module/util"
	"github.com/gin-gonic/gin"
)

// FriendList /douyin/relation/friend/list/ - 用户好友列表
func FriendList(c *gin.Context) {
	ownerID := util.GetUserIDFromGinContext(c)

	var userList []response.FriendUserResponse

	followerIDs, err := dao.GetFollowerIDsByUserID(ownerID)
	if err != nil {
		log.Logger.Error(err)
		response.ErrRespWithMsg(c, err.Error())
		return
	}

	for _, followerID := range followerIDs {
		userRes, err := response.GetUserResponseByID(followerID, ownerID)
		if err != nil {
			log.Logger.Error(err)
			continue
		}

		message, err := dao.GetLatestBidirectionalMessage(ownerID, followerID)
		if err != nil {
			// 第一次加好友时，没有消息可以获取，这里忽略错误
			log.Logger.Error(err)
		}

		var msgType int64
		if message.FromUserID == ownerID {
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
