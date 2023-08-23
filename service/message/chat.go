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

// Chat /douyin/message/chat/ - 聊天记录
func Chat(c *gin.Context) {
	userID := util.GetUserIDFromGinContext(c)
	toUserID, _ := strconv.ParseInt(c.Query("to_user_id"), 10, 64)
	preMsgTimeStamp, _ := strconv.ParseInt(c.Query("pre_msg_time"), 10, 64)

	// 保留精度
	seconds := preMsgTimeStamp / 1000
	milliseconds := preMsgTimeStamp % 1000
	preMsgTime := time.Unix(seconds, milliseconds*int64(time.Millisecond))

	messages, err := dao.GetMessageByBoth(userID, toUserID, preMsgTime)

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
