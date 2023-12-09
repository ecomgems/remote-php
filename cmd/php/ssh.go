package main

import (
	"fmt"
	"golang.org/x/crypto/ssh"
	"io/ioutil"
	"net"
	"os"
	"os/user"
	"strings"
	"time"
)

func NewSshConfig(serverConfig *AppConfig) (*ssh.ClientConfig, error) {
	config := &ssh.ClientConfig{
		User: serverConfig.SshUser,
		HostKeyCallback: func(hostname string, remote net.Addr, key ssh.PublicKey) error {
			return nil
		},
		Timeout: 15 * time.Second,
	}

	authMethod, err := getSSHAuthMethod(serverConfig)
	if err != nil {
		return nil, err
	}

	config.Auth = []ssh.AuthMethod{authMethod}

	return config, nil
}

func getSSHAuthMethod(serverConfig *AppConfig) (ssh.AuthMethod, error) {
	if serverConfig.SshPassword != "" {
		return ssh.Password(serverConfig.SshPassword), nil
	}

	var keyFile string
	if serverConfig.SshKeyFile == "" {
		usr, _ := user.Current()
		if usr != nil {
			keyFile = usr.HomeDir + "/.ssh/id_rsa"
		} else {
			keyFile = "/root/.ssh/id_rsa"
		}

	} else {
		keyFile = getFullKeyPath(serverConfig.SshKeyFile)
	}

	var key ssh.Signer
	var err error

	buf, err := ioutil.ReadFile(keyFile)
	if err != nil {
		return nil, fmt.Errorf("error reading SSH key file %s: %s", serverConfig.SshKeyFile, err.Error())
	}

	encrypted := serverConfig.SshKeyFilePassword != ""
	if encrypted {
		key, err = ssh.ParsePrivateKeyWithPassphrase(buf, []byte(serverConfig.SshKeyFilePassword))
		if err != nil {
			return nil, fmt.Errorf("error parsing encrypted key: %s", err.Error())
		}
	} else {
		key, err = ssh.ParsePrivateKey(buf)
		if err != nil {
			return nil, fmt.Errorf("error parsing key: %s", err.Error())
		}
	}

	return ssh.PublicKeys(key), nil
}

func getFullKeyPath(keyPath string) string {
	return strings.Replace(keyPath, "~", os.Getenv("HOME"), 1)
}
