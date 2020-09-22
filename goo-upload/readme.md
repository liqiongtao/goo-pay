# 发送短信

```
conf := gooSms.AliyunConfig{
    Region:       "",
    Appid:        "",
    Secret:       "",
    SignName:     "",
    TemplateCode: "",
}
code, err := gooSms.New(gooSms.Aliyun(conf)).Send("18512345678", "mob-login")
if err != nil {
    log.Println(err.Error())
    return
}
log.Println(code)
```

# 验证短信验证码

```
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
```