package goo

import (
	"github.com/go-redis/redis"
	"goo/cache"
)

func InitRedis(cf gooCache.Config) {
	gooCache.Init(cf)
}

func Redis() *redis.Client {
	return gooCache.Redis()
}
