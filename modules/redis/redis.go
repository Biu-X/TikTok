package redis

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"time"
)

type Client struct {
	*redis.Client
	ctx context.Context
}

func NewRedisClient() *Client {
	r := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "123456",
		DB:       0,
	})
	ctx := context.Background()
	return &Client{Client: r, ctx: ctx}
}

func (c Client) Set(key string, value interface{}) *redis.StatusCmd {
	fmt.Printf("%#v\n", key)
	return c.Client.Set(c.ctx, key, value, 10*time.Second)
}

func (c Client) Get(key interface{}) *redis.StringCmd {
	return c.Client.Get(c.ctx, key.(string))
}
