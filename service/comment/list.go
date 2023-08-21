package comment

import (
	"strconv"

	"biu-x.org/TikTok/dao"
	"biu-x.org/TikTok/module/log"
	"biu-x.org/TikTok/module/response"
	"github.com/gin-gonic/gin"
)

// List 评论列表 /douyin/comment/list/
func List(c *gin.Context) {
	// 从 RequireAuth 处读取 user_id
	userID, _ := strconv.ParseInt(c.GetString("user_id"), 10, 64)
	videoIDStr := c.Query("video_id")
	videoID, _ := strconv.ParseInt(videoIDStr, 10, 64)
	commentList, err := dao.GetCommentByVideoID(videoID)
	if err != nil {
		log.Logger.Errorf("list: GetCommentByVideoID failed, err: %v", err)
		response.ErrRespWithMsg(c, err.Error())
		return
	}

	commentResponseList := []response.CommentResponse{}
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
