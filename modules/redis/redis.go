package redis

import (
	"biu-x.org/TikTok/modules/config"
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"strconv"
	"time"
)

type Client struct {
	*redis.Client //nolint:typecheck
	ctx           context.Context
}

func NewRedisClient() *Client {
	db, _ := strconv.Atoi(fmt.Sprintf("%v", config.Get("redis.db")))
	r := redis.NewClient(&redis.Options{ //nolint:typecheck
		Addr:     fmt.Sprintf("%v:%v", config.Get("redis.host"), config.Get("redis.port")),
		Password: fmt.Sprintf("%v", config.Get("redis.password")),
		DB:       db,
	})
	ctx := context.Background()
	return &Client{Client: r, ctx: ctx}
}

func (c Client) Set(key string, value interface{}) *redis.StatusCmd { //nolint:typecheck
	fmt.Printf("%#v\n", key)
	return c.Client.Set(c.ctx, key, value, 10*time.Second)
}

func (c Client) Get(key interface{}) *redis.StringCmd { //nolint:typecheck
	return c.Client.Get(c.ctx, key.(string))
}
