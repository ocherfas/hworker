package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"

	"gopkg.in/yaml.v3"
)

type Config struct {
	maxActivityTime time.Duration
	inactivityTime  time.Duration
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

	data := make(map[string]string)

	err = yaml.Unmarshal(yfile, &data)
	if err != nil {
		return err
	}

	maxActivityTimeString := data["max-activity-time"]
	inactivityTimeString := data["inactivity-timeout"]

	if maxActivityTimeString == "" || inactivityTimeString == "" {
		return fmt.Errorf("missing configuration")
	}

	maxTimeActivity, errParse1 := time.ParseDuration(maxActivityTimeString)
	if errParse1 != nil {
		err := fmt.Errorf("Incorrect format for max time activity")
		panic(err)
	}
	inactivityTime, errParse2 := time.ParseDuration(inactivityTimeString)
	if errParse2 != nil {
		err := fmt.Errorf("Incorrect format for inactivity time")
		panic(err)
	}

	config.inactivityTime = inactivityTime
	config.maxActivityTime = maxTimeActivity

	return nil
}
