package gooWeixin

import (
	"encoding/json"
	"errors"
	"fmt"
	"goo"
	gooHttp "goo/http"
	gooLog "goo/log"
	"sync"
	"time"
)

type cgiAccessToken struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int64  `json:"expires_in"`
	ErrCode     int    `json:"errcode"`
	ErrMsg      string `json:"errmsg"`
}

var muGetCGIAccessToken sync.Mutex

func GetCGIAccessToken(appid, secret string) (string, error) {
	key := fmt.Sprintf(cgi_token_key, appid)

	ttl := int64(__cache.TTL(key).Val().Seconds())
	if ttl > 60 {
		return __cache.Get(key).Val(), nil
	}

	muGetCGIAccessToken.Lock()
	defer muGetCGIAccessToken.Unlock()

	rsp := &cgiAccessToken{}
	buf, _ := gooHttp.NewRequest().Get(fmt.Sprintf(cgi_token_url, appid, secret))

	if err := json.Unmarshal(buf, rsp); err != nil {
		gooLog.Error(err.Error())
		return "", err
	}

	if errCode := rsp.ErrCode; errCode != 0 {
		gooLog.Error(rsp.ErrMsg)
		return "", errors.New(rsp.ErrMsg)
	}
	if err := __cache.Set(key, rsp.AccessToken, time.Duration(rsp.ExpiresIn)*time.Second).Err(); err != nil {
		gooLog.Error(err.Error())
		return "", err
	}

	return rsp.AccessToken, nil
}

func AutoRefreshCGIAccessToken(appid, secret string) {
	t := time.NewTicker(300 * time.Second)

	goo.AsyncFunc(func() {
		for range t.C {
			accessToken, err := GetCGIAccessToken(appid, secret)

			if err != nil {
				gooLog.Error(
					fmt.Sprintf("appid=%s", appid),
					err.Error(),
				)
				continue
			}

			gooLog.Debug(
				fmt.Sprintf("appid=%s", appid),
				fmt.Sprintf("cgi-access-token=%s", accessToken),
			)
		}
	})
}
