package jwt

import (
	"net/http"
	"time"

	"github.com/Biu-X/TikTok/module/log"
	"github.com/Biu-X/TikTok/module/response"
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
		claims, err := ParseToken(token)
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
		isLogin := false
		if len(token) != 0 {
			cliams, err := ParseToken(token)
			if err != nil {
				c.AbortWithStatusJSON(http.StatusOK, response.AuthResponse{
					StatusCode:    -1,
					StatusMessage: "ERR_INVALID_TOKEN",
				})
				return
			}
			log.Logger.Infof("token 获取成功并验证通过,无需登录")
			userId = cliams.ID
			isLogin = true
		} else {
			userId = "0"
		}
		c.Set("user_id", userId)
		c.Set("is_login", isLogin)
		log.Logger.Infof("user_id: %v, is_login: %v", userId, isLogin)
		// 放行
		c.Next()
	}
}

// RequireAuthCookie ，使用 cookie 持久化。 本实现暂时用不到，需要网页端时直接修改为这种持久化存储实现即可
func RequireAuthCookie(c *gin.Context) {
	// get the coolie off request
	tokenString, err := c.Cookie("Authorization")
	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
	}

	// decode tokenString to jwt.token(user secret)
	claims, err := ParseToken(tokenString)
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
