package config

import (
	"os"

	"github.com/go-redis/redis/v8"
)

func ConnectRedis() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr: os.Getenv("REDIS_ADDRESS"),
	})
	return rdb
}