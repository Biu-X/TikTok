package sensitiveguard

import (
	"biu-x.org/TikTok/module/log"
	"biu-x.org/TikTok/module/response"
	"biu-x.org/TikTok/module/sensitive"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SensitiveGuard(queryKey string) gin.HandlerFunc {
	return func(c *gin.Context) {
		Text := c.Query(queryKey)
		if !sensitive.Validate(Text) {
			c.AbortWithStatusJSON(http.StatusOK, response.AuthResponse{
				StatusCode:    -1,
				StatusMessage: "Contain sensitive words",
			})
			log.Logger.Warn("Contain sensitive words")
			return
		}
		c.Next()
	}
}
