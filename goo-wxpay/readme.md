# 微信支付

golang 实现的微信支付SDK

## 统一下单

### 微信H5&小程序

```
req := &gooWXPay.UnifiedOrderRequest{
    Appid:      APP_ID,
    MchId:      MCH_ID,
    Body:       "购买小风扇一台",
    OutTradeNo: fmt.Sprintf("%d", time.Now().UnixNano()<<2),
    TotalFee:   1,
    NotifyUrl:  NOTIFY_URL,
    TradeType:  gooWXPay.TRADE_TYPE_JSAPI,
    Openid:     OPENID,
}

rsp, err := gooWXPay.UnifiedOrder(req, API_KEY)

fmt.Println(rsp, err)
```

### APP

```
req := &gooWXPay.UnifiedOrderRequest{
    Appid:      APP_ID,
    MchId:      MCH_ID,
    Body:       "购买小风扇一台",
    OutTradeNo: fmt.Sprintf("%d", time.Now().UnixNano()<<2),
    TotalFee:   1,
    NotifyUrl:  NOTIFY_URL,
    TradeType:  gooWXPay.TRADE_TYPE_APP,
    Openid:     OPENID,
}

rsp, err := gooWXPay.UnifiedOrder(req, API_KEY)

fmt.Println(rsp, err)
```

## 订单查询

```
req := &gooWXPay.QueryOrderRequest{
    Appid:      APP_ID,
    MchId:      MCH_ID,
    OutTradeNo: "6355074728607893693",
}

rsp, err := gooWXPay.QueryOrder(req, API_KEY)

fmt.Println(rsp, err)
```

## 申请退款

```
req := &gooWXPay.RefundRequest{
    Appid:       APP_ID,
    MchId:       MCH_ID,
    OutTradeNo:  "9369635507837472860",
    OutRefundNo: fmt.Sprintf("%d", time.Now().Unix()<<2),
    TotalFee:    1,
    RefundFee:   1,
    RefundDesc:  "申请退款",
    NotifyUrl:   NOTIFY_REFUND_URL,
}

rsp, err := gooWXPay.Refund(req, API_KEY,
    "./cert/apiclient_cert.pem",
    "./cert/apiclient_key.pem")
    
fmt.Println(rsp, err)
```

## 退款查询

```
req := &gooWXPay.RefundQueryRequest{
    Appid:       APP_ID,
    MchId:       MCH_ID,
    OutTradeNo:  "9369376355078472860",
}

rsp, err := gooWXPay.RefundQuery(req, API_KEY)

fmt.Println(rsp, err)
```

## 支付结果回调

```
rsp, err := gooWXPay.OrderPayNotify(buf, APP_ID)

fmt.Println(rsp, err)
```

## 退款结果回调

```
rsp, err := gooWXPay.RefundNotify(buf, APP_ID)

fmt.Println(rsp, err)
```

## 回调成功

```
gooWXPay.Success()
```