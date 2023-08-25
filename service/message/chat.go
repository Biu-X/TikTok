package message

import (
	"strconv"
	"time"

	"github.com/Biu-X/TikTok/dal/model"
	"github.com/Biu-X/TikTok/dao"
	"github.com/Biu-X/TikTok/module/log"
	"github.com/Biu-X/TikTok/module/response"
	"github.com/Biu-X/TikTok/module/util"
	"github.com/gin-gonic/gin"
)

// Chat /douyin/message/chat/ - 聊天记录
func Chat(c *gin.Context) {
	userID := util.GetUserIDFromGinContext(c)
	toUserID, _ := strconv.ParseInt(c.Query("to_user_id"), 10, 64)
	preMsgTimeStamp, _ := strconv.ParseInt(c.Query("pre_msg_time"), 10, 64)

	message_list := []response.MessageResponse{}
	messages := []*model.Message{}
	var err error

	if preMsgTimeStamp == 0 {
		// 返回全部聊天记录
		messages, err = dao.GetMessageByBoth(userID, toUserID, time.Unix(0, 0))
		if err != nil {
			log.Logger.Errorf("chat: GetMessageByBoth failed, err: %v", err)
			response.ErrRespWithMsg(c, err.Error())
			return
		}
	} else {
		// 返回对方之后传来的消息
		// 保留精度
		seconds := preMsgTimeStamp / 1000
		milliseconds := preMsgTimeStamp % 1000
		preMsgTime := time.Unix(seconds, milliseconds*int64(time.Millisecond))

		messages, err = dao.GetUserMessagesToUser(toUserID, userID, preMsgTime)
		if err != nil {
			log.Logger.Errorf("chat: GetMessageByBoth failed, err: %v", err)
			response.ErrRespWithMsg(c, err.Error())
			return
		}
	}

	log.Logger.Infof("------------> pre_msg_time: %v", preMsgTimeStamp)

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
