package gooWeixin

import (
	"fmt"
	"goo"
	gooLog "goo/log"
	gooUtils "goo/utils"
	"net/url"
	"strings"
	"time"
)

func JsApi(appid, secret, urlStr string) goo.Params {
	ticket, _ := GetCGITicket(appid, secret)

	ts := time.Now().Unix()
	nonceStr := gooUtils.NonceStr()

	urlStr, _ = url.QueryUnescape(urlStr)
	urlStr = strings.Split(urlStr, "#")[0]

	rawstr := fmt.Sprintf(jsapi_ticket_qs, ticket, nonceStr, ts, urlStr)
	rawstr = gooUtils.SHA1([]byte(rawstr))

	params := goo.Params{
		"debug":     false,
		"appId":     appid,
		"timestamp": ts,
		"nonceStr":  nonceStr,
		"signature": rawstr,
		"jsApiList": []string{"checkJsApi", "onMenuShareTimeline", "onMenuShareAppMessage", "chooseWXPay"},
	}

	gooLog.Debug("wx.jsapi:", params.Json())

	return params
}
