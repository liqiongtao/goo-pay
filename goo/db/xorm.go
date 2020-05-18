package gooDB

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	gooLog "googo.io/goo/log"
	"xorm.io/core"
)

func NewXORM(cf Config) *xorm.EngineGroup {
	conns := []string{cf.Master}
	if n := len(cf.Slaves); n > 0 {
		conns = append(conns, cf.Slaves...)
	}

	db, err := xorm.NewEngineGroup(cf.Driver, conns)
	if err != nil {
		panic(err.Error())
	}

	db.SetLogger(&XORMLogger{})

	db.ShowSQL(cf.LogModel)
	db.SetMaxIdleConns(cf.MaxIdle)
	db.SetMaxOpenConns(cf.MaxOpen)

	return db
}

type XORMLogger struct {
	LogLevel core.LogLevel
}

func (l XORMLogger) Debug(v ...interface{}) {
	gooLog.Debug(v...)
}

func (l XORMLogger) Debugf(format string, v ...interface{}) {
	gooLog.Debug(fmt.Sprintf(format, v...))
}

func (l XORMLogger) Error(v ...interface{}) {
	gooLog.Error(v...)
}

func (l XORMLogger) Errorf(format string, v ...interface{}) {
	gooLog.Error(fmt.Sprintf(format, v...))
}

func (l XORMLogger) Info(v ...interface{}) {
	gooLog.Info(v...)
}

func (l XORMLogger) Infof(format string, v ...interface{}) {
	gooLog.Info(fmt.Sprintf(format, v...))
}

func (l XORMLogger) Warn(v ...interface{}) {
	gooLog.Warn(v...)
}

func (l XORMLogger) Warnf(format string, v ...interface{}) {
	gooLog.Warn(fmt.Sprintf(format, v...))
}

func (l XORMLogger) Level() core.LogLevel {
	return l.LogLevel
}

func (l XORMLogger) SetLevel(ll core.LogLevel) {
	l.LogLevel = ll
}

func (l XORMLogger) ShowSQL(show ...bool) {
}

func (l XORMLogger) IsShowSQL() bool {
	return true
}
