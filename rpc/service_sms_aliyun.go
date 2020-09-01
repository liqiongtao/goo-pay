package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	gooLog "googo.io/goo/log"
	"math/rand"
	"strconv"
	"sync"
	"time"
)

var (
	smsAliyunCodes         sync.Map
	smsAliyunCodeExpireOut sync.Map
	smsAliyunClients       sync.Map
	expireIn               = 3600
)

func init() {
	go func() {
		for {
			if isQuit {
				break
			}

			smsAliyunCodeExpireOut.Range(func(key, expireOut interface{}) bool {
				if time.Now().Unix() > expireOut.(int64) {
					smsAliyunCodeExpireOut.Delete(key)
					smsAliyunCodes.Delete(key)
				}
				return true
			})

			time.Sleep(time.Duration(expireIn) * time.Second)
		}
	}()
}

type smsAliyun struct {
	Appid        string
	Secret       string
	SignName     string
	TemplateCode string
	client       *sdk.Client
}

func (this *smsAliyun) Send(mobile, action string) (string, error) {
	code := this.getSmsCode(mobile, action)

	request := requests.NewCommonRequest()

	request.Method = "POST"
	request.Scheme = "https"
	request.Domain = "dysmsapi.aliyuncs.com"
	request.Version = "2017-05-25"
	request.ApiName = "SendSms"

	request.QueryParams["RegionId"] = "cn-hangzhou"
	request.QueryParams["PhoneNumbers"] = mobile
	request.QueryParams["SignName"] = this.SignName
	request.QueryParams["TemplateCode"] = this.TemplateCode
	request.QueryParams["TemplateParam"] = fmt.Sprintf("{\"code\": %s}", code)

	response, err := this.client.ProcessCommonRequest(request)
	if err != nil {
		gooLog.Error("[aliyun-sms]", "[send-error]", request.String(), err.Error())
		return "", err
	}

	rsp := map[string]string{}
	if err := json.Unmarshal(response.GetHttpContentBytes(), &rsp); err != nil {
		gooLog.Error("[aliyun-sms]", "[send-error]", request.String(), err.Error())
		return "", err
	}

	if rsp["Code"] != "OK" {
		gooLog.Error("[aliyun-sms]", "[send-error]", request.String(), rsp)
		return "", errors.New(rsp["Message"])
	}

	key := fmt.Sprintf("%s_%s_%s_%s", this.Appid, mobile, action, code)
	smsAliyunCodes.Store(key, code)
	smsAliyunCodeExpireOut.Store(key, time.Now().Unix()+int64(expireIn))

	return code, nil
}

func (this *smsAliyun) Verify(mobile, action, code string) error {
	key := fmt.Sprintf("%s_%s_%s_%s", this.Appid, mobile, action, code)

	__code, ok := smsAliyunCodes.Load(key)
	if !ok {
		return errors.New("验证码无效")
	}
	if __code != code {
		return errors.New("验证码错误")
	}

	expireOut, _ := smsAliyunCodeExpireOut.Load(key)
	if time.Now().Unix() > expireOut.(int64) {
		return errors.New("验证码失效")
	}

	return nil
}
func (this *smsAliyun) getClient() error {
	if c, ok := smsAliyunClients.Load(this.Appid); ok {
		this.client = c.(*sdk.Client)
		return nil
	}

	client, err := sdk.NewClientWithAccessKey("cn-hangzhou", this.Appid, this.Secret)
	if err != nil {
		return err
	}

	smsAliyunClients.Store(this.Appid, client)
	this.client = client
	return nil
}

func (this *smsAliyun) getSmsCode(mobile, action string) string {
	rand.Seed(time.Now().UnixNano())
	return strconv.Itoa(rand.Intn(8)+1) +
		strconv.Itoa(rand.Intn(8)+1) +
		strconv.Itoa(rand.Intn(8)+1) +
		strconv.Itoa(rand.Intn(8)+1) +
		strconv.Itoa(rand.Intn(8)+1) +
		strconv.Itoa(rand.Intn(8)+1)
}
