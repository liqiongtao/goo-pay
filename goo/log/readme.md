# 日志

## 日志级别

- info
- debug
- warn
- error

## 输出内容

- 日期时间，格式: 0000-00-00 00:00:00
- 级别
- 日志内容
- 文件名+行号

```
0000-00-00 00:00:00 DEBUG xxxxx
```

## 默认日志对象

```
l := Default()
```

## 新建日志对象

### 输出到文件

```
l := New(&File{})
```

### 输出到控制台

```
l := New(&Console{})
```

## 示例

```
gooLog.Info("this is info log")
gooLog.Debug("this is debug log")
gooLog.Warn("this is warn log")
gooLog.Error("this is error log")

gooLog.Debug("name", "liqiongao")
gooLog.Debug("addr", []string{"北京", "北京"})
```

# 设计方案

## iWriter

定义输出接口对象，具体实现对象: File, Console

### Init()

初始化

### Output(info *LogInfo)

输出日志内容

## File

实现 iWriter

文件输出：logs/20060102.log

### Dir

自定义输出目录

### FileName

自定义输出日志文件名

## Console

实现 iWriter，输出到控制台

## LogInfo

### Level

日志级别

### Data

[]interface{} 类型，用于存放日志内容

### Trace

[]string 类型，用于存放文件名、行号

## Logger

### Writer

抽象"输出接口"对象

### MinLevel

定义最新输出日志级别，小于该级别的日志不再输出，默认DEBUG

### sync.Mutex

对象内输出日志产生并发时，加锁串行输出

### output(level Level, v ...interface{})

统一处理日志输出任务，使用 ```l.Writer.Output(info)``` 抽象输出
