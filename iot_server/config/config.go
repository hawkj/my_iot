package config

import (
	"gopkg.in/yaml.v3"
	"log"
	"os"
)

type Config struct {
}

func GetConfig(configFile string) *Config {
	content, err := os.ReadFile(configFile)
	if err != nil {
		log.Fatalf("[os.ReadFile] failed to read config file: %v", err)
	}

	var config Config
	if err := yaml.Unmarshal(content, &config); err != nil {
		log.Fatalf("[yaml.Unmarshal] failed to unmarshal config: %v", err)
	}
	return &config
}
