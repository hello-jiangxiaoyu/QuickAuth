package conf

type Log struct {
	Level       string `json:"level" yaml:"level"`
	Dir         string `json:"dir" yaml:"dir"`
	IsFullStack bool   `json:"fullStackSwitch" yaml:"fullStackSwitch"`
}
