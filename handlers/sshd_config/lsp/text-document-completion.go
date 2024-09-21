package lsp

import (
	"config-lsp/common"
	sshdconfig "config-lsp/handlers/sshd_config"
	"config-lsp/handlers/sshd_config/handlers"
	"regexp"

	"github.com/tliron/glsp"
	protocol "github.com/tliron/glsp/protocol_3_16"
)

var isEmptyPattern = regexp.MustCompile(`^\s*$`)

func TextDocumentCompletion(context *glsp.Context, params *protocol.CompletionParams) (any, error) {
	line := params.Position.Line
	cursor := common.LSPCharacterAsCursorPosition(params.Position.Character)

	d := sshdconfig.DocumentParserMap[params.TextDocument.URI]

	if _, found := d.Config.CommentLines[line]; found {
		return nil, nil
	}

	entry, matchBlock := d.Config.FindOption(line)

	if entry == nil ||
		entry.Separator == nil ||
		entry.Key == nil ||
		entry.Key.IsPositionBeforeEnd(cursor) {

		return handlers.GetRootCompletions(
			d,
			matchBlock,
			// Empty line, or currently typing a new key
			entry == nil || isEmptyPattern.Match([]byte(entry.Value.Value[cursor:])),
		)
	}

	if entry.Separator != nil && entry.OptionValue.IsPositionAfterStart(cursor) {
		return handlers.GetOptionCompletions(
			d,
			entry,
			matchBlock,
			cursor,
		)
	}

	return nil, nil
}
