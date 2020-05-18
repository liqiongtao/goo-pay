# 发送请求

## get

```
bts, err := gooHttp.Get("http://api.help.bj.cn/apis/weather6d/?id=101010200")

gooLog.Debug(string(bts), err)
```

## post

```
values := url.Values{}
values.Add("id", "101010200")

bts, err := gooHttp.Post("http://api.help.bj.cn/apis/weather6d/", []byte(values.Encode()))

gooLog.Debug(string(bts), err)
```

## upload

```
bts, err := gooHttp.Upload("https://s.weflys.com/upload/oss/5087ab15eb8e12ea", "file", "100.txt", nil)

gooLog.Debug(string(bts), err)
```

## tls

```
values := url.Values{}
values.Add("id", "100")

bts, err := gooHttp.NewTlsRequest("", "client.crt", "key.crt").Post("https://abc.com", []byte(values.Encode()))

gooLog.Debug(string(bts), err)
```

