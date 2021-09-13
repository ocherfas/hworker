package main

import (
	"io/ioutil"
	"log"
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
	configPath := os.Getenv("HWORKER_CONFIG")
	if configPath == "" {
		log.Println("HWORKER_CONFIG was not supplied. Reading from default path.")
		ex, err := os.Executable()
		if err != nil {
			return err
		}
		exPath := filepath.Dir(ex)

		configPath = filepath.Join(exPath, "..", "config.yaml")
	}

	yfile, err := ioutil.ReadFile(configPath)
	if err != nil {
		return err
	}

	err = yaml.Unmarshal(yfile, config)

	log.Printf("Successfully read config file from path %s\n", configPath)

	return err
}
