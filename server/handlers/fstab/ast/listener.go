package ast

import (
	"config-lsp/common"
	"config-lsp/handlers/fstab/ast/parser"

	commonparser "config-lsp/common/parser"
)

type fstabListenerContext struct {
	line         uint32
	currentEntry *FstabEntry
}

func createListenerContext() *fstabListenerContext {
	context := new(fstabListenerContext)

	return context
}

type fstabParserListener struct {
	*parser.BaseFstabListener
	Config       *FstabConfig
	Errors       []common.LSPError
	fstabContext *fstabListenerContext
}

func createListener(
	config *FstabConfig,
	context *fstabListenerContext,
) fstabParserListener {
	return fstabParserListener{
		Config:       config,
		Errors:       make([]common.LSPError, 0),
		fstabContext: context,
	}
}

func (s *fstabParserListener) EnterEntry(ctx *parser.EntryContext) {
	location := common.CharacterRangeFromCtx(ctx.BaseParserRuleContext)
	location.ChangeBothLines(s.fstabContext.line)

	s.fstabContext.currentEntry = &FstabEntry{
		Fields: &FstabFields{
			LocationRange: location,
		},
	}

	s.Config.Entries.Put(
		s.fstabContext.line,
		s.fstabContext.currentEntry,
	)
}

func (s *fstabParserListener) ExitEntry(ctx *parser.EntryContext) {
	s.fstabContext.currentEntry = nil
}

func (s *fstabParserListener) EnterSpec(ctx *parser.SpecContext) {
	location := common.CharacterRangeFromCtx(ctx.BaseParserRuleContext)
	location.ChangeBothLines(s.fstabContext.line)

	text := ctx.GetText()
	value := commonparser.ParseRawString(text, commonparser.FullFeatures)

	s.fstabContext.currentEntry.Fields.Spec = &FstabField{
		Value:         value,
		LocationRange: location,
	}
}

func (s *fstabParserListener) EnterMountPoint(ctx *parser.MountPointContext) {
	location := common.CharacterRangeFromCtx(ctx.BaseParserRuleContext)
	location.ChangeBothLines(s.fstabContext.line)

	text := ctx.GetText()
	value := commonparser.ParseRawString(text, commonparser.FullFeatures)

	s.fstabContext.currentEntry.Fields.MountPoint = &FstabField{
		LocationRange: location,
		Value:         value,
	}
}

func (s *fstabParserListener) EnterFileSystem(ctx *parser.FileSystemContext) {
	location := common.CharacterRangeFromCtx(ctx.BaseParserRuleContext)
	location.ChangeBothLines(s.fstabContext.line)

	text := ctx.GetText()
	value := commonparser.ParseRawString(text, commonparser.FullFeatures)

	s.fstabContext.currentEntry.Fields.FilesystemType = &FstabField{
		LocationRange: location,
		Value:         value,
	}
}

func (s *fstabParserListener) EnterMountOptions(ctx *parser.MountOptionsContext) {
	location := common.CharacterRangeFromCtx(ctx.BaseParserRuleContext)
	location.ChangeBothLines(s.fstabContext.line)

	text := ctx.GetText()
	value := commonparser.ParseRawString(text, commonparser.FullFeatures)

	s.fstabContext.currentEntry.Fields.Options = &FstabField{
		LocationRange: location,
		Value:         value,
	}
}

func (s *fstabParserListener) EnterFreq(ctx *parser.FreqContext) {
	location := common.CharacterRangeFromCtx(ctx.BaseParserRuleContext)
	location.ChangeBothLines(s.fstabContext.line)

	text := ctx.GetText()
	value := commonparser.ParseRawString(text, commonparser.FullFeatures)

	s.fstabContext.currentEntry.Fields.Freq = &FstabField{
		LocationRange: location,
		Value:         value,
	}
}

func (s *fstabParserListener) EnterPass(ctx *parser.PassContext) {
	location := common.CharacterRangeFromCtx(ctx.BaseParserRuleContext)
	location.ChangeBothLines(s.fstabContext.line)

	text := ctx.GetText()
	value := commonparser.ParseRawString(text, commonparser.FullFeatures)

	s.fstabContext.currentEntry.Fields.Pass = &FstabField{
		LocationRange: location,
		Value:         value,
	}
}
