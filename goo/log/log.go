package gooLog

var l = Default()

func Info(v ...interface{}) {
	l.Info(v...)
}

func Debug(v ...interface{}) {
	l.Debug(v...)
}

func Warn(v ...interface{}) {
	l.Warn(v...)
}

func Error(v ...interface{}) {
	l.Error(v...)
}
