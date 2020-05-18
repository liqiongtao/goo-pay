package goo

import (
	"github.com/go-redis/redis"
	gooCache "googo.io/goo/cache"
)

func InitRedis(cf gooCache.Config) {
	gooCache.Init(cf)
}

func Redis() *redis.Client {
	return gooCache.Redis()
}
