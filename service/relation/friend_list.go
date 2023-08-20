package relation

import (
	"biu-x.org/TikTok/dao"
	"biu-x.org/TikTok/module/log"
	"biu-x.org/TikTok/module/response"
	"github.com/gin-gonic/gin"
	"strconv"
)

// FriendList /douyin/relation/friend/list/ - 用户好友列表
func FriendList(c *gin.Context) {
	// 从 RequireAuth 处读取 user_id
	userId, _ := strconv.ParseInt(c.GetString("user_id"), 10, 64)

	var userList []response.FriendUserResponse

	followerIDs, err := dao.GetFollowerIDsByUserID(userId)
	if err != nil {
		log.Logger.Error(err)
		response.ErrRespWithMsg(c, err.Error())
		return
	}

	for _, followerID := range followerIDs {
		userRes, err := response.GetUserResponseByID(followerID, userId)
		if err != nil {
			log.Logger.Error(err)
			continue
		}

		message, err := dao.GetLatestBidirectionalMessage(userId, followerID)
		if err != nil {
			log.Logger.Error(err)
			continue
		}

		var msgType int64
		if message.FromUserID == userId {
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
