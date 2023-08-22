package cache

import (
	"biu-x.org/TikTok/module/config"
	"biu-x.org/TikTok/module/db"
	"biu-x.org/TikTok/module/log"
	"biu-x.org/TikTok/module/oss"
	"testing"
)

func TestNewRedisClient(t *testing.T) {
	defer func(c map[RDB]*Client) {
		for k, _ := range c {
			err := c[k].C.Close()
			if err != nil {
				return
			}
		}
	}(Clients)
	config.Init()
	log.Init()
	db.Init()
	oss.Init()
	Init()
	feed := Clients[Feed]
	result, err := feed.Ping().Result()
	if err != nil {
		return
	}
	log.Logger.Infof("Ping -> %v", result)
}
