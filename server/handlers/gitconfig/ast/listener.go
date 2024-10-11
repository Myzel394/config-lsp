package ast

import (
	"config-lsp/common"
	commonparser "config-lsp/common/parser"
	"config-lsp/handlers/gitconfig/ast/parser"
)

type gitconfigListenerContext struct {
	line                uint32
	currentSection        *GitSection
	currentEntry *GitEntry
}

func createListenerContext() *gitconfigListenerContext {
	context := new(gitconfigListenerContext)

	return context
}

type gitconfigParserListener struct {
	*parser.BaseConfigListener
	Config     *GitConfig
	Errors     []common.LSPError
	gitconfigContext *gitconfigListenerContext
}

func createListener(
	config *GitConfig,
	context *gitconfigListenerContext,
) gitconfigParserListener {
	return gitconfigParserListener{
		Config:     config,
		Errors:     make([]common.LSPError, 0),
		gitconfigContext: context,
	}
}

func (s *gitconfigParserListener) EnterEntry(ctx *parser.EntryContext) {
	location := common.CharacterRangeFromCtx(ctx.BaseParserRuleContext)
	location.ChangeBothLines(s.gitconfigContext.line)

	s.gitconfigContext.currentEntry = &GitEntry{
		LocationRange: location,
	}

	s.gitconfigContext.currentSection.Entries.Put(
		location.Start.Line,
		s.gitconfigContext.currentEntry,
	)
}

func (s *gitconfigParserListener) ExitEntry(ctx *parser.EntryContext) {
	s.gitconfigContext.currentEntry = nil
}

func (s *gitconfigParserListener) EnterKey(ctx *parser.KeyContext) {
	location := common.CharacterRangeFromCtx(ctx.BaseParserRuleContext)
	location.ChangeBothLines(s.gitconfigContext.line)

	text := ctx.GetText()
	value := commonparser.ParseRawString(text, commonparser.FullFeatures)

	s.gitconfigContext.currentEntry.Key = &GitKey{
		LocationRange: location,
		Value: value,
	}
}

func (s *gitconfigParserListener) EnterSeparator(ctx *parser.SeparatorContext) {
	location := common.CharacterRangeFromCtx(ctx.BaseParserRuleContext)
	location.ChangeBothLines(s.gitconfigContext.line)

	text := ctx.GetText()
	value := commonparser.ParseRawString(text, commonparser.FullFeatures)

	s.gitconfigContext.currentEntry.Separator = &GitSeparator{
		LocationRange: location,
		Value: value,
	}
}

func (s *gitconfigParserListener) EnterValue(ctx *parser.ValueContext) {
	location := common.CharacterRangeFromCtx(ctx.BaseParserRuleContext)
	location.ChangeBothLines(s.gitconfigContext.line)

	text := ctx.GetText()
	value := commonparser.ParseRawString(text, commonparser.FullFeatures)

	s.gitconfigContext.currentEntry.Value = &GitValue{
		LocationRange: location,
		Value: value,
	}
}

