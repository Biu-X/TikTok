package message

import (
	"strconv"
	"time"

	"biu-x.org/TikTok/dao"
	"biu-x.org/TikTok/module/log"
	"biu-x.org/TikTok/module/response"
	"biu-x.org/TikTok/module/util"
	"github.com/gin-gonic/gin"
)

var tmpTimeStamp int64

// Chat /douyin/message/chat/ - 聊天记录
func Chat(c *gin.Context) {
	userID := util.GetUserIDFromGinContext(c)
	toUserID, _ := strconv.ParseInt(c.Query("to_user_id"), 10, 64)
	preMsgTimeStamp, _ := strconv.ParseInt(c.Query("pre_msg_time"), 10, 64)

	if preMsgTimeStamp == 0 {
		tmp, err := dao.GetEarliestTimeMessageByBoth(userID, toUserID)
		if err != nil {
			response.ErrResp(c)
			return
		}
		preMsgTimeStamp = tmp.UnixMilli()
	}

	if tmpTimeStamp == preMsgTimeStamp {
		log.Logger.Infof("------------> pre_msg_time: %v", preMsgTimeStamp)
		response.OKResp(c)
		return
	}
	tmpTimeStamp = preMsgTimeStamp

	preMsgTime := time.Unix(preMsgTimeStamp/1000, 0)

	messages, err := dao.GetMessageByBoth(userID, toUserID, preMsgTime)

	log.Logger.Infof("------> ownerID: %v, targetID: %v, pre_msg_time: %v", userID, toUserID, preMsgTimeStamp)

	if err != nil {
		log.Logger.Errorf("chat: GetMessageByBoth failed, err: %v", err)
		response.ErrRespWithMsg(c, err.Error())
		return
	}
	message_list := []response.MessageResponse{}
	for _, message := range messages {
		message_list = append(message_list, response.MessageResponse{
			ID:         message.ID,
			ToUserID:   message.ToUserID,
			FromUserID: message.FromUserID,
			Content:    message.Content,
			CreateTime: strconv.FormatInt(message.CreatedAt.UnixMilli(), 10),
		})
	}

	response.OKRespWithData(c, map[string]interface{}{
		"message_list": message_list,
	})

}
