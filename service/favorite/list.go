package favorite

import (
	"strconv"

	"biu-x.org/TikTok/module/response"
	"github.com/gin-gonic/gin"
)

// List /douyin/favorite/list/ - 喜欢列表
func List(c *gin.Context) {
	// 从 RequireAuth 处读取 user_id
	userId, _ := strconv.ParseInt(c.GetString("user_id"), 10, 64)

	videoList, err := response.GetFavoriteVideoListResponseByOwnerID(userId)
	if err != nil {
		response.ErrRespWithMsg(c, err.Error())
		return
	}

	response.OKRespWithData(c, map[string]interface{}{
		"video_list": videoList,
	})
}
