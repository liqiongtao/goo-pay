package gooWXPay

const (
	// 统一下单
	URL_UNIFIED_ORDER = "https://api.mch.weixin.qq.com/pay/unifiedorder"

	// 查询订单
	URL_ORDER_QUERY = "https://api.mch.weixin.qq.com/pay/orderquery"

	// 关闭订单
	URL_CLOSE_QUERY = "https://api.mch.weixin.qq.com/pay/closeorder"

	// 申请退款
	URL_REFUND = "https://api.mch.weixin.qq.com/secapi/pay/refund"

	// 查询退款
	URL_REFUND_QUERY = "https://api.mch.weixin.qq.com/pay/refundquery"
)

const (
	SUCCESS     = "SUCCESS"     // 成功
	FAIL        = "FAIL"        // 失败
	REFUND      = "REFUND"      // 转入退款
	NOTPAY      = "NOTPAY"      // 未支付
	CLOSED      = "CLOSED"      // 已关闭
	REVOKED     = "REVOKED"     // 已撤销（刷卡支付）
	USERPAYING  = "USERPAYING"  // 用户支付中
	PAYERROR    = "PAYERROR"    // 支付失败(其他原因，如银行返回失败)
	CHANGE      = "CHANGE"      // 退款异常
	REFUNDCLOSE = "REFUNDCLOSE" // 退款关闭
)

var tradeStateMsg = map[string]string{
	SUCCESS:    "支付成功",
	FAIL:       "成功",
	REFUND:     "转入退款",
	NOTPAY:     "未支付",
	CLOSED:     "已关闭",
	REVOKED:    "已撤销（刷卡支付）",
	USERPAYING: "用户支付中",
	PAYERROR:   "支付失败(其他原因，如银行返回失败)",
}

var tradeStatusMsg = map[string]string{
	SUCCESS:     "退款成功",
	CHANGE:      "退款异常",
	REFUNDCLOSE: "退款关闭",
}

// 签名类型
type SignType string

const (
	SIGN_TYPE_MD5         SignType = "MD5"
	SIGN_TYPE_HMAC_SHA256 SignType = "HMAC-SHA256"
)

// 交易类型
type TradeType string

const (
	TRADE_TYPE_JSAPI    TradeType = "JSAPI"    // 微信h5, 小程序
	TRADE_TYPE_NATIVE   TradeType = "NATIVE"   // native
	TRADE_TYPE_APP      TradeType = "APP"      // android, ios
	TRADE_TYPE_MICROPAY TradeType = "MICROPAY" // 付款码支付
)
