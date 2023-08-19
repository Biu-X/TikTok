package v1

import (
	"net/http"

	"biu-x.org/TikTok/service/auth"
	comment_service "biu-x.org/TikTok/service/comment"
	favorite_service "biu-x.org/TikTok/service/favorite"
	message_service "biu-x.org/TikTok/service/message"
	publish_service "biu-x.org/TikTok/service/publish"
	relation_service "biu-x.org/TikTok/service/relation"
	user_service "biu-x.org/TikTok/service/user"
	"github.com/gin-gonic/gin"
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
			user.POST("register/", user_service.Signup)
			// 用户登录
			user.POST("login/", user_service.Login)
			// 用户信息
			user.Use(auth.RequireAuth())
			user.GET("", user_service.UserInfo)
		}

		publish := tiktok.Group("publish/")
		{
			publish.Use(auth.RequireAuth())
			// 投稿接口
			publish.POST("action/", publish_service.Action)
			// 发布列表
			publish.GET("list/", publish_service.List)
		}

		favorite := tiktok.Group("favorite/")
		{
			favorite.Use(auth.RequireAuth())
			// 赞操作
			favorite.POST("action/", favorite_service.Action)
			// 喜欢列表
			favorite.GET("list/", favorite_service.List)
		}

		comment := tiktok.Group("comment/")
		{
			comment.Use(auth.RequireAuth())
			// 评论操作
			comment.POST("action/", comment_service.Action)
			// 评论列表
			comment.GET("list/", comment_service.List)
		}

		relation := tiktok.Group("relation/")
		{
			relation.Use(auth.RequireAuth())
			// 关注操作
			relation.POST("action/", relation_service.Action)

			follow := relation.Group("follow/")
			{
				// 关注列表
				follow.GET("list/", relation_service.FollowList)
			}

			follower := relation.Group("follower/")
			{
				// 粉丝列表
				follower.GET("list/", relation_service.FollowerList)
			}

			friend := relation.Group("friend/")
			{
				// 好友列表
				friend.GET("list/", relation_service.FriendList)
			}
		}

		message := tiktok.Group("message/")
		{
			message.Use(auth.RequireAuth())
			// 发送消息
			message.POST("action/", message_service.Action)
			// 聊天记录
			message.GET("chat/", message_service.Chat)
		}
	}

	return r
}
