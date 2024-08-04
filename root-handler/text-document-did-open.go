package roothandler

import (
	"config-lsp/common"
	fstab "config-lsp/handlers/fstab"
	"fmt"

	"github.com/tliron/glsp"
	protocol "github.com/tliron/glsp/protocol_3_16"
)

func TextDocumentDidOpen(context *glsp.Context, params *protocol.DidOpenTextDocumentParams) error {
	common.ClearDiagnostics(context, params.TextDocument.URI)

	// Find the file type
	content := params.TextDocument.Text
	language, err := DetectLanguage(content, params.TextDocument.LanguageID, params.TextDocument.URI)

	if err != nil {
		parseError := err.(common.ParseError)
		showParseError(
			context,
			params.TextDocument.URI,
			parseError,
		)

		return parseError.Err
	}

	// Everything okay, now we can handle the file
	rootHandler.AddDocument(params.TextDocument.URI, language)

	switch language {
	case LanguageFstab:
		return fstab.TextDocumentDidOpen(context, params)
	case LanguageSSHDConfig:
	default:
		panic(fmt.Sprintf("unexpected roothandler.SupportedLanguage: %#v", language))
	}

	return nil
}

func showParseError(
	context *glsp.Context,
	uri protocol.DocumentUri,
	err common.ParseError,
) {
	severity := protocol.DiagnosticSeverityError

	common.SendDiagnostics(
		context,
		uri,
		[]protocol.Diagnostic{
			{
				Severity: &severity,
				Message:  err.Err.Error(),
				Range: protocol.Range{
					Start: protocol.Position{
						Line:      err.Line,
						Character: 0,
					},
					End: protocol.Position{
						Line:      err.Line,
						Character: 99999,
					},
				},
			},
		},
	)
}