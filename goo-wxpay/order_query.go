package gooWXPay

import (
	"encoding/xml"
	"errors"
	"fmt"
	gooHttp "goo/http"
	gooLog "goo/log"
	gooUtils "goo/utils"
	"strings"
)

type QueryOrderRequest struct {
	Appid         string   `xml:"appid"`
	MchId         string   `xml:"mch_id"`
	TransactionId string   `xml:"transaction_id"`
	OutTradeNo    string   `xml:"out_trade_no"`
	NonceStr      string   `xml:"nonce_str"`
	Sign          string   `xml:"sign"`
	SignType      SignType `xml:"sign_type"`
}

func (qo *QueryOrderRequest) toXml(apiKey string) []byte {
	if qo.NonceStr == "" {
		qo.NonceStr = gooUtils.NonceStr()
	}
	if qo.SignType == "" {
		qo.SignType = SIGN_TYPE_HMAC_SHA256
	}

	str := obj2querystring(qo) + fmt.Sprintf("&key=%s", apiKey)
	gooLog.Debug("[UnifiedOrderRequest.querystring]", str)

	if qo.SignType == SIGN_TYPE_HMAC_SHA256 {
		qo.Sign = strings.ToUpper(gooUtils.HMacSha256([]byte(str), []byte(apiKey)))
	} else if qo.SignType == SIGN_TYPE_MD5 {
		qo.Sign = strings.ToUpper(gooUtils.MD5([]byte(str)))
	}

	return obj2xml(qo)
}

type QueryOrderResponse struct {
	ReturnCode string `xml:"return_code"`
	ReturnMsg  string `xml:"return_msg"`

	ResultCode string `xml:"result_code"`
	ErrCode    string `xml:"err_code"`
	ErrCodeDes string `xml:"err_code_des"`

	Appid    string `xml:"appid"`
	MchId    string `xml:"mch_id"`
	NonceStr string `xml:"nonce_str"`
	Sign     string `xml:"sign"`

	DeviceInfo         string    `xml:"device_info"`
	Openid             string    `xml:"openid"`
	IsSubscribe        string    `xml:"is_subscribe"`
	TradeType          TradeType `xml:"trade_type"`
	TradeState         string    `xml:"trade_state"`
	BankType           string    `xml:"bank_type"`
	TotalFee           int64     `xml:"total_fee"`
	SettlementTotalFee int64     `xml:"settlement_total_fee"`
	CashFee            int64     `xml:"cash_fee"`
	CouponFee          int64     `xml:"coupon_fee"`
	CouponCount        int64     `xml:"coupon_count"`
	OutTradeNo         string    `xml:"out_trade_no"`
	Attach             string    `xml:"attach"`
	TimeEnd            string    `xml:"time_end"`
	TradeStateDesc     string    `xml:"trade_state_desc"`
}

func QueryOrder(req *QueryOrderRequest, apiKey string) (*QueryOrderResponse, error) {
	buf := req.toXml(apiKey)
	gooLog.Debug("[QueryOrderRequest.xml]", string(buf))

	rstBuf, err := gooHttp.NewRequest().Post(URL_ORDER_QUERY, buf)
	if err != nil {
		return nil, err
	}

	gooLog.Debug("[QueryOrderResponse.xml]", string(rstBuf))

	rsp := &QueryOrderResponse{}
	if err := xml.Unmarshal(rstBuf, rsp); err != nil {
		return nil, err
	}
	if rsp.ReturnCode == FAIL {
		return nil, errors.New(rsp.ReturnMsg)
	}
	if rsp.ResultCode == FAIL {
		return nil, errors.New(rsp.ErrCodeDes)
	}
	if rsp.TradeState != SUCCESS {
		return nil, errors.New(tradeStateMsg[rsp.TradeState])
	}

	return rsp, nil
}
