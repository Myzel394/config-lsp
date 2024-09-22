package ast

import (
	"config-lsp/common"
	commonparser "config-lsp/common/parser"
	"config-lsp/handlers/sshd_config/ast/parser"
	"config-lsp/handlers/sshd_config/match-parser"
	"strings"

	"github.com/emirpasic/gods/maps/treemap"
	gods "github.com/emirpasic/gods/utils"
)

type sshdListenerContext struct {
	line              uint32
	currentOption     *SSHDOption
	currentMatchBlock *SSHDMatchBlock
	isKeyAMatchBlock  bool
}

func createListenerContext() *sshdListenerContext {
	context := new(sshdListenerContext)
	context.isKeyAMatchBlock = false

	return context
}

func createListener(
	config *SSHDConfig,
	context *sshdListenerContext,
) sshdParserListener {
	return sshdParserListener{
		Config:      config,
		Errors:      make([]common.LSPError, 0),
		sshdContext: context,
	}
}

type sshdParserListener struct {
	*parser.BaseConfigListener
	Config      *SSHDConfig
	Errors      []common.LSPError
	sshdContext *sshdListenerContext
}

func (s *sshdParserListener) EnterEntry(ctx *parser.EntryContext) {
	location := common.CharacterRangeFromCtx(ctx.BaseParserRuleContext)
	location.ChangeBothLines(s.sshdContext.line)

	option := &SSHDOption{
		LocationRange: location,
		Value:         commonparser.ParseRawString(ctx.GetText(), commonparser.FullFeatures),
	}

	s.sshdContext.currentOption = option
}

func (s *sshdParserListener) EnterKey(ctx *parser.KeyContext) {
	location := common.CharacterRangeFromCtx(ctx.BaseParserRuleContext)
	location.ChangeBothLines(s.sshdContext.line)

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
		s.sshdContext.isKeyAMatchBlock = true
	}

	s.sshdContext.currentOption.Key = &SSHDKey{
		LocationRange: location,
		Value:         value,
		Key:           key,
	}
}

func (s *sshdParserListener) EnterSeparator(ctx *parser.SeparatorContext) {
	location := common.CharacterRangeFromCtx(ctx.BaseParserRuleContext)
	location.ChangeBothLines(s.sshdContext.line)

	text := ctx.GetText()
	value := commonparser.ParseRawString(text, commonparser.FullFeatures)

	s.sshdContext.currentOption.Separator = &SSHDSeparator{
		LocationRange: location,
		Value:         value,
	}
}

func (s *sshdParserListener) EnterValue(ctx *parser.ValueContext) {
	location := common.CharacterRangeFromCtx(ctx.BaseParserRuleContext)
	location.ChangeBothLines(s.sshdContext.line)

	s.sshdContext.currentOption.OptionValue = &SSHDValue{
		LocationRange: location,
		Value:         commonparser.ParseRawString(ctx.GetText(), commonparser.FullFeatures),
	}
}

func (s *sshdParserListener) ExitEntry(ctx *parser.EntryContext) {
	location := common.CharacterRangeFromCtx(ctx.BaseParserRuleContext)
	location.ChangeBothLines(s.sshdContext.line)

	defer (func() {
		s.sshdContext.currentOption = nil
	})()

	if s.sshdContext.isKeyAMatchBlock {
		// Add new match block
		var match *matchparser.Match

		if s.sshdContext.currentOption.OptionValue != nil {
			matchParser := matchparser.NewMatch()
			errors := matchParser.Parse(
				s.sshdContext.currentOption.OptionValue.Value.Raw,
				location.Start.Line,
				s.sshdContext.currentOption.OptionValue.Start.Character,
			)

			if len(errors) > 0 {
				for _, err := range errors {
					s.Errors = append(s.Errors, common.LSPError{
						Range: err.Range.ShiftHorizontal(s.sshdContext.currentOption.Start.Character),
						Err:   err.Err,
					})
				}
			} else {
				match = matchParser
			}
		}

		matchBlock := &SSHDMatchBlock{
			LocationRange: location,
			MatchOption:   s.sshdContext.currentOption,
			MatchValue:    match,
			Options:       treemap.NewWith(gods.UInt32Comparator),
		}
		s.Config.Options.Put(
			location.Start.Line,
			matchBlock,
		)

		s.sshdContext.currentMatchBlock = matchBlock

		s.sshdContext.isKeyAMatchBlock = false

		return
	}

	if s.sshdContext.currentMatchBlock != nil {
		s.sshdContext.currentMatchBlock.Options.Put(
			location.Start.Line,
			s.sshdContext.currentOption,
		)
		s.sshdContext.currentMatchBlock.End = s.sshdContext.currentOption.End
	} else {
		s.Config.Options.Put(
			location.Start.Line,
			s.sshdContext.currentOption,
		)
	}
}
