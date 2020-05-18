# gooDB

## 初始化配置

```
gooDB.Init(cf)
```

## 获取连接对象

```
db := gooDB.Orm()
```

## 示例

### 查询

```
u := &User{}
has, err := gooDB.Orm().Where("name", "hnatao").Get(u)
fmt.Println(has, err, u)
```

### 添加

```
u := &User{
    Name: "hantao",
}
_, err := gooDB.Orm().Insert()
fmt.Println(err, u)
```

### 修改

```
u := &User{
    Name: "hantao",
}
_, err := gooDB.Orm().ID(1).Update(u)
fmt.Println(err)
```

### 删除

```
u := &User{}
_, err := gooDB.Orm().ID(1).Delete(u)
fmt.Println(err)
```