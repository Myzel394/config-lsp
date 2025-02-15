package ast

import (
	"config-lsp/common"
	commonparser "config-lsp/common/parser"
	"config-lsp/handlers/gitconfig/ast/parser"
)

type gitconfigListenerContext struct {
	line                    uint32
	currentSection          *GitSection
	currentEntry            *GitEntry
	isWaitingForNextSection bool
	virtualLine             common.VirtualLine
}

func createListenerContext() *gitconfigListenerContext {
	return &gitconfigListenerContext{
		isWaitingForNextSection: false,
	}
}

type gitconfigParserListener struct {
	*parser.BaseConfigListener
	Config           *GitConfig
	Errors           []common.LSPError
	gitconfigContext *gitconfigListenerContext
}

func createListener(
	config *GitConfig,
	context *gitconfigListenerContext,
) gitconfigParserListener {
	return gitconfigParserListener{
		Config:           config,
		Errors:           make([]common.LSPError, 0),
		gitconfigContext: context,
	}
}

func (s *gitconfigParserListener) EnterEntry(ctx *parser.EntryContext) {
	location := common.CharacterRangeFromCtx(ctx.BaseParserRuleContext)
	location.ChangeBothLines(s.gitconfigContext.line)

	s.gitconfigContext.currentEntry = &GitEntry{
		LocationRange: location,
	}

	s.gitconfigContext.currentSection.Entries = append(
		s.gitconfigContext.currentSection.Entries,
		s.gitconfigContext.currentEntry,
	)
	s.gitconfigContext.currentSection.End = location.End
}

func (s *gitconfigParserListener) ExitEntry(ctx *parser.EntryContext) {
	if s.gitconfigContext.currentEntry.Value != nil {
		s.gitconfigContext.currentEntry.End = s.gitconfigContext.currentEntry.Value.Raw.End
	}

	s.gitconfigContext.currentEntry = nil
}

func (s *gitconfigParserListener) EnterKey(ctx *parser.KeyContext) {
	location := common.CharacterRangeFromCtx(ctx.BaseParserRuleContext)
	location.ChangeBothLines(s.gitconfigContext.line)

	text := ctx.GetText()
	value := commonparser.ParseRawString(text, commonparser.FullFeatures)

	s.gitconfigContext.currentEntry.Key = &GitKey{
		LocationRange: location,
		Value:         value,
	}
}

func (s *gitconfigParserListener) EnterSeparator(ctx *parser.SeparatorContext) {
	location := common.CharacterRangeFromCtx(ctx.BaseParserRuleContext)
	location.ChangeBothLines(s.gitconfigContext.line)

	text := ctx.GetText()
	value := commonparser.ParseRawString(text, commonparser.FullFeatures)

	s.gitconfigContext.currentEntry.Separator = &GitSeparator{
		LocationRange: location,
		Value:         value,
	}
}

func (s *gitconfigParserListener) EnterValue(ctx *parser.ValueContext) {
	location := common.CharacterRangeFromCtx(ctx.BaseParserRuleContext)
	location.ChangeBothLines(s.gitconfigContext.line)

	virtualLine := s.gitconfigContext.virtualLine.GetSubset(location.Start.Character, location.End.Character)
	value := commonparser.ParseRawString(
		virtualLine.GetText(),
		commonparser.ParseFeatures{
			ParseDoubleQuotes:      true,
			TrimWhitespace:         true,
			ParseEscapedCharacters: false,
			Replacements:           &map[string]string{},
		},
	).Value
	s.gitconfigContext.currentEntry.Value = &GitValue{
		Raw:   virtualLine,
		Value: value,
	}
}
