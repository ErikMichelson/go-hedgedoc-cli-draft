package config

import (
	"log"
	"os"
	"path"
)

func GetConfigPath() string {
	var envVars = []string{"XDG_CONFIG_HOME", "APPDATA", "HOME"}
	var configPath string
	for i := 0; i < len(envVars); i++ {
		configPath = os.Getenv(envVars[i])
		if len(configPath) > 0 {
			if envVars[i] == "HOME" {
				configPath = path.Join(configPath, ".config")
			}
			if _, err := os.Stat(configPath); err == nil {
				break
			}
		}
	}
	configPath = path.Join(configPath, "hedgedoc-cli")
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		errDirCreation := os.MkdirAll(configPath, 0755)
		if errDirCreation != nil {
			log.Fatal(errDirCreation)
		}
	}
	return configPath
}
