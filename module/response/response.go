package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ErrReponse struct {
	StatusCode int    `json:"status_code"` // 错误码
	Message    string `json:"messgae"`     // 错误信息
}

func OKResp(c *gin.Context) {
	resp := map[string]interface{}{
		"status_code": OK,
		"status_msg":  Msg(OK),
	}
	c.JSON(http.StatusOK, resp)
}

func OKRespWithData(c *gin.Context, data map[string]interface{}) {
	resp := map[string]interface{}{
		"status_code": OK,
		"status_msg":  Msg(OK),
	}
	for key, value := range data {
		resp[key] = value
	}
	c.JSON(http.StatusOK, resp)
}

func ErrResp(c *gin.Context) {
	resp := map[string]interface{}{
		"status_code": Error,
		"status_msg":  Msg(Error),
	}
	c.JSON(http.StatusOK, resp)
}

func ErrRespWithMsg(c *gin.Context, msg string) {
	resp := map[string]interface{}{
		"status_code": Error,
		"status_msg":  msg,
	}
	c.JSON(http.StatusOK, resp)
}

type VideoResponse struct {
	VideoID       int64        `json:"id"`             // 视频唯一标识
	Author        UserResponse `json:"author"`         // 视频作者信息
	PlayURL       string       `json:"play_url"`       // 视频播放地址
	CoverURL      string       `json:"cover_url"`      // 视频封面地址
	FavoriteCount int64        `json:"favorite_count"` // 视频的点赞总数
	CommentCount  int64        `json:"comment_count"`  // 视频的评论总数
	IsFavorite    bool         `json:"is_favorite"`    // true-已点赞，false-未点赞
	Title         string       `json:"title"`          // 视频标题
}

type UserResponse struct {
	UserID         int64  `json:"id"`               // 用户ID
	Username       string `json:"name"`             // 用户名
	FollowCount    int64  `json:"follow_count"`     // 该用户关注了多少个其他用户
	FollowerCount  int64  `json:"follower_count"`   // 该用户粉丝总数
	IsFollow       bool   `json:"is_follow"`        // true: 已关注 false: 未关注
	Avatar         string `json:"avatar"`           // 头像
	BackGroudImage string `json:"background_image"` // 背景大图
	Signature      string `json:"signature"`        // 个人简介
	TotalFavorite  int64  `json:"total_favorite"`   // 该用户获赞总量
	WorkCount      int64  `json:"work_count"`       // 作品数量
	FavoriteCount  int64  `json:"favorite_count"`   // 喜欢的作品数量
}
