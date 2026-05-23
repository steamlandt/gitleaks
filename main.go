package main

import (
	"fmt"
	"os"

	"github.com/zricethezav/gitleaks/v8/cmd"
)

// main is the entry point for gitleaks.
// Gitleaks is a SAST tool for detecting hardcoded secrets like passwords,
// API keys, and tokens in git repositories.
//
// Personal fork: added exit code output for easier debugging in CI pipelines.
// Note: exit code 1 = error running gitleaks, exit code 2 = leaks found (handled by cmd package).
//
// Exit code reference:
//   0 = no leaks found
//   1 = error running gitleaks (e.g. invalid config, bad args)
//   2 = leaks detected in scan
func main() {
	if err := cmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "[gitleaks] fatal error: %v\n", err)
		os.Exit(1)
	}
}
