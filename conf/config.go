package conf

import (
	"gopkg.in/yaml.v3"
	"os"
)

const (
	DBPostgres  = "postgres"
	DBTimeScale = "timescale"
	DBMySQL     = "mysql"
	DBSqlite    = "sqlite"
)

type SysConfig struct {
	Gorm Gorm `json:"gorm"`
	Log  Log  `json:"log"`
}

func NewSystemConfig() (*SysConfig, error) {
	yamlFile, err := os.ReadFile(GetSystemConfigFileName())
	if err != nil {
		return nil, err
	}

	var res SysConfig
	if err = yaml.Unmarshal(yamlFile, &res); err != nil {
		return nil, err
	}

	return &res, nil
}
