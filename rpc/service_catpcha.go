package main

import (
	"context"
	"github.com/mojocn/base64Captcha"
	rpc_goo "googo.io/rpc/protos.pb"
)

type CaptchaService struct {
}

func (*CaptchaService) Get(ctx context.Context, in *rpc_goo.CaptchaGetParams) (*rpc_goo.CaptchaGetResponse, error) {
	if in.GetWidth() == 0 {
		in.Width = 240
	}
	if in.GetHeight() == 0 {
		in.Height = 80
	}

	var configDigit = base64Captcha.ConfigDigit{
		Height:     int(in.Height),
		Width:      int(in.Width),
		MaxSkew:    0.7,
		DotCount:   80,
		CaptchaLen: 6,
	}

	id, cap := base64Captcha.GenerateCaptcha("", configDigit)

	rsp := rpc_goo.CaptchaGetResponse{
		Id:          id,
		Base64Image: base64Captcha.CaptchaWriteToBase64Encoding(cap),
	}

	return &rsp, nil
}

func (*CaptchaService) Verify(ctx context.Context, in *rpc_goo.CaptchaVerifyParams) (*rpc_goo.Response, error) {
	if in.GetId() == "" {
		return &rpc_goo.Response{ErrCode: 30011, ErrMsg: "图片验证码ID为空"}, nil
	}

	if in.GetCode() == "" {
		return &rpc_goo.Response{ErrCode: 30012, ErrMsg: "图片验证码为空"}, nil
	}

	rst := base64Captcha.VerifyCaptcha(in.GetId(), in.GetCode())
	if !rst {
		return &rpc_goo.Response{ErrCode: 30013, ErrMsg: "图片验证码错误"}, nil
	}

	return nil, nil
}
