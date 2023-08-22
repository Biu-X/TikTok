package cache

import (
	"biu-x.org/TikTok/module/config"
	"biu-x.org/TikTok/module/db"
	"biu-x.org/TikTok/module/log"
	"biu-x.org/TikTok/module/oss"
	"testing"
)

func TestNewRedisClient(t *testing.T) {
	config.Init()
	log.Init()
	db.Init()
	oss.Init()

	clients := map[string]*Client{
		"IPLimit": &Client{},
		"Feed":    &Client{},
	}
	NewRedisClients(clients)

	for i, c := range clients {
		log.Logger.Infof("client %v: %v", i, c)
	}
	ipLimit := clients["IPLimit"]
	ipLimit.Set("name", "lala")
	name := ipLimit.Get("name")
	log.Logger.Infof("%v", name.Val())
	name = ipLimit.C.Get(ipLimit.ctx, "name")
	log.Logger.Infof("%v", name.Val())
}
