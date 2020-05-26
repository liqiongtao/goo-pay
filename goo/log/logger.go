package gooLog

import (
	"fmt"
	"runtime"
	"strings"
	"sync"
)

type Logger struct {
	Writer   iWriter
	MinLevel Level
	sync.Mutex
}

func New(w iWriter) *Logger {
	l := &Logger{
		Writer:   w,
		MinLevel: DEFAULT_LEVEL,
	}
	l.Writer.Init()
	return l
}

func Default() *Logger {
	return New(&File{})
}

func (l *Logger) Info(v ...interface{}) {
	l.output(LEVEL_INFO, v...)
}

func (l *Logger) Debug(v ...interface{}) {
	l.output(LEVEL_DEBUG, v...)
}

func (l *Logger) Warn(v ...interface{}) {
	l.output(LEVEL_WARN, v...)
}

func (l *Logger) Error(v ...interface{}) {
	l.output(LEVEL_ERROR, v...)
}

func (l *Logger) output(level Level, v ...interface{}) {
	if level < l.MinLevel {
		return
	}

	l.Mutex.Lock()
	defer l.Mutex.Unlock()

	info := &LogInfo{
		Level: level,
		Data:  v,
		Trace: []string{},
	}

	if level >= LEVEL_ERROR {
		for i := 4; i < 12; i++ {
			if _, file, line, ok := runtime.Caller(i); ok {
				filename := strings.Split(file, "src/")[1]
				info.Trace = append(info.Trace, fmt.Sprintf("%s %dL", filename, line))
			}
		}
	}

	l.Writer.Output(info)
}
