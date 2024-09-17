package ast

import (
	"config-lsp/common"
	commonparser "config-lsp/common/parser"
	"config-lsp/handlers/sshd_config/ast/parser"
	match_parser "config-lsp/handlers/sshd_config/fields/match-parser"
	"strings"

	"github.com/emirpasic/gods/maps/treemap"
	gods "github.com/emirpasic/gods/utils"
)

type sshListenerContext struct {
	line              uint32
	currentOption     *SSHDOption
	currentMatchBlock *SSHDMatchBlock
	isKeyAMatchBlock  bool
}

func createSSHListenerContext() *sshListenerContext {
	context := new(sshListenerContext)
	context.isKeyAMatchBlock = false

	return context
}

func createListener(
	config *SSHDConfig,
	context *sshListenerContext,
) sshParserListener {
	return sshParserListener{
		Config:     config,
		Errors:     make([]common.LSPError, 0),
		sshContext: context,
	}
}

type sshParserListener struct {
	*parser.BaseConfigListener
	Config     *SSHDConfig
	Errors     []common.LSPError
	sshContext *sshListenerContext
}

func (s *sshParserListener) EnterEntry(ctx *parser.EntryContext) {
	location := common.CharacterRangeFromCtx(ctx.BaseParserRuleContext)
	location.ChangeBothLines(s.sshContext.line)

	option := &SSHDOption{
		LocationRange: location,
		Value:         commonparser.ParseRawString(ctx.GetText(), commonparser.FullFeatures),
	}

	s.sshContext.currentOption = option
}

func (s *sshParserListener) EnterKey(ctx *parser.KeyContext) {
	location := common.CharacterRangeFromCtx(ctx.BaseParserRuleContext)
	location.ChangeBothLines(s.sshContext.line)

	text := ctx.GetText()
	value := commonparser.ParseRawString(text, commonparser.FullFeatures)
	key := strings.TrimRight(
		strings.TrimLeft(
			value.Value,
			" ",
		),
		" ",
	)

	if strings.ToLower(text) == "match" {
		s.sshContext.isKeyAMatchBlock = true
	}

	s.sshContext.currentOption.Key = &SSHDKey{
		LocationRange: location,
		Value:         value,
		Key:           key,
	}
}

func (s *sshParserListener) EnterSeparator(ctx *parser.SeparatorContext) {
	location := common.CharacterRangeFromCtx(ctx.BaseParserRuleContext)
	location.ChangeBothLines(s.sshContext.line)

	s.sshContext.currentOption.Separator = &SSHDSeparator{
		LocationRange: location,
	}
}

func (s *sshParserListener) EnterValue(ctx *parser.ValueContext) {
	location := common.CharacterRangeFromCtx(ctx.BaseParserRuleContext)
	location.ChangeBothLines(s.sshContext.line)

	s.sshContext.currentOption.OptionValue = &SSHDValue{
		LocationRange: location,
		Value:         commonparser.ParseRawString(ctx.GetText(), commonparser.FullFeatures),
	}
}

func (s *sshParserListener) ExitEntry(ctx *parser.EntryContext) {
	location := common.CharacterRangeFromCtx(ctx.BaseParserRuleContext)
	location.ChangeBothLines(s.sshContext.line)

	defer (func() {
		s.sshContext.currentOption = nil
	})()

	if s.sshContext.isKeyAMatchBlock {
		// Add new match block
		var match *match_parser.Match

		if s.sshContext.currentOption.OptionValue != nil {
			matchParser := match_parser.NewMatch()
			errors := matchParser.Parse(
				s.sshContext.currentOption.OptionValue.Value.Raw,
				location.Start.Line,
				s.sshContext.currentOption.OptionValue.Start.Character,
			)

			if len(errors) > 0 {
				for _, err := range errors {
					s.Errors = append(s.Errors, common.LSPError{
						Range: err.Range.ShiftHorizontal(s.sshContext.currentOption.Start.Character),
						Err:   err.Err,
					})
				}
			} else {
				match = matchParser
			}
		}

		matchBlock := &SSHDMatchBlock{
			LocationRange: location,
			MatchEntry:    s.sshContext.currentOption,
			MatchValue:    match,
			Options:       treemap.NewWith(gods.UInt32Comparator),
		}
		s.Config.Options.Put(
			location.Start.Line,
			matchBlock,
		)

		s.sshContext.currentMatchBlock = matchBlock

		s.sshContext.isKeyAMatchBlock = false

		return
	}

	if s.sshContext.currentMatchBlock != nil {
		s.sshContext.currentMatchBlock.Options.Put(
			location.Start.Line,
			s.sshContext.currentOption,
		)
		s.sshContext.currentMatchBlock.End = s.sshContext.currentOption.End
	} else {
		s.Config.Options.Put(
			location.Start.Line,
			s.sshContext.currentOption,
		)
	}
}
