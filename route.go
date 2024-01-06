package main

import (
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"shivamsinghal.me/ratelimiting4e/ratelimiter"
)

func NewRouter() *gin.Engine {
	r := gin.Default()

	redisDB := redis.NewClient(GetRedisConfig())

	healthcheckRateLimiter := ratelimiter.GetHealthcheckRateLimiter(redisDB)

	r.GET("/healthcheck", healthcheckRateLimiter.Limit, Healthcheck)

	return r
}
