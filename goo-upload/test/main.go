package main

import (
	"googo.io/goo-sms"
	"log"
)

func main() {
	conf := gooSms.AliyunConfig{
		Region:       "",
		Appid:        "",
		Secret:       "",
		SignName:     "",
		TemplateCode: "",
	}
	err := gooSms.New(gooSms.Aliyun(conf)).Verify("18512345678", "mob-login", "1234")
	if err != nil {
		log.Println(err.Error())
		return
	}
}
