package main

import (
	roothandler "config-lsp/root-handler"
	"fmt"
	"os"

	"github.com/tliron/commonlog"

	// Must include a backend implementation
	// See CommonLog for other options: https://github.com/tliron/commonlog
	_ "github.com/tliron/commonlog/simple"
)

func main() {
	if len(os.Args) > 1 && (os.Args[1] == "--version" || os.Args[1] == "version") {
		fmt.Println(roothandler.Version)

		os.Exit(0)
		return
	}

	// This increases logging verbosity (optional)
	commonlog.Configure(1, nil)

	roothandler.SetUpRootHandler()
}
