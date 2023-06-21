package conf

import "os"

const (
	DefaultSystemConfigName = "sys.yaml"
)

func getEnv(env, defaultValue string) string {
	value := os.Getenv(env)
	if value == "" {
		value = defaultValue
	}
	return value
}

func GetSystemConfigFileName() string {
	return getEnv("SYS_CONFIG", DefaultSystemConfigName)
}
