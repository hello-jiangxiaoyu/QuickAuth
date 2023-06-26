package conf

import "os"

const (
	DefaultSystemConfigName = "deploy/dev.yaml"
	DefaultDeployType       = "local"
	DefaultIsFirstDeploy    = "false"
)

func getEnv(env, defaultValue string) string {
	value := os.Getenv(env)
	if value == "" {
		value = defaultValue
	}
	return value
}

func GetSystemConfigFileName() string {
	return getEnv("SYSTEM_CONFIG", DefaultSystemConfigName)
}

func GetDeployType() string {
	return getEnv("DEPLOY_TYPE", DefaultDeployType)
}

func IsFirstDeploy() bool {
	isFirst := getEnv("FIRST_DEPLOY", DefaultIsFirstDeploy)
	return isFirst == "true"
}
