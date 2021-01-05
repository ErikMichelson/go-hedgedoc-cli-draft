package config

import (
	"io/ioutil"
	"log"
	"os"
	"path"

	"gopkg.in/yaml.v2"
)

type Profile struct {
	Id     int
	Server string
	Token  string
}

type Config []Profile

var configDir = GetConfigPath()
var configFile = path.Join(configDir, "profiles.yml")

var loadedConfig Config

func Exists() bool {
	stat, err := os.Stat(configFile)
	return err == nil && !stat.IsDir()
}

func LoadConfig() {
	data, err := ioutil.ReadFile(configFile)
	if err != nil {
		log.Fatal(err)
	}
	err = yaml.Unmarshal(data, &loadedConfig)
	if err != nil {
		log.Fatal(err)
	}
}

func GetCurrentProfile() Profile {
	// TODO Use profile that was specified via flag or default profile as fallback
	return GetProfile(0)
}

func GetProfile(id int) Profile {
	for _, profile := range loadedConfig {
		if profile.Id == id {
			return profile
		}
	}
	return Profile{}
}

func SaveProfile(profile Profile) {
	config := Config{profile}
	data, err := yaml.Marshal(config)
	if err != nil {
		log.Fatal(err)
	}
	err = ioutil.WriteFile(configFile, data, 0600)
	if err != nil {
		log.Fatal(err)
	}
}
