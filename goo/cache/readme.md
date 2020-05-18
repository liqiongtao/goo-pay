# gooCache

## 初始化配置

```
gooCache.Init(cf)
```

## 获取连接对象

```
client := gooCache.Redis()
```

## 示例

### set

```
gooCache.Redis().Set("name", "hnatao", 0).Err()
```

### get

```
gooCache.Redis().Get("name").String()
```