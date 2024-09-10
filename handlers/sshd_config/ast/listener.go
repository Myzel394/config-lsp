package ast

import (
	"config-lsp/common"
	"config-lsp/handlers/sshd_config/ast/parser"
	"strings"
)

func createListener(
	config *SSHConfig,
	context *sshListenerContext,
) sshParserListener {
	return sshParserListener{
		Config:     config,
		Errors:     make([]common.LSPError, 0),
		sshContext: context,
	}
}

type sshListenerContext struct {
	line              uint32
	currentOption     *SSHOption
	currentMatchBlock *SSHMatchBlock
	isKeyAMatchBlock  bool
}

func createSSHListenerContext() *sshListenerContext {
	context := new(sshListenerContext)
	context.isKeyAMatchBlock = false

	return context
}

type sshParserListener struct {
	*parser.BaseConfigListener
	Config     *SSHConfig
	Errors     []common.LSPError
	sshContext *sshListenerContext
}

func (s *sshParserListener) EnterEntry(ctx *parser.EntryContext) {
	location := common.CharacterRangeFromCtx(ctx.BaseParserRuleContext)
	location.ChangeBothLines(s.sshContext.line)

	option := &SSHOption{
		LocationRange: location,
		Value:         ctx.GetText(),
	}

	if s.sshContext.currentMatchBlock == nil {
		s.Config.Options.Put(
			location.Start.Line,
			option,
		)

		s.sshContext.currentOption = option
	} else {
		s.sshContext.currentMatchBlock.Options.Put(
			location.Start.Line,
			option,
		)

		s.sshContext.currentOption = option
	}
}

func (s *sshParserListener) EnterKey(ctx *parser.KeyContext) {
	location := common.CharacterRangeFromCtx(ctx.BaseParserRuleContext)
	location.ChangeBothLines(s.sshContext.line)

	text := ctx.GetText()

	if strings.ToLower(text) == "match" {
		s.sshContext.isKeyAMatchBlock = true
	}

	s.sshContext.currentOption.Key = &SSHKey{
		LocationRange: location,
		Value:         ctx.GetText(),
	}
}

func (s *sshParserListener) EnterValue(ctx *parser.ValueContext) {
	location := common.CharacterRangeFromCtx(ctx.BaseParserRuleContext)
	location.ChangeBothLines(s.sshContext.line)

	s.sshContext.currentOption.OptionValue = &SSHValue{
		LocationRange: location,
		Value:         ctx.GetText(),
	}
}

func (s *sshParserListener) ExitValue(ctx *parser.ValueContext) {
	if s.sshContext.isKeyAMatchBlock {
		location := common.CharacterRangeFromCtx(ctx.BaseParserRuleContext)
		location.ChangeBothLines(s.sshContext.line)

		rawEntry, _ := s.Config.Options.Get(location.Start.Line)
		entry := rawEntry.(*SSHOption)

		// Overwrite the current match block
		matchBlock := &SSHMatchBlock{
			LocationRange: location,
			MatchEntry:    entry,
		}
		s.Config.Options.Put(
			location.Start.Line,
			matchBlock,
		)

		s.sshContext.currentMatchBlock = matchBlock
		s.sshContext.isKeyAMatchBlock = false
	}

	s.sshContext.currentOption = nil
}
