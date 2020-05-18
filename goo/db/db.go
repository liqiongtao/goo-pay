package gooDB

import (
	"github.com/go-xorm/xorm"
)

var (
	__db *db
)

type db struct {
	*xorm.EngineGroup
}

func Init(cf Config) {
	__db = &db{
		EngineGroup: NewXORM(cf),
	}
}

func Orm() *xorm.EngineGroup {
	return __db.EngineGroup
}
