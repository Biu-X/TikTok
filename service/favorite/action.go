package favorite

import (
	"strconv"

	"biu-x.org/TikTok/module/log"

	"biu-x.org/TikTok/dal/model"
	"biu-x.org/TikTok/dao"
	"biu-x.org/TikTok/module/response"
	"github.com/gin-gonic/gin"
)

// Action /douyin/favorite/action/ - 赞操作
func Action(c *gin.Context) {
	// 从 RequireAuth 处读取 user_id
	userId, _ := strconv.ParseInt(c.GetString("user_id"), 10, 64)
	// 从 request 中查询 video_id 和 action_type
	videoId, _ := strconv.ParseInt(c.Query("video_id"), 10, 64)       // 视频id
	actionType, _ := strconv.ParseInt(c.Query("action_type"), 10, 64) // 1-点赞，2-取消点赞

	// 根据 action_type 执行不同的操作
	if actionType == 1 {
		isLove := dao.GetUserIsFavoriteVideo(userId, videoId)
		if isLove {
			log.Logger.Info("用户已经点赞过该视频")
			return
		}
		err := dao.CreateFavorite(&model.Favorite{
			UserID:  userId,
			VideoID: videoId,
		})
		if err != nil {
			response.ErrRespWithMsg(c, err.Error())
			return
		}
	} else {
		err := dao.SetFavoriteCancelByVideoID(videoId, 1)
		if err != nil {
			response.ErrRespWithMsg(c, err.Error())
			return
		}
	}

	response.OKResp(c)
}
