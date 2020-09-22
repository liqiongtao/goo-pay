package gooLog

// 日志级别
type Level int

// 日志级别 定义
const (
	LEVEL_DEBUG Level = iota // 调试输出
	LEVEL_INFO               // 信息输出
	LEVEL_WARN               // 警告输出
	LEVEL_ERROR              // 错误输出
)

// 日志级别 文本
var LevelText = map[Level]string{
	LEVEL_DEBUG: "DEBUG",
	LEVEL_INFO:  "INFO",
	LEVEL_WARN:  "WARN",
	LEVEL_ERROR: "ERROR",
}

// 默认
const (
	DEFAULT_LEVEL = LEVEL_DEBUG
)
