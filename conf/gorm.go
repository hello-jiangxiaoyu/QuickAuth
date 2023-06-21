package conf

import (
	"fmt"
	"gorm.io/gorm"
)

type Gorm struct {
	DbType   string `json:"db-type"`
	DbName   string `json:"db-name"`
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Username string `json:"username"`
	Password string `json:"password"`
	LogLevel string `json:"log-level"`
	Timeout  string `json:"timeout"`
	Config   string `json:"config"`
}

func (g *Gorm) GetDsn() string {
	res := ""
	switch g.DbType {
	case DBPostgres, DBTimeScale:
		res = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d %s",
			g.Host, g.Username, g.Password, g.DbName, g.Port, g.Config)
	case DBMySQL:
		res = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local&timeout=%s",
			g.Username, g.Password, g.Host, g.Port, g.DbName, g.Timeout)
	case DBSqlite:
		res = g.DbName
	}
	return res
}

func (g *Gorm) GetGormConfig() *gorm.Config {
	return &gorm.Config{}
}
