package conf

import (
	"encoding/json"
)

var (
	conf *Conf
)

type Conf struct {
	Mysql *DbConf
}

type DbConf struct {
	Host   string
	Port   uint16
	User   string
	Pass   string
	DbName string
}

func GetIns() *Conf {
	return conf
}

func Init(confJson string) (*Conf, error) {
	conf = &Conf{}
	err := json.Unmarshal([]byte(confJson), conf)
	return conf, err
}
