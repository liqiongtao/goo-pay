package gooUpload

import (
	"bytes"
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"log"
	"path"
)

var OSS *gooOSS

func InitOSS(config OSSConfig) {
	var err error
	OSS, err = NewOSS(config)
	if err != nil {
		log.Panic(err.Error())
	}
}

func NewOSS(config OSSConfig) (*gooOSS, error) {
	this := &gooOSS{
		Config: config,
	}

	client, err := this.getClient()
	if err != nil {
		return nil, err
	}

	this.Client = client

	bucket, err := this.getBucket()
	if err != nil {
		return nil, err
	}

	this.Bucket = bucket

	return this, nil
}

type gooOSS struct {
	Config OSSConfig
	Client *oss.Client
	Bucket *oss.Bucket
}

func (this *gooOSS) Upload(filename string, body []byte) (string, error) {
	md5str := goo.Util.MD5(body)
	filename = fmt.Sprintf("%s/%s/%s%s", md5str[0:2], md5str[2:4], md5str[8:24], path.Ext(filename))

	if err := this.Bucket.PutObject(filename, bytes.NewReader(body)); err != nil {
		return "", err
	}

	if this.Config.Domain != "" {
		return this.Config.Domain + filename, nil
	}

	url := "https://" + this.Config.Bucket + "." + this.Config.Endpoint + "/" + filename
	return url, nil
}

func (this *gooOSS) getClient() (*oss.Client, error) {
	return oss.New(this.Config.Endpoint, this.Config.AccessKeyId, this.Config.AccessKeySecret)
}

func (this *gooOSS) getBucket() (*oss.Bucket, error) {
	return this.Client.Bucket(this.Config.Bucket)
}
