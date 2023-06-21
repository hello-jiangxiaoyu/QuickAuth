package conf

type Log struct {
	Type   string `json:"type"`
	Dir    string `json:"dir"`
	Level  string `json:"level"`
	MaxAge int    `json:"max-age"`
}
