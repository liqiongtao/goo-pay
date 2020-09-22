package goo__

import (
	"github.com/go-xorm/xorm"
	gooDB "googo.io/goo/db"
)

func InitDB(cf gooDB.Config) {
	gooDB.Init(cf)
}

func DB() *xorm.EngineGroup {
	return gooDB.Orm()
}
