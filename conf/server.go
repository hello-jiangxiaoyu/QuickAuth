package conf

type Server struct {
	Listen         string `json:"listen" yaml:"listen"`
	IsUpload       bool   `json:"uploadSwitch" yaml:"uploadSwitch"`
	IsWildcard     bool   `json:"wildcardSwitch" yaml:"wildcardSwitch"`
	WildcardSuffix string `json:"wildcardSuffix" yaml:"wildcardSuffix"`
}

type Default struct {
	DefaultUsername string `json:"defaultUsername" yaml:"defaultUsername"`
	DefaultPassword string `json:"defaultPassword" yaml:"defaultPassword"`
	DefaultTenant   string `json:"defaultTenant" yaml:"defaultTenant"`
	DefaultHost     string `json:"defaultHost" yaml:"defaultHost"`
}
