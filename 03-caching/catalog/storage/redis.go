package storage

import (
	"context"
	"github.com/go-redis/redis/v8"
	"time"
)

type RedisDriver interface {
	Get(context context.Context, key string) *redis.StringCmd
	Set(context context.Context, key string, value interface{}, expiration time.Duration) *redis.StatusCmd
	Del(context context.Context, keys ...string) *redis.IntCmd
}

type RedisConfig struct {
	Addr string
}

func NewRedisClient(config RedisConfig) RedisDriver {
	rdb := redis.NewClient(&redis.Options{
		Addr: config.Addr,
	})

	return rdb
}
