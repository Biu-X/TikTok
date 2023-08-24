package cache

import (
	"bytes"
	"encoding/json"
	"github.com/Biu-X/TikTok/module/log"
	"github.com/gin-gonic/gin"
)

func PublishMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Logger.Infof("in publish")
		writer := responseWriter{
			c.Writer,
			bytes.NewBuffer([]byte{}),
		}
		c.Writer = writer
		var m map[string]interface{}
		err := json.Unmarshal(writer.b.Bytes(), &m)
		if err != nil {
			return
		}
		c.Next()
		log.Logger.Infof("map ------------> %v", m)
		log.Logger.Infof("response body: %v", writer.b.String())
		log.Logger.Infof("out publish")
	}
}
