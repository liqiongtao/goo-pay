package gooWXPay

import (
	"encoding/xml"
	"errors"
	"fmt"
	"strings"
	"time"
)

type UnifiedOrderRequest struct {
	Appid          string    `xml:"appid"`
	MchId          string    `xml:"mch_id"`
	NonceStr       string    `xml:"nonce_str"`
	Sign           string    `xml:"sign"`
	SignType       SignType  `xml:"sign_type"`
	Body           string    `xml:"body"`
	Detail         string    `xml:"detail"`
	Attach         string    `xml:"attach"`
	OutTradeNo     string    `xml:"out_trade_no"`
	TotalFee       int       `xml:"total_fee"`
	SpbillCreateIp string    `xml:"spbill_create_ip"`
	GoodsTag       string    `xml:"goods_tag"`
	NotifyUrl      string    `xml:"notify_url"`
	TradeType      TradeType `xml:"trade_type"`
	ProductId      string    `xml:"product_id"`
	LimitPay       string    `xml:"limit_pay"`
	Openid         string    `xml:"openid"`
	Receipt        string    `xml:"receipt"`
	SceneInfo      string    `xml:"scene_info"`
}

func (uo *UnifiedOrderRequest) toXml(apiKey string) []byte {
	if uo.NonceStr == "" {
		uo.NonceStr = gooUtils.NonceStr()
	}
	if uo.SignType == "" {
		uo.SignType = SIGN_TYPE_HMAC_SHA256
	}

	str := obj2querystring(uo) + fmt.Sprintf("&key=%s", apiKey)
	gooLog.Debug("[UnifiedOrderRequest.querystring]", str)

	if uo.SignType == SIGN_TYPE_HMAC_SHA256 {
		uo.Sign = strings.ToUpper(gooUtils.HMacSha256([]byte(str), []byte(apiKey)))
	} else if uo.SignType == SIGN_TYPE_MD5 {
		uo.Sign = strings.ToUpper(gooUtils.MD5([]byte(str)))
	}

	return obj2xml(uo)
}

type UnifiedOrderResponse struct {
	ReturnCode string `xml:"return_code"`
	ReturnMsg  string `xml:"return_msg"`

	ResultCode string `xml:"result_code"`
	ErrCode    string `xml:"err_code"`
	ErrCodeDes string `xml:"err_code_des"`

	Appid    string `xml:"appid"`
	MchId    string `xml:"mch_id"`
	NonceStr string `xml:"nonce_str"`
	Sign     string `xml:"sign"`

	TradeType TradeType `xml:"trade_type"`
	PrepayId  string    `xml:"prepay_id"`
	CodeUrl   string    `xml:"code_url"`
}

func (uo *UnifiedOrderResponse) toJsApi(apiKey string, signType SignType) map[string]interface{} {
	data := map[string]interface{}{
		"appId":     uo.Appid,
		"timeStamp": fmt.Sprintf("%d", time.Now().Unix()),
		"nonceStr":  uo.NonceStr,
		"package":   fmt.Sprintf("prepay_id=%s", uo.PrepayId),
		"signType":  signType,
		"paySign":   "",
	}

	str := map2querystring(data) + fmt.Sprintf("&key=%s", apiKey)
	gooLog.Debug("[UnifiedOrder.JSAPI.querystring]", str)

	if signType == SIGN_TYPE_HMAC_SHA256 {
		data["paySign"] = strings.ToUpper(gooUtils.HMacSha256([]byte(str), []byte(apiKey)))
	} else if signType == SIGN_TYPE_MD5 {
		data["paySign"] = strings.ToUpper(gooUtils.MD5([]byte(str)))
	}

	return data
}

func (uo *UnifiedOrderResponse) toApp(apiKey string, signType SignType) map[string]interface{} {
	data := map[string]interface{}{
		"appId":        uo.Appid,
		"partnerId":    uo.MchId,
		"prepayId":     uo.PrepayId,
		"packageValue": "Sign=WXPay",
		"nonceStr":     uo.NonceStr,
		"timeStamp":    fmt.Sprintf("%d", time.Now().Unix()),
		"sign":         "",
	}

	str := map2querystring(data) + fmt.Sprintf("&key=%s", apiKey)
	gooLog.Debug("[UnifiedOrder.APP.querystring]", str)

	if signType == SIGN_TYPE_HMAC_SHA256 {
		data["sign"] = strings.ToUpper(gooUtils.HMacSha256([]byte(str), []byte(apiKey)))
	} else if signType == SIGN_TYPE_MD5 {
		data["sign"] = strings.ToUpper(gooUtils.MD5([]byte(str)))
	}

	return data
}

func UnifiedOrder(req *UnifiedOrderRequest, apiKey string) (map[string]interface{}, error) {
	buf := req.toXml(apiKey)
	gooLog.Debug("[UnifiedOrderRequest.xml]", string(buf))

	rstBuf, err := gooHttp.NewRequest().Post(URL_UNIFIED_ORDER, buf)
	if err != nil {
		return nil, err
	}

	gooLog.Debug("[UnifiedOrderResponse.xml]", string(rstBuf))

	rsp := UnifiedOrderResponse{}
	if err := xml.Unmarshal(rstBuf, &rsp); err != nil {
		return nil, err
	}
	if rsp.ReturnCode == FAIL {
		return nil, errors.New(rsp.ReturnMsg)
	}
	if rsp.ResultCode == FAIL {
		return nil, errors.New(rsp.ErrCodeDes)
	}

	if rsp.TradeType == TRADE_TYPE_JSAPI {
		return rsp.toJsApi(apiKey, req.SignType), nil
	}

	if rsp.TradeType == TRADE_TYPE_APP {
		return rsp.toApp(apiKey, req.SignType), nil
	}

	return nil, nil
}
