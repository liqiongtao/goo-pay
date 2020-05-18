package gooWeixin

import (
	"encoding/json"
	"errors"
	"fmt"
	gooHttp "goo/http"
	gooLog "goo/log"
	"sync"
	"time"
)

type cgiTicket struct {
	Ticket    string `json:"ticket"`
	ExpiresIn int64  `json:"expires_in"`
	ErrCode   int    `json:"errcode"`
	ErrMsg    string `json:"errmsg"`
}

var muGetCGITicket sync.Mutex

func GetCGITicket(appid, secret string) (string, error) {
	key := fmt.Sprintf(cgi_ticket_key, appid)

	ttl := int64(__cache.TTL(key).Val().Seconds())
	if ttl > 30 {
		return __cache.Get(key).Val(), nil
	}

	muGetCGITicket.Lock()
	defer muGetCGITicket.Unlock()

	accessToken, err := GetCGIAccessToken(appid, secret)
	if err != nil {
		return "", err
	}

	rsp := &cgiTicket{}
	buf, _ := gooHttp.NewRequest().Get(fmt.Sprintf(cgi_ticket_url, accessToken))

	if err := json.Unmarshal(buf, rsp); err != nil {
		gooLog.Error(err.Error())
		return "", err
	}
	if errCode := rsp.ErrCode; errCode != 0 {
		gooLog.Error(rsp.ErrMsg)
		return "", errors.New(rsp.ErrMsg)
	}
	if err := __cache.Set(key, rsp.Ticket, time.Duration(rsp.ExpiresIn)*time.Second).Err(); err != nil {
		gooLog.Error(err.Error())
		return "", err
	}

	return rsp.Ticket, nil
}
