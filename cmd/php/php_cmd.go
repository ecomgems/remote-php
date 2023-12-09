package main

import (
	"fmt"
	"os"
	"strings"
)

func buildPhpCmd(config *AppConfig) string {
	cmdPieces := []string{
		"docker",
		"exec",
	}

	if config.DockerContainerUser != "" {
		cmdPieces = append(cmdPieces, fmt.Sprintf("--user=%s", config.DockerContainerUser))
	}

	cmdPieces = append(cmdPieces, config.DockerContainer, "php")

	additionalArgs := os.Args[1:]
	cmdPieces = append(cmdPieces, additionalArgs...)
	cmdBody := strings.Join(cmdPieces, " ")

	return cmdBody
}
