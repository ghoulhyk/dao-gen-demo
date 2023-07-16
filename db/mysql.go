package db

import (
	"demo/conf"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"os"
	"xorm.io/xorm"
	"xorm.io/xorm/log"
)

var orm *xorm.Engine

func OpenXorm() error {
	dbConf := conf.GetIns().Mysql
	connConfig := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbConf.User, dbConf.Pass, dbConf.Host, dbConf.Port, dbConf.DbName)

	engine, err := xorm.NewEngine("mysql", connConfig)
	if err != nil {
		return err
	}
	engine.ShowSQL(true)
	engine.SetLogLevel(log.LOG_DEBUG)
	engine.SetLogger(log.NewSimpleLogger(os.Stdout))
	err = engine.Ping()
	if err != nil {
		return err
	}
	orm = engine
	return nil
}

func NewSession() *xorm.Session {
	return orm.NewSession()
}
