package main

import (
	roothandler "config-lsp/root-handler"

	"github.com/tliron/commonlog"

	// Must include a backend implementation
	// See CommonLog for other options: https://github.com/tliron/commonlog
	_ "github.com/tliron/commonlog/simple"
)

func main() {
	// This increases logging verbosity (optional)
	commonlog.Configure(1, nil)

	roothandler.SetUpRootHandler()
}
