package roothandler

import (
	"config-lsp/root-handler/lsp"
	"config-lsp/root-handler/shared"
	"github.com/tliron/glsp"
	protocol "github.com/tliron/glsp/protocol_3_16"

	"github.com/tliron/glsp/server"
)

const lsName = "config-lsp"

// The comment below at the end of the line is required for the CI:CD to work.
// Do not remove it
var version = "0.1.1" // CI:CD-VERSION

var lspHandler protocol.Handler

// The root handler which handles all the LSP requests and then forwards them to the appropriate handler
func SetUpRootHandler() {
	shared.Handler = shared.NewRootHandler()

	lspHandler = protocol.Handler{
		Initialize:                  initialize,
		Initialized:                 initialized,
		Shutdown:                    shutdown,
		SetTrace:                    setTrace,
		TextDocumentDidOpen:         lsp.TextDocumentDidOpen,
		TextDocumentDidChange:       lsp.TextDocumentDidChange,
		TextDocumentCompletion:      lsp.TextDocumentCompletion,
		TextDocumentHover:           lsp.TextDocumentHover,
		TextDocumentDidClose:        lsp.TextDocumentDidClose,
		TextDocumentCodeAction:      lsp.TextDocumentCodeAction,
		TextDocumentDefinition:      lsp.TextDocumentDefinition,
		WorkspaceExecuteCommand:     lsp.WorkspaceExecuteCommand,
		TextDocumentRename:          lsp.TextDocumentRename,
		TextDocumentPrepareRename:   lsp.TextDocumentPrepareRename,
		TextDocumentSignatureHelp:   lsp.TextDocumentSignatureHelp,
		TextDocumentRangeFormatting: lsp.TextDocumentRangeFormattingFunc,
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
