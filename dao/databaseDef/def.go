// Code generated with the command 'go generate', DO NOT EDIT.

package databaseDef

import (
	"demo/conf"
	"demo/dao/base"
)

func Db1() base.DatabaseInfo { return base.DatabaseInfo(conf.GetIns().Mysql.DbName) }
func Db2() base.DatabaseInfo { return base.DatabaseInfo(conf.GetIns().Mysql.DbName) }
