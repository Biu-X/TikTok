package relation

import (
	"strconv"

	"github.com/Biu-X/TikTok/dao"
	"github.com/Biu-X/TikTok/module/log"
	"github.com/Biu-X/TikTok/module/response"
	"github.com/Biu-X/TikTok/module/util"
	"github.com/gin-gonic/gin"
)

// Action /douyin/relation/action/ - 关系操作
func Action(c *gin.Context) {
	userID := util.GetUserIDFromGinContext(c)
	// 从 request 中查询
	toUserID, _ := strconv.ParseInt(c.Query("to_user_id"), 10, 64)    // 对方用户id
	actionType, _ := strconv.ParseInt(c.Query("action_type"), 10, 64) // 1-关注，2-取消关注

	if userID == toUserID {
		response.ErrRespWithMsg(c, "can't follow yourself")
		return
	}

	// 根据 action_type 执行不同的操作
	if actionType == 1 {
		err := dao.SetFollowingByBoth(toUserID, userID)
		if err != nil {
			log.Logger.Error(err.Error())
			response.ErrRespWithMsg(c, err.Error())
			return
		}
	} else {
		err := dao.SetFollowCancelByBoth(toUserID, userID)
		if err != nil {
			log.Logger.Error(err.Error())
			response.ErrRespWithMsg(c, err.Error())
			return
		}
	}

	response.OKResp(c)
}
