package config

import (
	"gopkg.in/yaml.v3"
	"log"
	"os"
)

type AppEnv struct {
	Name string `yaml:"name"`
}

type Config struct {
	AppEnv AppEnv `yaml:"app-env"`
}

func GetConfig(configFile string) *Config {
	content, err := os.ReadFile(configFile)
	if err != nil {
		log.Fatalf("failed to read config file: %v", err)
	}

	var config Config
	if err := yaml.Unmarshal(content, &config); err != nil {
		log.Fatalf("failed to unmarshal config: %v", err)
	}
	return &config
}
