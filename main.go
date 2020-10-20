package main

import (
    _ "fmt"
    "os"
    rootCmd "github.com/latelee/cmdtool/cmd"
)

func main() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
