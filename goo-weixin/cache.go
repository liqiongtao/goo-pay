package gooWeixin

import "github.com/go-redis/redis"

var (
	__cache *redis.Client
)

func InitCache(redisClient *redis.Client) {
	__cache = redisClient
}
