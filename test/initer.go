package test

import (
	"demo/conf"
	"demo/db"
)

func initConf() {
	conf.Init(`
{
  "Mysql": {
    "Host": "127.0.0.1",
    "Port": 3306,
    "User": "root",
    "Pass": "",
    "DbName": "db_1_test"
  }
}
`)
}

func connectDb() {
	err := db.OpenXorm()
	if err != nil {
		panic(err)
	}
}
