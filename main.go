package main

import (
	"fmt"
	"os"

	"github.com/hiyongliz/upimage/cmd"
)

// Version information set by build flags
var (
	version   = "dev"
	buildTime = "unknown"
	commitSHA = "unknown"
)

func main() {
	if err := cmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
