package ast

import (
	"config-lsp/common"
	"config-lsp/handlers/sshd_config/ast/parser"
	"github.com/emirpasic/gods/maps/treemap"
	gods "github.com/emirpasic/gods/utils"
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

	s.sshContext.currentOption = option
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

func (s *sshParserListener) EnterSeparator(ctx *parser.SeparatorContext) {
	location := common.CharacterRangeFromCtx(ctx.BaseParserRuleContext)
	location.ChangeBothLines(s.sshContext.line)

	s.sshContext.currentOption.Separator = &SSHSeparator{
		LocationRange: location,
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

func (s *sshParserListener) ExitEntry(ctx *parser.EntryContext) {
	location := common.CharacterRangeFromCtx(ctx.BaseParserRuleContext)
	location.ChangeBothLines(s.sshContext.line)

	if s.sshContext.isKeyAMatchBlock {
		// Add new match block
		matchBlock := &SSHMatchBlock{
			LocationRange: location,
			MatchEntry:    s.sshContext.currentOption,
			Options:       treemap.NewWith(gods.UInt32Comparator),
		}
		s.Config.Options.Put(
			location.Start.Line,
			matchBlock,
		)

		s.sshContext.currentMatchBlock = matchBlock
		s.sshContext.isKeyAMatchBlock = false
	} else if s.sshContext.currentMatchBlock != nil {
		s.sshContext.currentMatchBlock.Options.Put(
			location.Start.Line,
			s.sshContext.currentOption,
		)
	} else {
		s.Config.Options.Put(
			location.Start.Line,
			s.sshContext.currentOption,
		)
	}

	s.sshContext.currentOption = nil
}
