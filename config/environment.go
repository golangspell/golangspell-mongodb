package config

import (
	"fmt"

	coreconfig "github.com/golangspell/golangspell/config"
)

const (
	//configFileName defines the configuration file name
	configFileName = ".golangspell.json"
)

//Values stores the current configuration values
var (
	Values coreconfig.Config
)

//ConfigFilePath contains the path of the config file
var ConfigFilePath = fmt.Sprintf("%s/%s", coreconfig.GetGolangspellHome(), configFileName)

// GetEnv gets an environment variable content or a default value
func GetEnv(key, defaultValue string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return defaultValue
}

func init() {
	Values.TestRun = GetEnv("TESTRUN", "false") == "true"
	Values.GoPath = GetEnv("GOPATH", fmt.Sprintf("%s%sgo", coreconfig.GetHomeDir(), coreconfig.PlatformSeparator))
	Values.LogLevel = GetEnv("LOG_LEVEL", "INFO")
}
