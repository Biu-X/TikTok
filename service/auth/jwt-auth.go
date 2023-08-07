package auth

import (
	"net/http"
	"time"

	"biu-x.org/TikTok/module/middleware/jwt"
	"github.com/gin-gonic/gin"
)

// 使用 cookie 持久化
func RequireAuth1(c *gin.Context) {
	// get the coolie off request
	tokenString, err := c.Cookie("Authorization")
	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
	}

	// decode tokenString to jwt.token(user secret)
	claims, err := jwt.ParseToken(tokenString)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"Messgae": err.Error(),
		})
		return
	}

	if float64(time.Now().Unix()) > float64(claims.ExpiresAt.Second()) {
		c.AbortWithStatus(http.StatusUnauthorized)
	}

	// 验证通过，将用户 ID 添加到 context 中
	c.Set("user_id", claims.ID)
	// 放行
	c.Next()
}

// 非持久化
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
		ctx.Set("user_id", claims.ID)
		// 验证通过，通行
		ctx.Next()
	}
}
