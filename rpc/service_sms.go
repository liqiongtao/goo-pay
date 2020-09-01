package main

import (
	"context"
	rpc_goo "googo.io/rpc/protos.pb"
)

type SmsService struct {
}

func (*SmsService) Send(ctx context.Context, in *rpc_goo.SmsSendParams) (*rpc_goo.SmsSendResponse, error) {

}

func (*SmsService) Verify(ctx context.Context, in *rpc_goo.SmsVerifyParams) (*rpc_goo.Response, error) {

}
