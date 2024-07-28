package handlers

import (
	"config-lsp/common"
	"fmt"
	"unicode/utf8"

	"github.com/tliron/glsp"
	protocol "github.com/tliron/glsp/protocol_3_16"
)

func ClearDiagnostics(context *glsp.Context, uri protocol.DocumentUri) {
	    context.Notify(
	"textDocument/publishDiagnostics",
	protocol.PublishDiagnosticsParams{
		URI: uri,
		Diagnostics: []protocol.Diagnostic{},
	},
    )
}

func SendDiagnosticsFromParserErrors(context *glsp.Context, uri protocol.DocumentUri, parserErrors []error) {
    diagnosticErrors := make([]protocol.Diagnostic, 0)

    for _, parserError := range parserErrors {
	switch parserError.(type) {
	case common.OptionAlreadyExistsError: {
		err := parserError.(common.OptionAlreadyExistsError)
		existingOption, _ := Parser.GetOption(err.Option)

		diagnosticErrors = append(diagnosticErrors, protocol.Diagnostic{
		    Range: protocol.Range{
			Start: protocol.Position{
			    Line: err.FoundOnLine,
			    Character: 0,
			},
			End: protocol.Position{
			    Line: err.FoundOnLine,
			    Character: uint32(utf8.RuneCountInString(err.Option)),
		        },
		    },
		    Message: fmt.Sprintf("Option '%s' has already been declared on line %v", err.Option, existingOption.Position.Line + 1),
		})
	}
	}
    }

    context.Notify(
	"textDocument/publishDiagnostics",
	protocol.PublishDiagnosticsParams{
		URI: uri,
		Diagnostics: diagnosticErrors,
	},
    )
}

