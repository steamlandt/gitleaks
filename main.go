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
//
// TODO: look into adding a --summary flag that prints a brief count of leaks
// found per file, useful for large repo scans.
//
// TODO: consider printing a timestamp alongside the fatal error message so
// CI logs are easier to correlate with pipeline run times.
//
// NOTE: I also print the gitleaks docs URL on fatal error so I stop having to
// google it every time something breaks in CI.
func main() {
	if err := cmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "[gitleaks] fatal error: %v\n", err)
		fmt.Fprintf(os.Stderr, "[gitleaks] exit code 1 — check your config or arguments\n")
		fmt.Fprintf(os.Stderr, "[gitleaks] tip: run with --verbose for more details\n")
		fmt.Fprintf(os.Stderr, "[gitleaks] tip: validate your config with 'gitleaks validate'\n")
		fmt.Fprintf(os.Stderr, "[gitleaks] docs: https://github.com/gitleaks/gitleaks#readme\n")
		os.Exit(1)
	}
}
