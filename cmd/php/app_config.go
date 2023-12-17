package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

type AppConfig struct {
	SshHost             string `yaml:"ssh_host"`
	SshPort             string `yaml:"ssh_port"`
	SshUser             string `yaml:"ssh_user"`
	SshKeyFile          string `yaml:"ssh_key_file"`
	SshKeyFilePassword  string `yaml:"ssh_key_file_password"`
	SshPassword         string `yaml:"ssh_password"`
	DockerContainer     string `yaml:"docker_container"`
	DockerContainerUser string `yaml:"docker_container_user"`
	DockerContainerPath string `yaml:"docker_container_path"`
}

func buildAppConfig() (*AppConfig, error) {
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
