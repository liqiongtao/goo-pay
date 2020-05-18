package gooWXPay

import (
	"encoding/xml"
	"errors"
	gooUtils "goo/utils"
	"strings"
)

type OrderPayData struct {
	ReturnCode string `xml:"return_code"`
	ReturnMsg  string `xml:"return_msg"`

	ResultCode string `xml:"result_code"`
	ErrCode    string `xml:"err_code"`
	ErrCodeDes string `xml:"err_code_des"`

	Appid      string   `xml:"appid"`
	MchId      string   `xml:"mch_id"`
	DeviceInfo string   `xml:"device_info"`
	NonceStr   string   `xml:"nonce_str"`
	Sign       string   `xml:"sign"`
	SignType   SignType `xml:"sign_type"`

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
	TransactionId      string    `xml:"transaction_id"`
	OutTradeNo         string    `xml:"out_trade_no"`
	Attach             string    `xml:"attach"`
	TimeEnd            string    `xml:"time_end"`
}

func OrderPayNotify(buf []byte, apiKey string) (*OrderPayData, error) {
	data := &OrderPayData{}

	if err := xml.Unmarshal(buf, data); err != nil {
		return nil, err
	}
	if data.ReturnCode == FAIL {
		return nil, errors.New(data.ReturnMsg)
	}
	if data.ResultCode == FAIL {
		return nil, errors.New(data.ErrCodeDes)
	}

	params := xml2map(buf)
	str := map2querystring(params) + "&key=" + apiKey

	var signStr string
	if data.SignType == SIGN_TYPE_MD5 {
		signStr = strings.ToUpper(gooUtils.MD5([]byte(str)))
	} else {
		signStr = strings.ToUpper(gooUtils.HMacSha256([]byte(str), []byte(apiKey)))
	}

	if signStr != data.Sign {
		return nil, errors.New("签名验证失败")
	}

	return data, nil
}
