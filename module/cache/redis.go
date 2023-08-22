package cache

import (
	"biu-x.org/TikTok/module/config"
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"time"
)

type Client struct {
	C   *redis.Client //nolint:typecheck
	ctx context.Context
}

func NewRedisClients(clients map[string]*Client) {
	ctx := context.Background()
	n := 0
	for k, _ := range clients {
		r := redis.NewClient(&redis.Options{ //nolint:typecheck
			Addr:     fmt.Sprintf("%v:%v", config.Get("redis.host"), config.Get("redis.port")),
			Password: fmt.Sprintf("%v", config.Get("redis.password")),
			DB:       n + 1,
		})
		n++
		clients[k] = &Client{C: r, ctx: ctx}
	}
}

func NewRedisClient(n int) *Client {
	ctx := context.Background()
	r := redis.NewClient(&redis.Options{ //nolint:typecheck
		Addr:     fmt.Sprintf("%v:%v", config.Get("redis.host"), config.Get("redis.port")),
		Password: fmt.Sprintf("%v", config.Get("redis.password")),
		DB:       n,
	})
	return &Client{C: r, ctx: ctx}
}

func (c Client) Set(key string, value interface{}) *redis.StatusCmd { //nolint:typecheck
	return c.C.Set(c.ctx, key, value, 10*time.Hour)
}

func (c Client) Get(key interface{}) *redis.StringCmd { //nolint:typecheck
	return c.C.Get(c.ctx, key.(string))
}
