package auth

import (
	"net/http"
	"time"

	"biu-x.org/TikTok/module/log"
	"biu-x.org/TikTok/module/middleware/jwt"
	"biu-x.org/TikTok/module/response"
	"github.com/gin-gonic/gin"
)

// RequireAuth 鉴权中间件
// 如果用户携带的 token 验证通过，将 user_id 存入上下文中然后执行下一个 Handler
func RequireAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		url := c.Request.URL
		log.Logger.Infof("url: %v", url.Path)
		// 从输入的 url 中查询 token 值
		token := c.Query("token")
		if len(token) == 0 {
			// 从输入的表单中查询 token 值
			token = c.PostForm("token")
		}

		if len(token) == 0 {
			// 终止调用链，并不是返回
			c.AbortWithStatusJSON(http.StatusOK, response.AuthResponse{
				StatusCode:    -1,
				StatusMessage: "JSON WEB TOKEN IS NULL",
			})
			return
		}

		log.Logger.Info("token 读取成功")
		// auth = [[header][cliams][signature]]
		// 解析 token
		claims, err := jwt.ParseToken(token)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusOK, response.AuthResponse{
				StatusCode:    -1,
				StatusMessage: "ERR_INVALID_TOKEN",
			})
			return
		}
		// validate expire time
		if time.Now().Unix() > claims.ExpiresAt.Unix() {
			c.AbortWithStatusJSON(http.StatusOK, response.AuthResponse{
				StatusCode:    -1,
				StatusMessage: "TOKEN IS ALREADY EXPIRED, Please Log In Again",
			})
			return
		}

		userId := claims.ID
		c.Set("user_id", userId)
		c.Set("is_login", true)
		// 放行
		c.Next()
	}
}

// RequireAuthWithoutLogin 用户在未登录情况如果携带了 token
// 验证 token 有效性，如果有效，解析出用户 id 存入上下文，否则存入默认值 0
func RequireAuthWithoutLogin() gin.HandlerFunc {
	return func(c *gin.Context) {
		url := c.Request.URL
		log.Logger.Infof("url: %v", url.Path)
		token := c.Query("token")
		userId := "0"
		if len(token) != 0 {
			cliams, err := jwt.ParseToken(token)
			if err != nil {
				c.AbortWithStatusJSON(http.StatusOK, response.AuthResponse{
					StatusCode:    -1,
					StatusMessage: "ERR_INVALID_TOKEN",
				})
				return
			}

			userId = cliams.ID
			c.Set("user_id", userId)
			c.Set("is_login", true)
			c.Next()
		} else {
			c.AbortWithStatus(http.StatusOK)
		}
	}
}
