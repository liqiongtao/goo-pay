package gooUpload

import (
	"bytes"
	"fmt"
	gooLog "goo/log"
	gooUtils "goo/utils"
	"path"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

var Oss *gooOss

func InitOss(config OssConfig) {
	var err error
	Oss, err = NewOss(config)
	if err != nil {
		gooLog.Error(err.Error())
	}
}

func NewOss(config OssConfig) (*gooOss, error) {
	this := &gooOss{
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

type gooOss struct {
	Config OssConfig
	Client *oss.Client
	Bucket *oss.Bucket
}

func (this *gooOss) Upload(filename string, body []byte) (string, error) {
	md5str := gooUtils.MD5(body)
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

func (this *gooOss) getClient() (*oss.Client, error) {
	return oss.New(this.Config.Endpoint, this.Config.AccessKeyId, this.Config.AccessKeySecret)
}

func (this *gooOss) getBucket() (*oss.Bucket, error) {
	return this.Client.Bucket(this.Config.Bucket)
}
