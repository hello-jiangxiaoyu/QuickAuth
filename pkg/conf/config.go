package conf

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
)

type (
	Log struct {
		Level       string `json:"level" yaml:"level"`
		Dir         string `json:"dir" yaml:"dir"`
		IsFullStack bool   `json:"fullStackSwitch" yaml:"fullStackSwitch"`
	}
	Server struct {
		Listen         string `json:"listen" yaml:"listen"`
		IsUpload       bool   `json:"uploadSwitch" yaml:"uploadSwitch"`
		IsWildcard     bool   `json:"wildcardSwitch" yaml:"wildcardSwitch"`
		WildcardSuffix string `json:"wildcardSuffix" yaml:"wildcardSuffix"`
	}
	Default struct {
		DefaultUsername string `json:"defaultUsername" yaml:"defaultUsername"`
		DefaultPassword string `json:"defaultPassword" yaml:"defaultPassword"`
		DefaultTenant   string `json:"defaultTenant" yaml:"defaultTenant"`
		DefaultHost     string `json:"defaultHost" yaml:"defaultHost"`
	}

	SystemConfig struct {
		Svc        Server   `json:"server" yaml:"server"`
		Def        Default  `json:"default" yaml:"default"`
		Log        Log      `json:"log" yaml:"log"`
		Clickhouse Database `json:"clickhouse" yaml:"clickhouse"`
		DB         Database `json:"database" yaml:"database"`
		Redis      Database `json:"redis" yaml:"redis"`
	}
)

func NewSystemConfig(path string) (*SystemConfig, error) {
	if path == "" {
		path = DefaultSystemConfigName
	}
	yamlFile, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var res SystemConfig
	if err = yaml.Unmarshal(yamlFile, &res); err != nil {
		return nil, err
	}

	fmt.Println("new system configuration ok: ", path)
	return &res, nil
}
