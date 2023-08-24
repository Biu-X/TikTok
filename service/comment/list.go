package comment

import (
	"strconv"

	"github.com/Biu-X/TikTok/dao"
	"github.com/Biu-X/TikTok/module/log"
	"github.com/Biu-X/TikTok/module/response"
	"github.com/Biu-X/TikTok/module/util"
	"github.com/gin-gonic/gin"
)

// List /douyin/comment/list/ - 视频评论列表
func List(c *gin.Context) {
	userID := util.GetUserIDFromGinContext(c)
	videoIDStr := c.Query("video_id")
	videoID, _ := strconv.ParseInt(videoIDStr, 10, 64)
	commentList, err := dao.GetCommentByVideoID(videoID)
	if err != nil {
		log.Logger.Errorf("list: GetCommentByVideoID failed, err: %v", err)
		response.ErrRespWithMsg(c, err.Error())
		return
	}

	var commentResponseList []response.CommentResponse
	for _, comment := range commentList {
		user, err := response.GetUserResponseByID(comment.UserID, userID)
		if err != nil {
			log.Logger.Errorf("list: GetUserResponseByID failed, err: %v", err)
			response.ErrRespWithMsg(c, err.Error())
			return
		}
		commentResponseList = append(commentResponseList, response.CommentResponse{
			CommentID:  comment.ID,
			User:       *user,
			Content:    comment.Content,
			CreateDate: comment.CreatedAt.Format("01-02"),
		})
	}
	response.OKRespWithData(c, map[string]interface{}{
		"comment_list": commentResponseList,
	})
}
