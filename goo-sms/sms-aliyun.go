package gooSms

import (
	"encoding/json"
	"errors"
	"fmt"
	gooLog "googo.io/goo/log"
	"math/rand"
	"strconv"
	"time"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
)

type SmsAliyun struct {
	Config AliyunConfig
	Client *sdk.Client
}

func (this *SmsAliyun) Init() {
	this.Client, _ = sdk.NewClientWithAccessKey("cn-hangzhou", this.Config.Appid, this.Config.Secret)
}

func (this *SmsAliyun) Send(mobile, action string) (string, error) {
	code := this.getSmsCode(mobile, action)

	request := requests.NewCommonRequest()

	request.Method = "POST"
	request.Scheme = "https"
	request.Domain = "dysmsapi.aliyuncs.com"
	request.Version = "2017-05-25"
	request.ApiName = "SendSms"

	request.QueryParams["RegionId"] = "cn-hangzhou"
	request.QueryParams["PhoneNumbers"] = mobile
	request.QueryParams["SignName"] = this.Config.SignName
	request.QueryParams["TemplateCode"] = this.Config.TemplateCode
	request.QueryParams["TemplateParam"] = fmt.Sprintf("{\"code\": %s}", code)

	response, err := this.Client.ProcessCommonRequest(request)
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

	__cache.setCode(this.Config.Appid, mobile, action, code, expireIn*time.Second)

	return code, nil
}

func (this *SmsAliyun) Verify(mobile, action, code string) error {
	__code := __cache.getCode(this.Config.Appid, mobile, action)
	if __code == "" {
		return errors.New("验证码无效")
	}
	if __code != code {
		return errors.New("验证码错误")
	}
	return nil
}

func (this *SmsAliyun) getSmsCode(mobile, action string) string {
	rand.Seed(time.Now().UnixNano())
	return strconv.Itoa(rand.Intn(8)+1) +
		strconv.Itoa(rand.Intn(8)+1) +
		strconv.Itoa(rand.Intn(8)+1) +
		strconv.Itoa(rand.Intn(8)+1) +
		strconv.Itoa(rand.Intn(8)+1) +
		strconv.Itoa(rand.Intn(8)+1)
}
