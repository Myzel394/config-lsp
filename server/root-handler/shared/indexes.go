package shared

import (
	protocol "github.com/tliron/glsp/protocol_3_16"
)

type LanguageOverwrite struct {
	Language SupportedLanguage

	// The start of the overwrite
	Raw       string
	Line      uint32
	Character uint32
}

var LanguagesOverwrites = map[protocol.DocumentUri]LanguageOverwrite{}
