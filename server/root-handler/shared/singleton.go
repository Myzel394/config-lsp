package shared

import (
	"config-lsp/root-handler/utils"

	protocol "github.com/tliron/glsp/protocol_3_16"
)

var Handler RootHandler

var OpenedFiles = make(map[protocol.DocumentUri]struct{})

type RootHandler struct {
	languageMap map[protocol.DocumentUri]utils.SupportedLanguage
}

func NewRootHandler() RootHandler {
	return RootHandler{
		languageMap: make(map[protocol.DocumentUri]utils.SupportedLanguage),
	}
}

func (h *RootHandler) AddDocument(uri protocol.DocumentUri, language utils.SupportedLanguage) {
	h.languageMap[uri] = language
}

func (h *RootHandler) GetLanguageForDocument(uri protocol.DocumentUri) *utils.SupportedLanguage {
	language, found := h.languageMap[uri]

	if !found {
		return nil
	}

	return &language
}

func (h *RootHandler) RemoveDocument(uri protocol.DocumentUri) {
	delete(h.languageMap, uri)
}
