# Goo框架

这是一个开发框架，整合使用第三方成熟组件，
比如：server使用gin，数据库使用xorm，缓存使用go-redis。

在此基础上，本人整理多年开发经验，将常用组件、方法、类库、日志等统一整理，已达到快速开发目的。

# server.go

```
使用Gin框架，封装web服务，提供异常处理，成功处理。
封装goroutine方法
实现controller抽象方法
```

# params.go

```
定义 Params 抽象对象
实现 Params -> Json
实现 Params -> Xml
实现 Params -> QueryString
实现 快速获取类型属性，如：GetString(), GetBool(), GetInt(), GetInt64()
```

# db.go

```
封装Mysql对象，供应用系统全局使用

初始化:
goo.InitDB(conf)

使用:
u := &User{}
has, err := gooDB.Orm().Where("name", "hnatao").Get(u)
fmt.Println(has, err, u)
```

# redis.go

```
封装redis对象，供应用系统全局使用

初始化:
goo.InitRedis(conf)

使用:
goo.Redis().Set("name", "liqiongtao").Err()
gooCache.Redis().Get("name").String()
```

# config

```
加载ymal配置文件

使用:
gooConfig.LoadFile(".yaml", &conf)
```

# http

```
封装http请求，实现了get、post、upload、tls

使用:
gooHttp.Get()
gooHttp.Post()
gooHttp.Upload()
gooHttp.NewTlsRequest().Post()
```

# log

```
自定义日志输出，默认输出到文件logs/20200101.log

使用:
gooLog.Info("hi")
gooLog.Debug("hi", 100, []string{"liqiongtao"})
gooLog.Warn("hi")
gooLog.Error("hi", map[string]interface{}{"errCode":500,"errMsg":"err"})
```

# utils

```
常用函数

大数字计算：
gooUtils.BigIntAdd(10000000, 100) // 加
gooUtils.BigIntReduce(10000000, 100) // 减
gooUtils.BigIntMul(10000000, 100) // 乘
gooUtils.BigIntDiv(10000000, 100) // 除
gooUtils.BigIntMod(10000000, 100) // 取模
gooUtils.BigIntCmp(10000000, 100) // 比较

随机数:
gooUtils.NonceStr()

id2code:
gooUtils.Id2Code(1000)
gooUtils.Code2Id("")

常用加密:
gooUtils.MD5()
gooUtils.SHA1()
gooUtils.SHA256()
gooUtils.HMacMd5()
gooUtils.HMacSha1()
gooUtils.HMacSha256()
gooUtils.Base64Encode()
gooUtils.Base64Decode()
gooUtils.AES256ECBEncrypt()
gooUtils.AES256ECBDecrypt()
gooUtils.Encrypt()
gooUtils.Decrypt()
gooUtils.SHAwithRSA()
```