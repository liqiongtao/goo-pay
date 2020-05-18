package gooCache

import (
	"github.com/go-redis/redis"
	"goo/log"
)

func NewRedis(cf Config) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     cf.Addr,
		Password: cf.Password,
		DB:       cf.DB,
	})

	if err := client.Ping().Err(); err != nil {
		gooLog.Error(err.Error())
		return nil
	}

	return client
}
