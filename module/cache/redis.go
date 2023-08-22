package cache

import (
	"context"
	"fmt"

	"biu-x.org/TikTok/module/config"
	"github.com/redis/go-redis/v9"
)

var (
	Ctx = context.Background()
	// 关注
	RedisFollowers *redis.Client
	RedisFollowing *redis.Client
	RedisFriends   *redis.Client
	// 赞
	RedisFavoriteByUserId  *redis.Client
	RedisFavoriteByVideoId *redis.Client
	// 评论和视频
	RedisRecordByVideoAndCommentId *redis.Client
	RedisRecordByCommentAndVideoId *redis.Client
)

func NewRedisClients() {
	// 粉丝信息存入 DB0
	RedisFollowers = NewRedisClient(0)
	// 关注信息存入 DB1
	RedisFollowing = NewRedisClient(1)
	// 相互关注信息存入 DB2
	RedisFriends = NewRedisClient(2)
	// 将某个用户所有点赞的视频 id 存入 DB3
	RedisFavoriteByUserId = NewRedisClient(3)
	// 将某个视频所有点赞的用户 id 存入 DB4
	RedisFavoriteByVideoId = NewRedisClient(4)
	// 将某个视频的所有评论 id 存入 DB5
	RedisRecordByVideoAndCommentId = NewRedisClient(5)
	// 将某个评论对应的视频 id 存入 DB6
	RedisRecordByCommentAndVideoId = NewRedisClient(6)
}

func NewRedisClient(DbNumber int) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%v:%v", config.Get("redis.host"), config.Get("redis.port")),
		Password: fmt.Sprintf("%v", config.Get("redis.password")),
		DB:       DbNumber,
	})
}
