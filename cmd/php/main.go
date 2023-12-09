package main

import (
	"fmt"
	"golang.org/x/crypto/ssh"
	"os"
)

func main() {
	appConfig, err := buildAppConfig()
	if err != nil {
		fmt.Printf("encountered an error: %v", err)
		os.Exit(1)
	}

	sshConfig, err := NewSshConfig(appConfig)
	if err != nil {
		fmt.Printf("encountered an error: %v", err)
		os.Exit(1)
	}

	sshServerStr := fmt.Sprintf(
		"%s:%d",
		appConfig.SshHost,
		appConfig.SshPort,
	)
	sshConn, err := ssh.Dial("tcp", sshServerStr, sshConfig)
	defer sshConn.Close()
	if err != nil {
		fmt.Printf("encountered an error: %v", err)
		os.Exit(1)
	}

	session, err := sshConn.NewSession()
	if err != nil {
		fmt.Printf("encountered an error: %v", err)
		os.Exit(1)
	}
	defer session.Close()

	session.Stdout = os.Stdout
	runCommand := buildPhpCmd(appConfig)

	err = session.Run(runCommand)
	if err != nil {
		fmt.Printf("encountered an error: %v", err)
		os.Exit(1)
	}
}
