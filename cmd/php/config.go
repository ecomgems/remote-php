package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

type AppConfig struct {
	Server    string `yaml:"server"`
	Container string `yaml:"container"`
}

func buildConfig() (*AppConfig, error) {
	configPath := ".remote-php.yml"

	configFilename, err := filepath.Abs(configPath)
	if err != nil {
		return nil, fmt.Errorf("failed to understand file path: %v", err)
	}

	_, err = os.Stat(configFilename)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, fmt.Errorf("config file not found: %s", configFilename)
		}

		return nil, fmt.Errorf("Failed to check if config file exists: %v", err)
	}

	config, err := readConfig(configFilename)
	if err != nil {
		return nil, fmt.Errorf("Failed to read config: %v", err)
	}

	return config, nil
}

// readConfig reads a YAML file and returns a AppConfig struct.
func readConfig(filename string) (*AppConfig, error) {
	config := &AppConfig{}

	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	err = yaml.Unmarshal(bytes, config)
	if err != nil {
		return nil, err
	}

	return config, nil
}
