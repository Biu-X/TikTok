package favorite

import (
	"strconv"

	"github.com/Biu-X/TikTok/dao"
	"github.com/Biu-X/TikTok/module/response"
	"github.com/Biu-X/TikTok/module/util"
	"github.com/gin-gonic/gin"
)

// Action /douyin/favorite/action/ - 赞操作
func Action(c *gin.Context) {
	userID := util.GetUserIDFromGinContext(c)
	// 从 request 中查询 video_id 和 action_type
	videoID, _ := strconv.ParseInt(c.Query("video_id"), 10, 64)       // 视频id
	actionType, _ := strconv.ParseInt(c.Query("action_type"), 10, 64) // 1-点赞，2-取消点赞

	// 根据 action_type 执行不同的操作
	if actionType == 1 {
		err := dao.SetFavoriteByUserIDAndVideoID(userID, videoID)
		if err != nil {
			response.ErrRespWithMsg(c, err.Error())
			return
		}
	} else {
		err := dao.SetFavoriteCancelByUserIDAndVideoID(userID, videoID)
		if err != nil {
			response.ErrRespWithMsg(c, err.Error())
			return
		}
	}

	response.OKResp(c)
}
