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

type RefundRequest struct {
	Appid         string   `xml:"appid"`
	MchId         string   `xml:"mch_id"`
	NonceStr      string   `xml:"nonce_str"`
	Sign          string   `xml:"sign"`
	SignType      SignType `xml:"sign_type"`
	TransactionId string   `xml:"transaction_id"`
	OutTradeNo    string   `xml:"out_trade_no"`
	OutRefundNo   string   `xml:"out_refund_no"`
	TotalFee      int64    `xml:"total_fee"`
	RefundFee     int64    `xml:"refund_fee"`
	RefundDesc    string   `xml:"refund_desc"`
	RefundAccount string   `xml:"refund_account"`
	NotifyUrl     string   `xml:"notify_url"`
}

func (r *RefundRequest) toXml(apiKey string) []byte {
	if r.NonceStr == "" {
		r.NonceStr = gooUtils.NonceStr()
	}
	if r.SignType == "" {
		r.SignType = SIGN_TYPE_HMAC_SHA256
	}

	str := obj2querystring(r) + fmt.Sprintf("&key=%s", apiKey)
	gooLog.Debug("[UnifiedOrderRequest.querystring]", str)

	if r.SignType == SIGN_TYPE_HMAC_SHA256 {
		r.Sign = strings.ToUpper(gooUtils.HMacSha256([]byte(str), []byte(apiKey)))
	} else if r.SignType == SIGN_TYPE_MD5 {
		r.Sign = strings.ToUpper(gooUtils.MD5([]byte(str)))
	}

	return obj2xml(r)
}

type RefundResponse struct {
	ReturnCode string `xml:"return_code"`
	ReturnMsg  string `xml:"return_msg"`

	ResultCode string `xml:"result_code"`
	ErrCode    string `xml:"err_code"`
	ErrCodeDes string `xml:"err_code_des"`

	Appid    string `xml:"appid"`
	MchId    string `xml:"mch_id"`
	NonceStr string `xml:"nonce_str"`
	Sign     string `xml:"sign"`

	TransactionId       string `xml:"transaction_id"`
	OutTradeNo          string `xml:"out_trade_no"`
	OutRefundNo         string `xml:"out_refund_no"`
	RefundId            string `xml:"refund_id"`
	RefundFee           int64  `xml:"refund_fee"`
	SettlementRefundFee int64  `xml:"settlement_refund_fee"`
	TotalFee            int64  `xml:"total_fee"`
	SettlementTotalFee  int64  `xml:"settlement_total_fee"`
	CashFee             int64  `xml:"cash_fee"`
	CashRefundFee       int64  `xml:"cash_refund_fee"`
	CouponRefundFee     int64  `xml:"coupon_refund_fee"`
	CouponRefundCount   int64  `xml:"coupon_refund_count"`
}

func Refund(req *RefundRequest, apiKey, clientCrtFile, clientKeyFile string) (*RefundResponse, error) {
	buf := req.toXml(apiKey)
	gooLog.Debug("[RefundRequest.xml]", string(buf))

	rstBuf, err := gooHttp.NewTlsRequest("", clientCrtFile, clientKeyFile).Post(URL_REFUND, buf)
	if err != nil {
		return nil, err
	}

	gooLog.Debug("[RefundResponse.xml]", string(rstBuf))

	rsp := &RefundResponse{}
	if err := xml.Unmarshal(rstBuf, rsp); err != nil {
		return nil, err
	}
	if rsp.ReturnCode == FAIL {
		return nil, errors.New(rsp.ReturnMsg)
	}
	if rsp.ResultCode == FAIL {
		return nil, errors.New(rsp.ErrCodeDes)
	}

	return rsp, nil
}
