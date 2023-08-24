package favorite

import (
	"github.com/Biu-X/TikTok/module/response"
	"github.com/Biu-X/TikTok/module/util"
	"github.com/gin-gonic/gin"
)

// List /douyin/favorite/list/ - 喜欢列表
func List(c *gin.Context) {
	userID := util.GetUserIDFromGinContext(c)

	videoList, err := response.GetFavoriteVideoListResponseByOwnerID(userID)
	if err != nil {
		response.ErrRespWithMsg(c, err.Error())
		return
	}

	response.OKRespWithData(c, map[string]interface{}{
		"video_list": videoList,
	})
}
