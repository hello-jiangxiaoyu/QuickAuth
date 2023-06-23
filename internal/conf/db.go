package conf

import (
	"fmt"
)

const (
	DBPostgres  = "postgres"
	DBTimeScale = "timescale"
	DBMySQL     = "mysql"
	DBSqlite    = "sqlite"
)

type Database struct {
	Host     string `json:"host" yaml:"host"`
	Port     uint   `json:"port" yaml:"port"`
	DbName   string `json:"dbName" yaml:"dbName"`
	DbType   string `json:"dbType" yaml:"dbType"`
	UserName string `json:"username" yaml:"username"`
	Password string `json:"password" yaml:"password"`
	Timeout  uint   `json:"timeout" yaml:"timeout"`
	Config   string `json:"config" yaml:"config"`
}

func (g *Database) GetDsn() string {
	res := ""
	switch g.DbType {
	case DBPostgres, DBTimeScale:
		res = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d %s",
			g.Host, g.UserName, g.Password, g.DbName, g.Port, g.Config)
	case DBMySQL:
		res = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local&timeout=%s",
			g.UserName, g.Password, g.Host, g.Port, g.DbName, g.Timeout)
	case DBSqlite:
		res = g.DbName
	}
	return res
}
