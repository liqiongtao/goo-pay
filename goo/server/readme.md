# gooServer

基于gin框架，增加了日志、跨域等中间，增加了统一返回处理

## 使用介绍

```
g := goo.Gin()

g.Use(func(c *gin.Context) {
    c.Next()
})

app := g.Group('/app')
app.Use(verify(), token(), authorize())
{
    app.Get('ping', goo.Handler(controller.Ping{}))
}

g.Serve(":8080")
```

## gin.go

对gin进行一层包装

### GinEngine

```
type GinEngine struct {
	*gin.Engine
	noLogPaths map[string]struct{}
	requestId  int64
}
```

- Engine: gin.Engine 引擎
- noLogPaths: 用于过滤不打印日志的uri
- requestId: 用于每次请求计数

### Gin引擎

```
func Gin() *GinEngine
```

## middleware.go

- 请求日志

```
func logger(g *GinEngine) gin.HandlerFunc
```

- 异常捕获&统一输出

```
func recovery() gin.HandlerFunc
```

- 跨域处理

```
func cors() gin.HandlerFunc
```

- 不处理的请求

```
func noAccess() gin.HandlerFunc
```

- 404请求

```
func noRoute() gin.HandlerFunc
```

## controller.go

### 抽象接口

定义抽象接口，每个request定义为struct结构体，结构体参数即请求参数

```
type iController interface {
	DoHandle(c *gin.Context)
}
```

### 接口实现

```
func Handler(controller iController) gin.HandlerFunc {
	return controller.DoHandle
}
```

## response.go

### Response

- Code: 返回编码，0表示成功，非0表示失败。（采用"没有消息就是好消息"机制）
- Message: 成功或失败描述信息
- Data: 返回数据，泛型
- 返回信息转换为JSON字符串

```
func (rsp Response) ToString() string
```

### 抛异常

程序执行到逻辑异常时，直接抛异常，可减少程序嵌套return处理

```
func Exception(code int, message string)
```

### 成功

```
func Success(data interface{})
```