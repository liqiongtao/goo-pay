package gooHttp

import (
	"bytes"
	gooLog "goo/log"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
)

func NewRequest() *Request {
	return &Request{
		Headers: map[string]string{
			"Content-Type": CONTENT_TYPE_FORM,
		},
	}
}

func NewTlsRequest(caCrtFile, clientCrtFile, clientKeyFile string) *Request {
	return &Request{
		Headers: map[string]string{
			"Content-Type": CONTENT_TYPE_FORM,
		},
		Tls: &Tls{
			CaCrtFile:     caCrtFile,
			ClientCrtFile: clientCrtFile,
			ClientKeyFile: clientKeyFile,
		},
	}
}

func Get(url string) ([]byte, error) {
	return NewRequest().Get(url)
}

func Post(url string, data []byte) ([]byte, error) {
	return NewRequest().Post(url, data)
}

func Upload(url, field, file string, data map[string]string) ([]byte, error) {
	fh, err := os.Open(file)
	if err != nil {
		gooLog.Error(err.Error())
		return nil, err
	}
	defer fh.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile(field, filepath.Base(file))
	if err != nil {
		gooLog.Error(err.Error())
		return nil, err
	}
	if _, err = io.Copy(part, fh); err != nil {
		gooLog.Error(err.Error())
		return nil, err
	}

	for k, v := range data {
		writer.WriteField(k, v)
	}

	if err = writer.Close(); err != nil {
		gooLog.Error(err.Error())
		return nil, err
	}

	request := NewRequest()
	request.SetHearder("Content-Type", writer.FormDataContentType())
	return request.Do("POST", url, body)
}
