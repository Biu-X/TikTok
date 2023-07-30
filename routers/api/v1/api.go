package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func NewAPI() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery())

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "Welcome to Tiktok!",
		})
	})

	tiktok := r.Group("/douyin/")
	{
		// 视频流接口
		tiktok.GET("feed/")

		user := tiktok.Group("user/")
		{
			// 用户注册
			user.POST("register/")
			// 用户登录
			user.POST("login/")
			// 用户信息
			user.GET("")
		}

		publish := tiktok.Group("publish/")
		{
			// 投稿接口
			publish.POST("action/")
			// 发布列表
			publish.GET("list/")
		}

		favorite := tiktok.Group("favorite/")
		{
			// 赞操作
			favorite.POST("action/")
			// 喜欢列表
			favorite.GET("list/")
		}

		comment := tiktok.Group("comment/")
		{
			// 评论操作
			comment.POST("action/")
			// 评论列表
			comment.GET("list/")
		}

		relation := tiktok.Group("relation/")
		{
			// 关注操作
			relation.POST("action/")

			follow := relation.Group("follow/")
			{
				// 关注列表
				follow.GET("list/")
			}

			follower := relation.Group("follower/")
			{
				// 粉丝列表
				follower.GET("list/")
			}

			friend := relation.Group("friend/")
			{
				// 好友列表
				friend.GET("list/")
			}
		}

		message := tiktok.Group("message/")
		{
			// 发送消息
			message.POST("action/")
			// 聊天记录
			message.GET("chat/")
		}
	}

	return r
}
