package lsp

import (
	"config-lsp/common"
	"config-lsp/handlers/ssh_config"
	"config-lsp/handlers/ssh_config/analyzer"
	"config-lsp/handlers/ssh_config/ast"
	"config-lsp/utils"

	"github.com/tliron/glsp"
	protocol "github.com/tliron/glsp/protocol_3_16"
)

func TextDocumentDidOpen(
	context *glsp.Context,
	params *protocol.DidOpenTextDocumentParams,
) error {
	common.ClearDiagnostics(context, params.TextDocument.URI)

	var document *sshconfig.SSHDocument

	if foundDocument, ok := sshconfig.DocumentParserMap[params.TextDocument.URI]; ok {
		document = foundDocument
	} else {
		config := ast.NewSSHConfig()
		document = &sshconfig.SSHDocument{
			Config: config,
		}
		sshconfig.DocumentParserMap[params.TextDocument.URI] = document
	}

	errors := document.Config.Parse(params.TextDocument.Text)

	diagnostics := utils.Map(
		errors,
		func(err common.LSPError) protocol.Diagnostic {
			return err.ToDiagnostic()
		},
	)

	diagnostics = append(
		diagnostics,
		analyzer.Analyze(document)...,
	)

	if len(diagnostics) > 0 {
		common.SendDiagnostics(context, params.TextDocument.URI, diagnostics)
	}

	return nil
}
