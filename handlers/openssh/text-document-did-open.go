package handlers

import (
	"os"

	"github.com/tliron/glsp"
	protocol "github.com/tliron/glsp/protocol_3_16"
)

func TextDocumentDidOpen(context *glsp.Context, params *protocol.DidOpenTextDocumentParams) error {
	readBytes, err := os.ReadFile(params.TextDocument.URI[len("file://"):])

	if err != nil {
		return err
	}

	errors := Parser.ParseFromFile(string(readBytes))

	if len(errors) > 0 {
		return errors[0]
	}

	return nil
}


