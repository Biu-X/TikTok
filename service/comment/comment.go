package comment

import (
	"strconv"

	"biu-x.org/TikTok/dao"
	"biu-x.org/TikTok/model"
	"biu-x.org/TikTok/module/log"
	"biu-x.org/TikTok/module/response"
	"github.com/gin-gonic/gin"
)

// Action 评论操作 /douyin/comment/action/
func Action(c *gin.Context) {
	userIDStr := c.GetString("user_id")
	userID, _ := strconv.ParseInt(userIDStr, 10, 64)
	videoIDStr := c.Query("video_id")
	actionTypeStr := c.Query("action_type")
	videoID, _ := strconv.ParseInt(videoIDStr, 10, 64)
	actionType, _ := strconv.Atoi(actionTypeStr)

	if actionType == 1 {
		// create comment
		commentText := c.Query("comment_text")
		comment := &model.Comment{
			UserID:  int64(userID),
			VideoID: int64(videoID),
			Content: commentText,
		}
		err := dao.CreateComment(comment)
		if err != nil {
			log.Logger.Errorf("action: CreateComment failed, err: %v", err)
			response.ErrRespWithMsg(c, err.Error())
			return
		}
		user, err := response.GetUserResponseByID(comment.UserID, userID)
		response.OKRespWithData(c, map[string]interface{}{
			"comment": response.CommentResponse{
				CommentID:  comment.ID,
				User:       *user,
				Content:    comment.Content,
				CreateDate: comment.CreatedAt.Format("01-02"),
			},
		})
	} else {
		// delete comment
		commentIDStr := c.Query("comment_id")
		commentID, _ := strconv.ParseInt(commentIDStr, 10, 64)
		err := dao.DeleteCommentByID(commentID)
		if err != nil {
			log.Logger.Errorf("action: DeleteCommentByID failed, err: %v", err)
			response.ErrRespWithMsg(c, err.Error())
			return
		}
		response.OKResp(c)
	}
}

// List 评论操作 /douyin/comment/list/
func List(c *gin.Context) {
	videoIDStr := c.Query("video_id")
	videoID, _ := strconv.ParseInt(videoIDStr, 10, 64)
	userIDStr := c.GetString("user_id")
	userID, _ := strconv.ParseInt(userIDStr, 10, 64)
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
