package main

import "os"

func buildPHPCommand() []string {
	additionalArgs := os.Args[1:]
	return append([]string{"php"}, additionalArgs...)
}
