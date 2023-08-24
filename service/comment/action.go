package comment

import (
	"strconv"

	"github.com/Biu-X/TikTok/dal/model"
	"github.com/Biu-X/TikTok/dao"
	"github.com/Biu-X/TikTok/module/log"
	"github.com/Biu-X/TikTok/module/response"
	"github.com/Biu-X/TikTok/module/util"
	"github.com/gin-gonic/gin"
)

// Action /douyin/comment/action/ - 评论操作
func Action(c *gin.Context) {
	userID := util.GetUserIDFromGinContext(c)
	commentText := util.GetInsensitiveTextFromGinContext(c, "comment_text")
	videoIDStr := c.Query("video_id")
	actionTypeStr := c.Query("action_type")
	videoID, _ := strconv.ParseInt(videoIDStr, 10, 64)
	actionType, _ := strconv.Atoi(actionTypeStr) // 1-评论, 2-删除评论

	if actionType == 1 {
		comment := &model.Comment{
			UserID:  userID,
			VideoID: videoID,
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
