package main

import (
	"fmt"
	"os"
)

func main() {
	config, err := buildConfig()
	if err != nil {
		fmt.Printf("encountered an error: %v", err)
		os.Exit(1)
		return
	}

	fmt.Printf("server: %s, container: %s\n", config.Server, config.Container)
	fmt.Printf("%v\n", buildPHPCommand())
}
