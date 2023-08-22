package cache

import (
	"biu-x.org/TikTok/module/config"
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"strconv"
	"time"
)

type RType uint8

const (
	Feed RType = iota
	Comment
	Favorite
	Message
	Publish
	Follow
	Follower
	Friend
	User
)

var Clients = map[RType]*Client{
	Feed:     &Client{},
	Comment:  &Client{},
	Favorite: &Client{},
	Message:  &Client{},
	Publish:  &Client{},
	Follow:   &Client{},
	Follower: &Client{},
	Friend:   &Client{},
	User:     &Client{},
}

type Client struct {
	C   *redis.Client //nolint:typecheck
	ctx context.Context
}

func Init() {
	NewRedisClients(Clients)
}

func NewRedisClients(clients map[RType]*Client) {
	poolSize, _ := strconv.ParseInt(config.GetString("redis.poolSize"), 10, 64)
	n := 0
	for k, _ := range clients {
		r := redis.NewClient(&redis.Options{ //nolint:typecheck
			Addr:     fmt.Sprintf("%v:%v", config.Get("redis.host"), config.Get("redis.port")),
			Password: fmt.Sprintf("%v", config.Get("redis.password")),
			DB:       n + 1,
			PoolSize: int(poolSize),
		})
		n++
		ctx := context.Background()
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

// 封装常用接口

// ClientGetName returns the name of the connection.
func (c Client) ClientGetName() *redis.StringCmd {
	return c.C.ClientGetName(c.ctx)
}

func (c Client) Echo(message interface{}) *redis.StringCmd {
	return c.Echo(message)
}

func (c Client) Ping() *redis.StatusCmd {
	return c.C.Ping(c.ctx)
}

func (c Client) Del(keys ...string) *redis.IntCmd {
	return c.C.Del(c.ctx, keys...)
}

func (c Client) Unlink(keys ...string) *redis.IntCmd {
	return c.C.Unlink(c.ctx, keys...)
}

func (c Client) Dump(key string) *redis.StringCmd {
	return c.Dump(key)
}

func (c Client) Exists(keys ...string) *redis.IntCmd {
	return c.C.Exists(c.ctx, keys...)
}

func (c Client) Expire(key string, expiration time.Duration) *redis.BoolCmd {
	return c.C.Expire(c.ctx, key, expiration)
}

func (c Client) ExpireNX(key string, expiration time.Duration) *redis.BoolCmd {
	return c.C.ExpireNX(c.ctx, key, expiration)
}

func (c Client) ExpireXX(key string, expiration time.Duration) *redis.BoolCmd {
	return c.C.ExpireXX(c.ctx, key, expiration)
}

func (c Client) ExpireGT(key string, expiration time.Duration) *redis.BoolCmd {
	return c.C.ExpireGT(c.ctx, key, expiration)
}

func (c Client) ExpireLT(key string, expiration time.Duration) *redis.BoolCmd {
	return c.C.ExpireLT(c.ctx, key, expiration)
}

func (c Client) ExpireAt(key string, tm time.Time) *redis.BoolCmd {
	return c.C.ExpireAt(c.ctx, key, tm)
}

func (c Client) ExpireTime(key string) *redis.DurationCmd {
	return c.C.ExpireTime(c.ctx, key)
}
