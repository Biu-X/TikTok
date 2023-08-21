package comment

import (
	"strconv"

	"biu-x.org/TikTok/dao"
	"biu-x.org/TikTok/dal/model"
	"biu-x.org/TikTok/module/log"
	"biu-x.org/TikTok/module/response"
	"github.com/gin-gonic/gin"
)

// Action 评论操作 /douyin/comment/action/
func Action(c *gin.Context) {
	// 从 RequireAuth 处读取 user_id
	userID, _ := strconv.ParseInt(c.GetString("user_id"), 10, 64)
	videoIDStr := c.Query("video_id")
	actionTypeStr := c.Query("action_type")
	videoID, _ := strconv.ParseInt(videoIDStr, 10, 64)
	actionType, _ := strconv.Atoi(actionTypeStr) // 1-评论, 2-删除评论
	
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
		if err != nil {
			log.Logger.Errorf("action: GetUserResponseByID failed, err: %v", err)
			response.ErrRespWithMsg(c, err.Error())
			return
		}
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
