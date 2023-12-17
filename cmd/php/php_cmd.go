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
		"-i",
	}

	if config.DockerContainerUser != "" {
		cmdPieces = append(cmdPieces, fmt.Sprintf("--user=%s", config.DockerContainerUser))
	}

	cmdPieces = append(cmdPieces, config.DockerContainer, "php")

	additionalArgs := os.Args[1:]

	workDir, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	for _, arg := range additionalArgs {
		cmdPieces = append(cmdPieces, strings.ReplaceAll(arg, workDir, config.DockerContainerPath))
	}

	cmdBody := strings.Join(cmdPieces, " ")

	return cmdBody
}
