package gooQrcode

import (
	"errors"
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

func (this *Qrcode) Get() ([]byte, error) {
	if this.Url == "" {
		return nil, errors.New("url 为空")
	}

	if this.Size == 0 {
		this.Size = DEFAULT_SIZE
	}

	this.Url, _ = url.QueryUnescape(this.Url)

	png, err := qrcode.Encode(this.Url, qrcode.Medium, this.Size)
	if err != nil {
		return nil, err
	}

	return png, nil
}

func (this *Qrcode) Output(c gin.Context) error {
	png, err := this.Get()
	if err != nil {
		return err
	}

	c.Header("Content-Type", "image/png")
	c.Writer.Write(png)

	return nil
}
