package storage

import (
	"context"
	"github.com/go-redis/redis/v8"
)

type RedisDriver interface {
	Get(context context.Context, key string) *redis.StringCmd
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
