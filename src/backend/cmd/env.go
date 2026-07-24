package cmd

import (
	"os"
	"strings"
)

func GetDebug() bool {
	debug := os.Getenv("DEBUG")
	debug = strings.ToLower(debug)

	if debug != "true" {
		return false
	}

	return true
}
