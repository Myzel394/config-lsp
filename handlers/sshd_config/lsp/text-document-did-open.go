package lsp

import (
	"config-lsp/common"
	"config-lsp/handlers/sshd_config"
	"config-lsp/handlers/sshd_config/analyzer"
	"config-lsp/handlers/sshd_config/ast"
	"config-lsp/utils"

	"github.com/tliron/glsp"
	protocol "github.com/tliron/glsp/protocol_3_16"
)

func TextDocumentDidOpen(
	context *glsp.Context,
	params *protocol.DidOpenTextDocumentParams,
) error {
	common.ClearDiagnostics(context, params.TextDocument.URI)

	parser := ast.NewSSHConfig()
	document := sshdconfig.SSHDocument{
		Config: parser,
	}
	sshdconfig.DocumentParserMap[params.TextDocument.URI] = &document

	errors := parser.Parse(params.TextDocument.Text)

	diagnostics := utils.Map(
		errors,
		func(err common.LSPError) protocol.Diagnostic {
			return err.ToDiagnostic()
		},
	)

	diagnostics = append(
		diagnostics,
		analyzer.Analyze(&document)...,
	)

	if len(diagnostics) > 0 {
		common.SendDiagnostics(context, params.TextDocument.URI, diagnostics)
	}

	return nil
}