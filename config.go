package main

import (
	"io/ioutil"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

type Config struct {
	ActivityMonitorConfigs []ActivityMonitorConfig `yaml:"activity-monitors"`
}

type ActivityMonitorConfig struct {
	MaxActivityTime string `yaml:"max-activity-time"`
	InactivityTime  string `yaml:"inactivity-time"`
	MessageFormat   string `yaml:"message-format"`
}

func NewConfig() Config {
	return Config{}
}

func (config *Config) readConfig() error {
	ex, err := os.Executable()
	if err != nil {
		return err
	}
	exPath := filepath.Dir(ex)

	configFilePath := filepath.Join(exPath, "..", "config.yaml")
	yfile, err := ioutil.ReadFile(configFilePath)
	if err != nil {
		return err
	}

	err = yaml.Unmarshal(yfile, config)

	return err
}
