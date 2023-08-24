package message

import (
	"strconv"

	"github.com/Biu-X/TikTok/dal/model"
	"github.com/Biu-X/TikTok/dao"
	"github.com/Biu-X/TikTok/module/log"
	"github.com/Biu-X/TikTok/module/response"
	"github.com/Biu-X/TikTok/module/util"
	"github.com/gin-gonic/gin"
)

// Action /douyin/message/action/ - 消息操作
func Action(c *gin.Context) {
	userID := util.GetUserIDFromGinContext(c)
	toUserID, _ := strconv.ParseInt(c.Query("to_user_id"), 10, 64)
	actionType, _ := strconv.ParseInt(c.Query("action_type"), 10, 32) // 1-发送消息
	content := util.GetInsensitiveTextFromGinContext(c, "content")

	if content == "" {
		response.OKResp(c)
		return
	}

	if actionType == 1 {
		err := dao.CreateMessage(&model.Message{
			ToUserID:   toUserID,
			FromUserID: userID,
			Content:    content,
		})
		if err != nil {
			log.Logger.Errorf("action: CreateMessage failed, err: %v", err)
			response.ErrRespWithMsg(c, err.Error())
			return
		}
		response.OKResp(c)
	}
}
