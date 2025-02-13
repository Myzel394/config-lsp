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
}

var ServerOptions = new(ServerOptionsType)

func InitServerOptions() {
	if slices.Contains(os.Args, "--no-undetectable-errors") {
		Log.Info("config-lsp will not return errors for undetectable files")
		ServerOptions.NoUndetectableErrors = true
	}
}
