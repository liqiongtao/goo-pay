package gooCache

import (
	"github.com/go-redis/redis"
)

var (
	__cache *cache
)

type cache struct {
	Prefix string
	*redis.Client
}

func Init(cf Config) {
	__cache = &cache{
		Prefix: cf.Prefix,
		Client: NewRedis(cf),
	}
}

func Redis() *redis.Client {
	return __cache.Client
}
