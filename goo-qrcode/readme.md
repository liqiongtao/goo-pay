# 生成二维码

- 二维码字节

```
buf, err := gooQrcode.New("http://googo.io").Get()
```

- 二维码base64

```
b64img, err := gooQrcode.New("http://googo.io").Base64Image()
if err != nil {
}
```

- 二维码图片

```
err := gooQrcode.New("http://googo.io").Output(c)
if err != nil {
}
```