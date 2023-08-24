package v1

import (
	"github.com/Biu-X/TikTok/module/cache"
	middleware_cache "github.com/Biu-X/TikTok/module/middleware/cache"
	"github.com/Biu-X/TikTok/module/middleware/sensitiveguard"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/Biu-X/TikTok/module/middleware/jwt"
	"github.com/Biu-X/TikTok/module/middleware/logger"
	comment_service "github.com/Biu-X/TikTok/service/comment"
	favorite_service "github.com/Biu-X/TikTok/service/favorite"
	feed_service "github.com/Biu-X/TikTok/service/feed"
	message_service "github.com/Biu-X/TikTok/service/message"
	publish_service "github.com/Biu-X/TikTok/service/publish"
	relation_service "github.com/Biu-X/TikTok/service/relation"
	user_service "github.com/Biu-X/TikTok/service/user"
)

func NewAPI() *gin.Engine {
	r := gin.New()
	r.Use(logger.DefaultLogger(), gin.Recovery()) // 日志中间件
	r.Use(middleware_cache.NewRateLimiter(middleware_cache.Clients[cache.IPLimit], "general", 200, 60*time.Second))

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "Welcome to Tiktok!",
		})
	})

	tiktok := r.Group("/douyin/")
	{
		// 视频流接口
		tiktok.GET("feed/", jwt.RequireAuthWithoutLogin(), feed_service.List)

		user := tiktok.Group("user/")
		{
			// 用户注册
			user.POST("register/", user_service.Register)
			// 用户登录
			user.POST("login/", user_service.Login)
			// 用户信息
			user.Use(jwt.RequireAuth())
			user.GET("", user_service.Info)
		}

		publish := tiktok.Group("publish/")
		{
			publish.Use(jwt.RequireAuth())
			// 投稿接口
			publish.POST("action/", publish_service.Action)
			// 发布列表
			publish.GET("list/", publish_service.List)
		}

		favorite := tiktok.Group("favorite/")
		{
			favorite.Use(jwt.RequireAuth())
			// 赞操作
			favorite.POST("action/", favorite_service.Action)
			// 喜欢列表
			favorite.GET("list/", favorite_service.List)
		}

		comment := tiktok.Group("comment/")
		{
			// 评论列表
			comment.GET("list/", comment_service.List)
			comment.Use(jwt.RequireAuth())
			// 评论操作
			comment.POST("action/",
				sensitiveguard.SensitiveGuard("comment_text"),
				comment_service.Action)
		}

		relation := tiktok.Group("relation/")
		{
			relation.Use(jwt.RequireAuth())
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
			message.Use(jwt.RequireAuth())
			// 发送消息
			message.POST("action/",
				sensitiveguard.SensitiveGuard("content"),
				message_service.Action)
			// 聊天记录
			message.GET("chat/", message_service.Chat)
		}
	}

	return r
}
