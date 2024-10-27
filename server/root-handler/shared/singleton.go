package shared

import (
	protocol "github.com/tliron/glsp/protocol_3_16"
)

var Handler RootHandler

var OpenedFiles = make(map[protocol.DocumentUri]struct{})

type RootHandler struct {
	languageMap map[protocol.DocumentUri]SupportedLanguage
}

func NewRootHandler() RootHandler {
	return RootHandler{
		languageMap: make(map[protocol.DocumentUri]SupportedLanguage),
	}
}

func (h *RootHandler) AddDocument(uri protocol.DocumentUri, language SupportedLanguage) {
	h.languageMap[uri] = language
}

func (h *RootHandler) GetLanguageForDocument(uri protocol.DocumentUri) *SupportedLanguage {
	language, found := h.languageMap[uri]

	if !found {
		return nil
	}

	return &language
}

func (h *RootHandler) RemoveDocument(uri protocol.DocumentUri) {
	delete(h.languageMap, uri)
}
