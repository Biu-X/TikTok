package sensitiveguard

import (
	"github.com/Biu-X/TikTok/module/log"
	"github.com/Biu-X/TikTok/module/response"
	"github.com/Biu-X/TikTok/module/sensitive"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SensitiveGuard(queryKey string) gin.HandlerFunc {
	return func(c *gin.Context) {
		text := c.Query(queryKey)
		// 检查是否包含重点敏感词，爆了直接返回
		if !sensitive.ValidateBoss(text) {
			c.AbortWithStatusJSON(http.StatusOK, response.AuthResponse{
				StatusCode:    -1,
				StatusMessage: "Contain sensitive words",
			})
			log.Logger.Warn("Contain sensitive words")
			return
		}
		log.Logger.Info(text)
		// 和谐普通敏感词，替换成*
		t := sensitive.Replace(text)
		insensitiveTextKey := "insensitive_text_" + queryKey
		c.Set(insensitiveTextKey, t)
		// 放行
		c.Next()
	}
}
