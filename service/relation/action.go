package relation

import (
	"biu-x.org/TikTok/dao"
	"biu-x.org/TikTok/module/log"
	"biu-x.org/TikTok/module/response"
	"github.com/gin-gonic/gin"
	"strconv"
)

// Action /douyin/relation/action/ - 关系操作
func Action(c *gin.Context) {
	// 从 RequireAuth 处读取 user_id
	userId, _ := strconv.ParseInt(c.GetString("user_id"), 10, 64)
	// 从 request 中查询
	toUserId, _ := strconv.ParseInt(c.Query("to_user_id"), 10, 64)    // 对方用户id
	actionType, _ := strconv.ParseInt(c.Query("action_type"), 10, 64) // 1-关注，2-取消关注

	// 根据 action_type 执行不同的操作
	if actionType == 1 {
		err := dao.SetFollowingByBoth(toUserId, userId)
		if err != nil {
			log.Logger.Error(err.Error())
			response.ErrRespWithMsg(c, err.Error())
			return
		}
	} else {
		err := dao.SetFollowCancelByBoth(toUserId, userId)
		if err != nil {
			log.Logger.Error(err.Error())
			response.ErrRespWithMsg(c, err.Error())
			return
		}
	}

	response.OKResp(c)
}
