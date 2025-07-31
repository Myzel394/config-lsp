package shared

import (
	"config-lsp/common"

	protocol "github.com/tliron/glsp/protocol_3_16"
)

type OpenedFile struct {
	Language *SupportedLanguage

	// Stores information when the language could not be determined
	UnavailableInfo *UnavailableInfo
}

type UnavailableInfo struct {
	// Position of `#?lsp.language`
	OverwritePosition *common.LocationRange
}

var OpenedFiles = make(map[protocol.DocumentUri]OpenedFile)

func AddDocument(uri protocol.DocumentUri, file OpenedFile) {
	OpenedFiles[uri] = file
}

func RemoveDocument(uri protocol.DocumentUri) {
	delete(OpenedFiles, uri)
}

func GetDocument(uri protocol.DocumentUri) *OpenedFile {
	if file, exists := OpenedFiles[uri]; exists {
		return &file
	}
	return nil
}
