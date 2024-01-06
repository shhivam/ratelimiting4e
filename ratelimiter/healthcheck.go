package ratelimiter

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"net/http"
	"time"
)

func errorHandler(c *gin.Context, err error) {
	c.JSON(http.StatusTooManyRequests, gin.H{
		"message": "Rate Limit exceeded",
		"error":   err.Error(),
	})
	c.Abort()
}

func keyGenerator(c *gin.Context) string {
	route := c.FullPath()
	ipAddress := c.ClientIP()

	return fmt.Sprintf("%s:%s", ipAddress, route)
}

func GetHealthcheckRateLimiter(redisDB *redis.Client) RateLimiter {
	return RateLimiter{
		RedisDB:        redisDB,
		Duration:       1 * time.Minute,
		BucketCapacity: 10,
		ErrorHandler:   errorHandler,
		GetBucketName:  keyGenerator,
	}
}
