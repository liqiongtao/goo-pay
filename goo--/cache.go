package goo__

type icache interface {
	init()
}

func NewCache(ca icache) icache {
	ca.init()
	return ca
}

var __cache icache

func InitCache(cache icache) {
	__cache = NewCache(cache)
}

func Cache() icache {
	return __cache
}

// var (
// 	__cache *cache
// )
//
// type cache struct {
// 	Prefix string
// 	*redis.Client
// }
//
// func Init(cf Config) {
// 	__cache = &cache{
// 		Prefix: cf.Prefix,
// 		Client: NewRedis(cf),
// 	}
// }
//
// func Redis() *redis.Client {
// 	return __cache.Client
// }
