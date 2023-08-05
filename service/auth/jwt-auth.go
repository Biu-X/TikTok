package auth

import (
	"net/http"
	"time"

	"biu-x.org/TikTok/module/middleware/jwt"
	"github.com/gin-gonic/gin"
)

// JWT Auth
func JWTAuthMiddleWare() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		// 获取 jwt
		tokenStr := ctx.Query("token")
		// 如果从上下文中没有查询到，尝试解析 url
		if len(tokenStr) == 0 {
			tokenStr = ctx.PostForm("token")
		}
		if len(tokenStr) == 0 {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"code": -1,
				"msg":  "get jwt token failed",
			})
			return
		}
		// get user info from claims
		claims, err := jwt.ParseToken(tokenStr)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"code": -1,
				"msg":  "ERR_INVALID_TOKEN",
			})
			return
		}
		// Is jwt out of date?
		if time.Now().Unix() > claims.ExpiresAt.Unix() {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"code": -1,
				"msg":  "ERR_TOKEN_EXPIRED",
			})
			return
		}

		// auth passed, Put valid information about the user into context
		ctx.Set("username", claims.UserName)
		ctx.Set("user_id", claims.ID)
		// 验证通过，通行
		ctx.Next()
	}
}
