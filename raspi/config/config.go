package config

import (
	"gopkg.in/yaml.v3"
	"log"
	"os"
)

type AppEnv struct {
	Name string `yaml:"name"`
}

type SiteInfo struct {
	Name string `yaml:"name"`
}

type Emq struct {
	BrokerAddress string `yaml:"broker_address"`
	Username      string `yaml:"username"`
	Password      string `yaml:"password"`
}

type Config struct {
	AppEnv   AppEnv   `yaml:"app-env"`
	SiteInfo SiteInfo `yaml:"site-info"`
	Emq      Emq      `yaml:"emq"`
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
