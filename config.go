package main

import (
	"github.com/redis/go-redis/v9"
	"os"
)

func GetRedisConfig() *redis.Options {
	host := os.Getenv("REDIS_HOST")
	port := os.Getenv("REDIS_PORT")
	return &redis.Options{
		Addr: host + ":" + port,
	}
}
