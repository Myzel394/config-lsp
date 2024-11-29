package ast

import (
	"config-lsp/common"

	"github.com/antlr4-go/antlr/v4"
)

type errorListenerContext struct {
	line uint32
}

type errorListener struct {
	*antlr.DefaultErrorListener
	Errors  []common.LSPError
	context errorListenerContext
}

func (d *errorListener) SyntaxError(
	recognizer antlr.Recognizer,
	offendingSymbol interface{},
	_ int,
	character int,
	message string,
	error antlr.RecognitionException,
) {
	line := d.context.line
	d.Errors = append(d.Errors, common.LSPError{
		Range: common.CreateSingleCharRange(uint32(line), uint32(character)),
		Err: common.SyntaxError{
			Message: message,
		},
	})
}

func createErrorListener(
	line uint32,
) errorListener {
	return errorListener{
		Errors: make([]common.LSPError, 0),
		context: errorListenerContext{
			line: line,
		},
	}
}
