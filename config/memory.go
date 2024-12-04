package config

import "github.com/go-redis/redis/v8"

func ConnectRedis() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr: "10.8.141.51:6379",
	})
	return rdb
}