package shared

import (
	protocol "github.com/tliron/glsp/protocol_3_16"
)

type OpenedFile struct {
	Language *SupportedLanguage
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
