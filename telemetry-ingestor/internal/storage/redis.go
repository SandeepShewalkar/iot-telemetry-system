package storage

import (
	"os"
	"sync"

	"github.com/go-redis/redis/v8"
)

var (
	rdb  *redis.Client
	once sync.Once
)

func NewRedisClient(addr string) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr: addr,
	})
}

func GetRedisClient() *redis.Client {
	once.Do(func() {
		redisAddr := os.Getenv("REDIS_ADDR")
		rdb = NewRedisClient(redisAddr)
	})
	return rdb
}
