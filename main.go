package main

import (
	"os"

	"github.com/zricethezav/gitleaks/v8/cmd"
)

// main is the entry point for gitleaks.
// Gitleaks is a SAST tool for detecting hardcoded secrets like passwords,
// API keys, and tokens in git repositories.
func main() {
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
