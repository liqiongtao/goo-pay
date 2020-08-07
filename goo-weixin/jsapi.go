package gooWeixin

import (
	"fmt"
	"googo.io/goo"
	gooLog "googo.io/goo/log"
	gooUtils "googo.io/goo/utils"
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
		"jsApiList": []string{"checkJsApi", "onMenuShareTimeline", "onMenuShareAppMessage", "chooseWXPay",
			"openLocation", "getLocation", "chooseImage", "previewImage", "uploadImage", "downloadImage",
			"startRecord", "stopRecord", "onVoiceRecordEnd", "playVoice", "pauseVoice", "stopVoice",
			"onVoicePlayEnd", "uploadVoice", "downloadVoice", "translateVoice", "getNetworkType", "scanQRCode",
			"addCard", "chooseCard", "openCard"},
	}

	gooLog.Debug(fmt.Sprintf("wx_jsapi=%s urlStr=%s", string(params.Json()), urlStr))

	return params
}
