# 获取验证码

```
rst := gooCaptcha.Get(240, 80)
```

返回数据结构

```
{
    "id": "",
    "base64image": "",
}
```

# 校验验证码

```
err := gooCaptcha.Verify("{id}", "{code}")
if err != nil {
}
```