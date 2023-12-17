package main

import (
	"fmt"
	"os"

	"golang.org/x/crypto/ssh"
	"golang.org/x/term"
)

func main() {
	appConfig, err := buildAppConfig()
	if err != nil {
		fmt.Printf("failed to build application configuration: %v\n", err)
		os.Exit(1)
	}

	sshConfig, err := NewSshConfig(appConfig)
	if err != nil {
		fmt.Printf("failed to create SSH configuration: %v\n", err)
		os.Exit(1)
	}

	sshServerStr := fmt.Sprintf("%s:%s", appConfig.SshHost, appConfig.SshPort)
	sshConn, err := ssh.Dial("tcp", sshServerStr, sshConfig)
	if err != nil {
		fmt.Printf("failed to establish SSH connection: %v\n", err)
		os.Exit(1)
	}
	defer sshConn.Close()

	session, err := sshConn.NewSession()
	if err != nil {
		fmt.Printf("failed to create SSH session: %v\n", err)
		os.Exit(1)
	}
	defer session.Close()

	session.Stdout = os.Stdout
	session.Stderr = os.Stderr
	session.Stdin = os.Stdin

	termWidth, termHeight, err := term.GetSize(int(os.Stdout.Fd()))
	if err != nil {
		termWidth = 80
		termHeight = 40
	}

	modes := ssh.TerminalModes{
		ssh.ECHO:          0,
		ssh.TTY_OP_ISPEED: 14400,
		ssh.TTY_OP_OSPEED: 14400,
	}

	err = session.RequestPty(os.Getenv("$TERM"), termHeight, termWidth, modes)
	if err != nil {
		fmt.Printf("failed to request PTY for SSH session: %v\n", err)
		os.Exit(1)
	}

	runCommand := buildPhpCmd(appConfig)
	err = session.Run(runCommand)
	if err != nil {
		fmt.Printf("failed to run the command over SSH: %v\n", err)
		os.Exit(1)
	}
}
