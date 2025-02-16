package common

import (
	"os"
	"slices"
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
}

var ServerOptions = new(ServerOptionsType)

func InitServerOptions() {
	if slices.Contains(os.Args, "--no-undetectable-errors") {
		Log.Info("config-lsp will not return errors for undetectable files")
		ServerOptions.NoUndetectableErrors = true
	}

	if slices.Contains(os.Args, "--no-typo-suggestions") {
		Log.Info("config-lsp will not detect typos for keywords")
		ServerOptions.NoTypoSuggestions = true
	}
}
