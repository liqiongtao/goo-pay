package main

import (
	"context"
	"encoding/base64"
	"fmt"
	"github.com/skip2/go-qrcode"
	rpc_goo "googo.io/rpc/protos.pb"
	"net/url"
)

type QRCodeService struct {
}

func (*QRCodeService) Get(ctx context.Context, in *rpc_goo.QRCodeGetParams) (*rpc_goo.QRCodeGetResponse, error) {
	if in.Url == "" {
		return &rpc_goo.QRCodeGetResponse{ErrCode: 30021, ErrMsg: "url is null"}, nil
	}

	if in.Size == 0 {
		in.Size = 256
	}

	urlStr, err := url.QueryUnescape(in.Url)
	if err != nil {
		return &rpc_goo.QRCodeGetResponse{ErrCode: 30022, ErrMsg: err.Error()}, nil
	}

	pngBuf, err := qrcode.Encode(urlStr, qrcode.Medium, int(in.Size))
	if err != nil {
		return &rpc_goo.QRCodeGetResponse{ErrCode: 30023, ErrMsg: err.Error()}, nil
	}

	base64image := fmt.Sprintf("data:%s;base64,%s", "image/png", base64.StdEncoding.EncodeToString(pngBuf))
	return &rpc_goo.QRCodeGetResponse{Base64Image: base64image}, nil
}
