package ast

import (
	"config-lsp/common"
	commonparser "config-lsp/common/parser"
	"config-lsp/handlers/ssh_config/ast/parser"
	hostparser "config-lsp/handlers/ssh_config/host-parser"
	"config-lsp/handlers/ssh_config/match-parser"
	"strings"

	"github.com/emirpasic/gods/maps/treemap"
	gods "github.com/emirpasic/gods/utils"
)

type sshListenerContext struct {
	line                uint32
	currentOption       *SSHOption
	currentBlock        SSHBlock
	currentKeyIsBlockOf *SSHBlockType
}

func createListenerContext() *sshListenerContext {
	context := new(sshListenerContext)

	return context
}

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

type sshParserListener struct {
	*parser.BaseConfigListener
	Config     *SSHConfig
	Errors     []common.LSPError
	sshContext *sshListenerContext
}

func (s *sshParserListener) EnterEntry(ctx *parser.EntryContext) {
	location := common.CharacterRangeFromCtx(ctx.BaseParserRuleContext)
	location.ChangeBothLines(s.sshContext.line)

	value := commonparser.ParseRawString(ctx.GetText(), commonparser.FullFeatures)
	option := &SSHOption{
		LocationRange: location,
		Value:         value,
	}

	s.sshContext.currentOption = option
}

func (s *sshParserListener) EnterKey(ctx *parser.KeyContext) {
	location := common.CharacterRangeFromCtx(ctx.BaseParserRuleContext)
	location.ChangeBothLines(s.sshContext.line)

	text := ctx.GetText()
	value := commonparser.ParseRawString(text, commonparser.FullFeatures)
	key := strings.Trim(
		value.Value,
		" ",
	)

	switch strings.ToLower(text) {
	case "match":
		value := SSHBlockTypeMatch
		s.sshContext.currentKeyIsBlockOf = &value
	case "host":
		value := SSHBlockTypeHost
		s.sshContext.currentKeyIsBlockOf = &value
	}

	s.sshContext.currentOption.Key = &SSHKey{
		LocationRange: location,
		Value:         value,
		Key:           key,
	}
}

func (s *sshParserListener) EnterSeparator(ctx *parser.SeparatorContext) {
	location := common.CharacterRangeFromCtx(ctx.BaseParserRuleContext)
	location.ChangeBothLines(s.sshContext.line)

	text := ctx.GetText()
	value := commonparser.ParseRawString(text, commonparser.FullFeatures)

	s.sshContext.currentOption.Separator = &SSHSeparator{
		LocationRange: location,
		Value:         value,
	}
}

func (s *sshParserListener) EnterValue(ctx *parser.ValueContext) {
	location := common.CharacterRangeFromCtx(ctx.BaseParserRuleContext)
	location.ChangeBothLines(s.sshContext.line)

	value := commonparser.ParseRawString(ctx.GetText(), commonparser.FullFeatures)
	s.sshContext.currentOption.OptionValue = &SSHValue{
		LocationRange: location,
		Value:         value,
	}
}

func (s *sshParserListener) ExitEntry(ctx *parser.EntryContext) {
	location := common.CharacterRangeFromCtx(ctx.BaseParserRuleContext)
	location.ChangeBothLines(s.sshContext.line)

	defer (func() {
		s.sshContext.currentOption = nil
	})()

	if s.sshContext.currentKeyIsBlockOf != nil {
		switch *s.sshContext.currentKeyIsBlockOf {
		case SSHBlockTypeMatch:
			var match *matchparser.Match

			matchParser := matchparser.NewMatch()
			errors := matchParser.Parse(
				s.sshContext.currentOption.OptionValue.Value.Raw,
				location.Start.Line,
				s.sshContext.currentOption.OptionValue.Start.Character,
			)

			if len(errors) > 0 {
				for _, err := range errors {
					s.Errors = append(s.Errors, common.LSPError{
						Range: err.Range,
						Err:   err.Err,
					})
				}
			} else {
				match = matchParser
			}

			matchBlock := &SSHMatchBlock{
				LocationRange: location,
				MatchOption:   s.sshContext.currentOption,
				MatchValue:    match,
				Options:       treemap.NewWith(gods.UInt32Comparator),
			}

			s.Config.Options.Put(
				location.Start.Line,
				matchBlock,
			)

			s.sshContext.currentKeyIsBlockOf = nil
			s.sshContext.currentBlock = matchBlock
		case SSHBlockTypeHost:
			var host *hostparser.Host

			hostParser := hostparser.NewHost()
			errors := hostParser.Parse(
				s.sshContext.currentOption.OptionValue.Value.Raw,
				location.Start.Line,
				s.sshContext.currentOption.OptionValue.Start.Character,
			)

			if len(errors) > 0 {
				for _, err := range errors {
					s.Errors = append(s.Errors, common.LSPError{
						Range: err.Range,
						Err:   err.Err,
					})
				}
			} else {
				host = hostParser
			}

			hostBlock := &SSHHostBlock{
				LocationRange: location,
				HostOption: s.sshContext.currentOption,
				HostValue: host,
				Options: treemap.NewWith(gods.UInt32Comparator),
			}

			s.Config.Options.Put(
				location.Start.Line,
				hostBlock,
			)

			s.sshContext.currentKeyIsBlockOf = nil
			s.sshContext.currentBlock = hostBlock
		}

		return
	}

	if s.sshContext.currentBlock == nil {
		s.Config.Options.Put(
			location.Start.Line,
			s.sshContext.currentOption,
		)
	} else {
		block := s.sshContext.currentBlock
		block.AddOption(s.sshContext.currentOption)
		block.SetEnd(location.End)
	}
}
