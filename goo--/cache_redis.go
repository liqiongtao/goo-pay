package goo__

import (
	"github.com/go-redis/redis"
)

type Redis struct {
	redis.Client
}

func (r *Redis) init() {

}

func (r *Redis) New() {

}

func NewRedis(cf CacheConfig) *redis.Client {
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
