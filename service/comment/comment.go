package comment

import (
	"net/http"
	"strconv"

	"biu-x.org/TikTok/dao"
	"biu-x.org/TikTok/model"
	"biu-x.org/TikTok/module/log"
	"github.com/gin-gonic/gin"
)

type Response struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"status_message"`
}

type CommentActionResponse struct {
	Response
	Comment *model.Comment
}
type CommentListResponse struct {
	Response
	CommentList []*model.Comment
}

// Action 评论操作 /douyin/comment/action/
func Action(c *gin.Context) {
	userIdStr := c.GetString("user_id")
	userId, _ := strconv.ParseInt(userIdStr, 10, 64)
	videoIdStr := c.Query("video_id")
	actionTypeStr := c.Query("action_type")
	videoId, _ := strconv.ParseInt(videoIdStr, 10, 64)
	actionType, _ := strconv.Atoi(actionTypeStr)

	if actionType == 1 {
		commentText := c.Query("comment_text")
		comment := &model.Comment{
			UserID:  int64(userId),
			VideoID: int64(videoId),
			Content: commentText,
		}
		err := dao.CreateComment(comment)
		if err != nil {
			log.Logger.Errorf("action: CreateComment failed, err: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{})
			return
		}
		c.JSON(http.StatusOK, CommentActionResponse{
			Response: Response{
				StatusCode: 0,
				Message:    "newcomment create success...",
			},
			Comment: comment,
		})
	} else {
		commentIdStr := c.Query("comment_id")
		commentId, _ := strconv.ParseInt(commentIdStr, 10, 64)
		err := dao.DeleteCommentByID(commentId)
		if err != nil {
			log.Logger.Errorf("action: DeleteCommentByID failed, err: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{})
			return
		}
		c.JSON(http.StatusOK, CommentActionResponse{
			Response: Response{
				StatusCode: 0,
				Message:    "comment delete success...",
			},
		})
	}
}

// List 评论操作 /douyin/comment/list/
func List(c *gin.Context) {
	videoIdStr := c.Query("video_id")
	videoId, _ := strconv.ParseInt(videoIdStr, 10, 64)
	commentList, err := dao.GetCommentByVideoID(videoId)
	if err != nil {
		log.Logger.Errorf("list: GetCommentByVideoID failed, err: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{})
		return
	}
	c.JSON(http.StatusOK, CommentListResponse{
		Response: Response{
			StatusCode: 0,
			Message:    "query success",
		},
		CommentList: commentList,
	})
}
