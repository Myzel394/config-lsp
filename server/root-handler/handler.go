package roothandler

import (
	"github.com/tliron/glsp"
	protocol "github.com/tliron/glsp/protocol_3_16"

	"github.com/tliron/glsp/server"
)

const lsName = "config-lsp"

var version string = "0.0.1"

var lspHandler protocol.Handler

// The root handler which handles all the LSP requests and then forwards them to the appropriate handler
func SetUpRootHandler() {
	rootHandler = NewRootHandler()
	lspHandler = protocol.Handler{
		Initialize:                  initialize,
		Initialized:                 initialized,
		Shutdown:                    shutdown,
		SetTrace:                    setTrace,
		TextDocumentDidOpen:         TextDocumentDidOpen,
		TextDocumentDidChange:       TextDocumentDidChange,
		TextDocumentCompletion:      TextDocumentCompletion,
		TextDocumentHover:           TextDocumentHover,
		TextDocumentDidClose:        TextDocumentDidClose,
		TextDocumentCodeAction:      TextDocumentCodeAction,
		TextDocumentDefinition:      TextDocumentDefinition,
		WorkspaceExecuteCommand:     WorkspaceExecuteCommand,
		TextDocumentRename:          TextDocumentRename,
		TextDocumentPrepareRename:   TextDocumentPrepareRename,
		TextDocumentSignatureHelp:   TextDocumentSignatureHelp,
		TextDocumentRangeFormatting: TextDocumentRangeFormattingFunc,
	}

	server := server.NewServer(&lspHandler, lsName, false)

	server.RunStdio()
}

func initialize(context *glsp.Context, params *protocol.InitializeParams) (any, error) {
	capabilities := lspHandler.CreateServerCapabilities()
	capabilities.TextDocumentSync = protocol.TextDocumentSyncKindFull
	capabilities.SignatureHelpProvider = &protocol.SignatureHelpOptions{}
	capabilities.ExecuteCommandProvider = &protocol.ExecuteCommandOptions{
		Commands: []string{
			"aliases.sendTestMail",

			"hosts.inlineAliases",

			"sshconfig.addToUnknown",

			"wireguard.generatePrivateKey",
			"wireguard.generatePresharedKey",
			"wireguard.addKeepalive",
		},
	}

	if (*params.Capabilities.TextDocument.Rename.PrepareSupport) == true {
		// Client supports rename preparation
		prepareRename := true
		capabilities.RenameProvider = protocol.RenameOptions{
			PrepareProvider: &prepareRename,
		}
	}

	return protocol.InitializeResult{
		Capabilities: capabilities,
		ServerInfo: &protocol.InitializeResultServerInfo{
			Name:    lsName,
			Version: &version,
		},
	}, nil
}

func initialized(context *glsp.Context, params *protocol.InitializedParams) error {
	return nil
}

func shutdown(context *glsp.Context) error {
	protocol.SetTraceValue(protocol.TraceValueOff)
	return nil
}

func setTrace(context *glsp.Context, params *protocol.SetTraceParams) error {
	protocol.SetTraceValue(params.Value)
	return nil
}
