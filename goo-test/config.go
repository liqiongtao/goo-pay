package main

import (
	gooSms "googo.io/goo-sms"
	gooUpload "googo.io/goo-upload"
	gooWeixin "googo.io/goo-weixin"
	"googo.io/goo/cache"
	gooDB "googo.io/goo/db"
	"time"
)

type config struct {
	Env    string `yaml:"env"`
	Server server `yaml:"server"`

	Redis gooCache.Config     `yaml:"redis"`
	Mysql gooDB.Config        `yaml:"mysql"`
	Oss   gooUpload.OssConfig `yaml:"oss"`
	Sms   gooSms.AliyunConfig `yaml:"sms"`

	Weixin wx    `yaml:"weixin"`
	WxPay  wxpay `yaml:"wxpay"`
}

type server struct {
	Port      string            `yaml:"port"`
	Accounts  map[string]string `yaml:"accounts"`
	AppKey    string            `yaml:"app_key"`
	AppSecret string            `yaml:"app_secret"`
	ExpireIn  time.Duration     `yaml:"expire_in"`
}

type wx struct {
	Cache   gooCache.Config  `yaml:"cache"`
	Android gooWeixin.Config `yaml:"android"`
	Ios     gooWeixin.Config `yaml:"ios"`
	Minip   gooWeixin.Config `yaml:"minip"`
	Wechat  gooWeixin.Config `yaml:"wechat"`
}

type wxpay struct {
	MchId     string `yaml:"mch_id"`
	ApiKey    string `yaml:"apikey"`
	NotifyUrl string `yaml:"notify_url"`
}
