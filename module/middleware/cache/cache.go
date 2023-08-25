package cache

import (
	"bytes"

	"github.com/Biu-X/TikTok/module/cache"
	"github.com/gin-gonic/gin"
)

var Clients = map[cache.RDB]*cache.Client{
	cache.Feed:     {},
	cache.Comment:  {},
	cache.Favorite: {},
	cache.Message:  {},
	cache.Publish:  {},
	cache.Follow:   {},
	cache.Follower: {},
	cache.Friend:   {},
	cache.User:     {},
	cache.IPLimit:  {},
}

type responseWriter struct {
	gin.ResponseWriter
	b *bytes.Buffer
}

func (w responseWriter) Write(b []byte) (int, error) {
	// 向一个bytes.buffer中写一份数据来为获取body使用
	w.b.Write(b)
	// 完成gin.Context.Writer.Write()原有功能
	return w.ResponseWriter.Write(b)
}

func Init() {
	cache.NewRedisClients(Clients)
}
