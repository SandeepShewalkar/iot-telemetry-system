package storage

import (
	"os"

	"github.com/go-redis/redis/v8"
)

var rdb *redis.Client

func NewRedisClient(addr string) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr: addr,
	})
}

func GetRedisClient() *redis.Client {
	redisAddr := os.Getenv("REDIS_ADDR")
	if rdb == nil {
		return NewRedisClient(redisAddr)
	}
	return rdb
}
