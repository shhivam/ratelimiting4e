package ratelimiter

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"strconv"
	"time"
)

type RateLimiter struct {
	RedisDB        *redis.Client
	Duration       time.Duration
	BucketCapacity int64
	ErrorHandler   func(c *gin.Context, err error)
	GetBucketName  func(c *gin.Context) string
}

func (r RateLimiter) handleIncrement(c *gin.Context, cacheKey string) error {
	value, err := r.RedisDB.Incr(c, cacheKey).Result()
	if value == 1 {
		// Key didn't exist before, so
		// we just created it, and we need to set an expiry
		r.RedisDB.Expire(c, cacheKey, r.Duration)
	}

	return err
}

func (r RateLimiter) Limit(c *gin.Context) {
	bucketName := r.GetBucketName(c)
	currentBucketSize, err := r.RedisDB.Get(c, bucketName).Result()

	if err != nil && errors.Is(err, redis.Nil) {
		err := r.handleIncrement(c, bucketName)
		if err != nil {
			r.ErrorHandler(c, err)
		}
		c.Next()
		return
	} else if err != nil {
		r.ErrorHandler(c, err)
		return
	}

	intValue, err := strconv.ParseInt(currentBucketSize, 10, 64)
	if intValue < r.BucketCapacity {
		r.handleIncrement(c, bucketName)
		c.Next()
	} else {
		fmt.Println("got too many requests")
		r.ErrorHandler(c, errors.New("too many requests"))
		return
	}
}
