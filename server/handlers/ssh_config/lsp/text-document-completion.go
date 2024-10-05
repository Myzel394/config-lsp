package lsp

import (
	"config-lsp/common"
	sshconfig "config-lsp/handlers/ssh_config"
	"config-lsp/handlers/ssh_config/handlers"
	"regexp"

	"github.com/tliron/glsp"
	protocol "github.com/tliron/glsp/protocol_3_16"
)

var isEmptyPattern = regexp.MustCompile(`^\s*$`)

func TextDocumentCompletion(context *glsp.Context, params *protocol.CompletionParams) (any, error) {
	line := params.Position.Line
	cursor := common.LSPCharacterAsCursorPosition(params.Position.Character)

	d := sshconfig.DocumentParserMap[params.TextDocument.URI]

	if _, found := d.Config.CommentLines[line]; found {
		return nil, nil
	}

	option, block := d.Config.FindOption(line)

	if option == nil ||
		option.Separator == nil ||
		option.Key == nil ||
		option.Key.IsPositionBeforeEnd(cursor) {

		return handlers.GetRootCompletions(
			d,
			block,
			// Empty line, or currently typing a new key
			option == nil || isEmptyPattern.Match([]byte(option.Value.Raw[cursor:])),
		)
	}

	if option.Separator != nil && option.OptionValue.IsPositionAfterStart(cursor) {
		return handlers.GetOptionCompletions(
			d,
			option,
			block,
			line,
			cursor,
		)
	}

	return nil, nil
}
