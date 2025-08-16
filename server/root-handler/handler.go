package roothandler

import (
	"config-lsp/common"
	"config-lsp/root-handler/lsp"

	"github.com/tliron/glsp"
	protocol "github.com/tliron/glsp/protocol_3_16"

	aliases_handlers "config-lsp/handlers/aliases/handlers"
	bitcoin_conf_handlers "config-lsp/handlers/bitcoin_conf/handlers"
	hosts_handlers "config-lsp/handlers/hosts/handlers"
	ssh_handlers "config-lsp/handlers/ssh_config/handlers"
	wireguard_handlers "config-lsp/handlers/wireguard/handlers"

	"github.com/tliron/glsp/server"
)

const lsName = "config-lsp"

var lspHandler protocol.Handler

// The root handler which handles all the LSP requests and then forwards them to the appropriate handler
func SetUpRootHandler() {
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

	common.Log.Info("config-lsp server started")

	server.RunStdio()
}

func initialize(context *glsp.Context, params *protocol.InitializeParams) (any, error) {
	capabilities := lspHandler.CreateServerCapabilities()
	capabilities.TextDocumentSync = protocol.TextDocumentSyncKindFull
	capabilities.SignatureHelpProvider = &protocol.SignatureHelpOptions{}
	capabilities.ExecuteCommandProvider = &protocol.ExecuteCommandOptions{
		Commands: []string{
			"aliases." + string(aliases_handlers.CodeActionSendTestMail),

			"hosts." + string(hosts_handlers.CodeActionInlineAliases),

			"sshconfig." + string(ssh_handlers.CodeActionAddToUnknown),

			"wireguard." + string(wireguard_handlers.CodeActionGeneratePrivateKey),
			"wireguard." + string(wireguard_handlers.CodeActionGeneratePresharedKey),
			"wireguard." + string(wireguard_handlers.CodeActionGenerateDownRule),
			"wireguard." + string(wireguard_handlers.CodeActionCreatePeer),

			"bitcoinconf." + string(bitcoin_conf_handlers.CodeActionGenerateRPCAuth),
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
			Version: &common.Version,
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
