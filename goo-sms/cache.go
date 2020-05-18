package gooSms

import (
	"fmt"
	"googo.io/goo"
	"time"
)

var __cache = cache{}

type cache struct{}

func (this cache) setCode(appid, mobile, action, code string, expireIn time.Duration) error {
	key := fmt.Sprintf(codeKey, appid, mobile, action)
	return goo.Redis().Set(key, code, expireIn).Err()
}

func (this cache) getCode(appid, mobile, action string) string {
	key := fmt.Sprintf(codeKey, appid, mobile, action)
	return goo.Redis().Get(key).Val()
}
