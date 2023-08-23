package message

import (
	"errors"
	"strconv"

	"biu-x.org/TikTok/dao"
	"biu-x.org/TikTok/module/log"
	"biu-x.org/TikTok/module/response"
	"biu-x.org/TikTok/module/util"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// Chat /douyin/message/chat/ - 聊天记录
func Chat(c *gin.Context) {
	userID := util.GetUserIDFromGinContext(c)
	toUserID, _ := strconv.ParseInt(c.Query("to_user_id"), 10, 64)

	messages, err := dao.GetMessageByBoth(userID, toUserID)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		response.OKRespWithData(c, map[string]interface{}{
			"message_list": messages,
		})
		return
	}

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
			CreateTime: message.CreatedAt.Format("2006-01-02 15:04:05"),
		})
	}

	response.OKRespWithData(c, map[string]interface{}{
		"message_list": message_list,
	})
}
