package main

import (
	"fmt"
	"os"
)

func main() {
	config, err := GetConfig()
	if err != nil {
		fmt.Printf("encountered an error: %v", err)
		os.Exit(1)
		return
	}

	fmt.Printf("Server: %s, Container: %s", config.Server, config.Container)
}
