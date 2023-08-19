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
	userIdStr := c.GetString("user_id")
	userId, _ := strconv.ParseInt(userIdStr, 10, 64)
	videoIdStr := c.Query("video_id")
	actionTypeStr := c.Query("action_type")
	videoId, _ := strconv.ParseInt(videoIdStr, 10, 64)
	actionType, _ := strconv.Atoi(actionTypeStr)

	if actionType == 1 {
		// create comment
		commentText := c.Query("comment_text")
		comment := &model.Comment{
			UserID:  int64(userId),
			VideoID: int64(videoId),
			Content: commentText,
		}
		err := dao.CreateComment(comment)
		if err != nil {
			log.Logger.Errorf("action: CreateComment failed, err: %v", err)
			response.ErrRespWithMsg(c, err.Error())
			return
		}
		response.OKRespWithData(c, map[string]interface{}{
			"comment": response.CommentResponse{
				CommentID: comment.ID,
				//User: ,
				Content:    comment.Content,
				CreateDate: comment.CreatedAt.Format("mm-dd"),
			},
		})
	} else {
		// delete comment
		commentIdStr := c.Query("comment_id")
		commentId, _ := strconv.ParseInt(commentIdStr, 10, 64)
		err := dao.DeleteCommentByID(commentId)
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
	videoIdStr := c.Query("video_id")
	videoId, _ := strconv.ParseInt(videoIdStr, 10, 64)
	commentList, err := dao.GetCommentByVideoID(videoId)
	if err != nil {
		log.Logger.Errorf("list: GetCommentByVideoID failed, err: %v", err)
		response.ErrRespWithMsg(c, err.Error())
		return
	}

	commentResponseList := []response.CommentResponse{}
	for _, comment := range commentList {
		commentResponseList = append(commentResponseList, response.CommentResponse{
			CommentID: comment.ID,
			//User: ,
			Content:    comment.Content,
			CreateDate: comment.CreatedAt.Format("mm-dd"),
		})
	}
	response.OKRespWithData(c, map[string]interface{}{
		"comment_list": commentResponseList,
	})
}
