package cache

import (
	"bytes"
	"github.com/Biu-X/TikTok/module/cache"
	"github.com/gin-gonic/gin"
)

var Clients = map[cache.RDB]*cache.Client{
	cache.Feed:     &cache.Client{},
	cache.Comment:  &cache.Client{},
	cache.Favorite: &cache.Client{},
	cache.Message:  &cache.Client{},
	cache.Publish:  &cache.Client{},
	cache.Follow:   &cache.Client{},
	cache.Follower: &cache.Client{},
	cache.Friend:   &cache.Client{},
	cache.User:     &cache.Client{},
	cache.IPLimit:  &cache.Client{},
}

type responseWriter struct {
	gin.ResponseWriter
	b *bytes.Buffer
}

func (w responseWriter) Write(b []byte) (int, error) {
	//向一个bytes.buffer中写一份数据来为获取body使用
	w.b.Write(b)
	//完成gin.Context.Writer.Write()原有功能
	return w.ResponseWriter.Write(b)
}

func Init() {
	cache.NewRedisClients(Clients)
}
