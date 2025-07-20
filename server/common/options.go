package common

import (
	"os"
	"slices"
)

type UsageReportType uint8

const (
	UsageReportNone UsageReportType = iota
	UsageReportErrorOnly
	UsageReportAll
)

// Global options for the server
type ServerOptionsType struct {
	// If true, the server will not return any errors if the
	// language was undetectable.
	// This is used for example in the VS Code extension, where
	// we show a native warning. The error message boxes just clutter
	// the interface.
	NoUndetectableErrors bool

	// If true, the server will not detect typos and suggest
	// the correct keywords.
	// Since the server finds typos using the Damerau-Levenshtein distance,
	// and this is done each time code actions are requested
	// (which happens quite often), these suggestions can eat a lot of resources.
	// You may want to enable this option if you are dealing with little
	// resources or if you're low on battery.
	NoTypoSuggestions bool

	// By default, usage reports are enabled.
	// Reports are sent to the self-hosted Sentry server.
	// 30 days retention, no personal data, data is not shared with third parties.
	// `--usage-reports-disable` disables all usage reports.
	// `--usage-reports-errors-only` only sends error reports.
	AllowedUsageReports UsageReportType

	IsDebug bool
}

var ServerOptions = new(ServerOptionsType)

func InitServerOptions() {
	ServerOptions.NoUndetectableErrors = false
	ServerOptions.NoTypoSuggestions = false
	ServerOptions.AllowedUsageReports = UsageReportAll
	ServerOptions.IsDebug = false

	if slices.Contains(os.Args, "--no-undetectable-errors") {
		ServerOptions.NoUndetectableErrors = true
	}

	if slices.Contains(os.Args, "--no-typo-suggestions") {
		ServerOptions.NoTypoSuggestions = true
	}

	if slices.Contains(os.Args, "--usage-reports-disable") {
		ServerOptions.AllowedUsageReports = UsageReportNone
	} else if slices.Contains(os.Args, "--usage-reports-errors-only") {
		ServerOptions.AllowedUsageReports = UsageReportErrorOnly
	}

	if slices.Contains(os.Args, "--env-debug") {
		println("config-lsp will run in debug mode")
		ServerOptions.IsDebug = true
	}
}
