package message

import (
	"strconv"

	"biu-x.org/TikTok/module/log"

	"biu-x.org/TikTok/dao"
	"biu-x.org/TikTok/model"
	"biu-x.org/TikTok/module/response"
	"github.com/gin-gonic/gin"
)

func Action(c *gin.Context) {
	userID, _ := strconv.ParseInt(c.GetString("user_id"), 10, 64)
	toUserID, _ := strconv.ParseInt(c.Query("to_user_id"), 10, 64)
	actionType, _ := strconv.ParseInt(c.Query("action_type"), 10, 32)
	content := c.Query("content")
	log.Logger.Info("action_type = ", actionType)
	if actionType == 1 {
		log.Logger.Info("action_type  1")
		err := dao.CreateMessage(&model.Message{
			ToUserID:   toUserID,
			FromUserID: userID,
			Content:    content,
		})
		if err != nil {
			response.ErrRespWithMsg(c, err.Error())
			return
		}
		response.OKResp(c)
	}
}
