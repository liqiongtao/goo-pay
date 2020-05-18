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

type RefundQueryRequest struct {
	Appid         string   `xml:"appid"`
	MchId         string   `xml:"mch_id"`
	NonceStr      string   `xml:"nonce_str"`
	Sign          string   `xml:"sign"`
	SignType      SignType `xml:"sign_type"`
	TransactionId string   `xml:"transaction_id"`
	OutTradeNo    string   `xml:"out_trade_no"`
	OutRefundNo   string   `xml:"out_refund_no"`
	RefundId      string   `xml:"refund_id"`
	Offset        int64    `xml:"offset"`
}

func (r *RefundQueryRequest) toXml(apiKey string) []byte {
	if r.NonceStr == "" {
		r.NonceStr = gooUtils.NonceStr()
	}
	if r.SignType == "" {
		r.SignType = SIGN_TYPE_HMAC_SHA256
	}

	str := obj2querystring(r) + fmt.Sprintf("&key=%s", apiKey)
	gooLog.Debug("[RefundQueryRequest.querystring]", str)

	if r.SignType == SIGN_TYPE_HMAC_SHA256 {
		r.Sign = strings.ToUpper(gooUtils.HMacSha256([]byte(str), []byte(apiKey)))
	} else if r.SignType == SIGN_TYPE_MD5 {
		r.Sign = strings.ToUpper(gooUtils.MD5([]byte(str)))
	}

	return obj2xml(r)
}

type RefundQueryResponse struct {
	ReturnCode string `xml:"return_code"`
	ReturnMsg  string `xml:"return_msg"`

	ResultCode string `xml:"result_code"`
	ErrCode    string `xml:"err_code"`
	ErrCodeDes string `xml:"err_code_des"`

	Appid    string `xml:"appid"`
	MchId    string `xml:"mch_id"`
	NonceStr string `xml:"nonce_str"`
	Sign     string `xml:"sign"`

	TotalRefundCount   int64  `xml:"total_refund_count"`
	TransactionId      string `xml:"transaction_id"`
	OutTradeNo         string `xml:"out_trade_no"`
	TotalFee           int64  `xml:"total_fee"`
	SettlementTotalFee int64  `xml:"settlement_total_fee"`
	CashFee            int64  `xml:"cash_fee"`
	RefundCount        int64  `xml:"refund_count"`

	OutRefundNo         string `xml:"out_refund_no_0"`
	RefundId            string `xml:"refund_id_0"`
	RefundChannel       int64  `xml:"refund_channel_0"`
	RefundFee           int64  `xml:"refund_fee_0"`
	SettlementRefundFee int64  `xml:"settlement_refund_fee_0"`
	CouponType          int64  `xml:"coupon_type_0_0"`
	CouponRefundFee     int64  `xml:"coupon_refund_fee_0"`
	CouponRefundCount   int64  `xml:"coupon_refund_count_0"`
	CouponRefundId      int64  `xml:"coupon_refund_id_0_0"`
	CouponRefundFee2    int64  `xml:"coupon_refund_fee_0_0"`
	RefundStatus        int64  `xml:"refund_status_0"`
	RefundAccount       int64  `xml:"refund_account_0"`
	RefundRecvAccout    int64  `xml:"refund_recv_accout_0"`
	RefundSuccessTime   int64  `xml:"refund_success_time_0"`
}

func RefundQuery(req *RefundQueryRequest, apiKey string) (*RefundQueryResponse, error) {
	buf := req.toXml(apiKey)
	gooLog.Debug("[RefundRequest.xml]", string(buf))

	rstBuf, err := gooHttp.NewRequest().Post(URL_REFUND_QUERY, buf)
	if err != nil {
		return nil, err
	}

	gooLog.Debug("[RefundResponse.xml]", string(rstBuf))

	rsp := &RefundQueryResponse{}
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
