package gooQrcode

import (
	"encoding/base64"
	"errors"
	"fmt"
	"net/url"

	"github.com/gin-gonic/gin"
	"github.com/skip2/go-qrcode"
)

func New(url string) *Qrcode {
	return &Qrcode{
		Url: url,
	}
}

type Qrcode struct {
	Size int
	Url  string
}

func (qr *Qrcode) Get() ([]byte, error) {
	if qr.Url == "" {
		return nil, errors.New("url 为空")
	}

	if qr.Size == 0 {
		qr.Size = DEFAULT_SIZE
	}

	qr.Url, _ = url.QueryUnescape(qr.Url)

	png, err := qrcode.Encode(qr.Url, qrcode.Medium, qr.Size)
	if err != nil {
		return nil, err
	}

	return png, nil
}

func (qr *Qrcode) Base64Image() (string, error) {
	png, err := qr.Get()
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("data:image/png;base64,%s", base64.StdEncoding.EncodeToString(png)), nil
}

func (qr *Qrcode) Output(c gin.Context) error {
	png, err := qr.Get()
	if err != nil {
		return err
	}

	c.Header("Content-Type", "image/png")
	c.Writer.Write(png)

	return nil
}
