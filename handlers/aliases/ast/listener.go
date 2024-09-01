package ast

import (
	"config-lsp/common"
	"config-lsp/handlers/aliases/parser"

	"github.com/antlr4-go/antlr/v4"
)

type aliasesListenerContext struct {
	line                   uint32
	currentIncludeIndex    *uint32
	currentErrorValueIndex *uint32
}

type aliasesParserListener struct {
	*parser.BaseAliasesListener
	Parser       *AliasesParser
	Errors       []common.LSPError
	aliasContext aliasesListenerContext
}

func (s *aliasesParserListener) EnterEntry(ctx *parser.EntryContext) {
	location := common.CharacterRangeFromCtx(ctx.BaseParserRuleContext)
	location.ChangeBothLines(s.aliasContext.line)

	s.Parser.Aliases.Put(
		location.Start.Line,
		&AliasEntry{
			Location: location,
		},
	)
}

func (s *aliasesParserListener) EnterSeparator(ctx *parser.SeparatorContext) {
	location := common.CharacterRangeFromCtx(ctx.BaseParserRuleContext)
	location.ChangeBothLines(s.aliasContext.line)

	rawEntry, _ := s.Parser.Aliases.Get(location.Start.Line)
	entry := rawEntry.(*AliasEntry)

	entry.Separator = &location
}

func (s *aliasesParserListener) EnterKey(ctx *parser.KeyContext) {
	location := common.CharacterRangeFromCtx(ctx.BaseParserRuleContext)
	location.ChangeBothLines(s.aliasContext.line)

	rawEntry, _ := s.Parser.Aliases.Get(location.Start.Line)
	entry := rawEntry.(*AliasEntry)

	entry.Key = &AliasKey{
		Location: location,
		Value:    ctx.GetText(),
	}
}

func (s *aliasesParserListener) EnterValues(ctx *parser.ValuesContext) {
	location := common.CharacterRangeFromCtx(ctx.BaseParserRuleContext)
	location.ChangeBothLines(s.aliasContext.line)

	rawEntry, _ := s.Parser.Aliases.Get(location.Start.Line)
	entry := rawEntry.(*AliasEntry)

	entry.Values = &AliasValues{
		Location: location,
		Values:   make([]AliasValueInterface, 0, 5),
	}
}

// === Value === //

func (s *aliasesParserListener) EnterUser(ctx *parser.UserContext) {
	location := common.CharacterRangeFromCtx(ctx.BaseParserRuleContext)
	location.ChangeBothLines(s.aliasContext.line)

	user := AliasValueUser{
		AliasValue: AliasValue{
			Location: location,
			Value:    ctx.GetText(),
		},
	}

	rawEntry, _ := s.Parser.Aliases.Get(location.Start.Line)
	entry := rawEntry.(*AliasEntry)

	entry.Values.Values = append(entry.Values.Values, user)
}

func (s *aliasesParserListener) EnterFile(ctx *parser.FileContext) {
	location := common.CharacterRangeFromCtx(ctx.BaseParserRuleContext)
	location.ChangeBothLines(s.aliasContext.line)

	rawEntry, _ := s.Parser.Aliases.Get(location.Start.Line)
	entry := rawEntry.(*AliasEntry)

	if s.aliasContext.currentIncludeIndex != nil {
		// This `file` is inside an `include`, so we need to set the path on the include
		values := entry.Values
		rawValue := values.Values[*s.aliasContext.currentIncludeIndex]

		// Set the path
		include := rawValue.(AliasValueInclude)
		include.Path = &AliasValueIncludePath{
			Location: location,
			Path:     path(ctx.GetText()),
		}
		values.Values[*s.aliasContext.currentIncludeIndex] = include

		// Clean up
		s.aliasContext.currentIncludeIndex = nil

		return
	}

	// Raw file, process it
	file := AliasValueFile{
		AliasValue: AliasValue{
			Location: location,
			Value:    ctx.GetText(),
		},
		Path: path(ctx.GetText()),
	}

	entry.Values.Values = append(entry.Values.Values, file)
}

func (s *aliasesParserListener) EnterCommand(ctx *parser.CommandContext) {
	location := common.CharacterRangeFromCtx(ctx.BaseParserRuleContext)
	location.ChangeBothLines(s.aliasContext.line)

	command := AliasValueCommand{
		AliasValue: AliasValue{
			Location: location,
			Value:    ctx.GetText(),
		},
		Command: ctx.GetText()[1:],
	}

	rawEntry, _ := s.Parser.Aliases.Get(location.Start.Line)
	entry := rawEntry.(*AliasEntry)

	entry.Values.Values = append(entry.Values.Values, command)
}

func (s *aliasesParserListener) EnterInclude(ctx *parser.IncludeContext) {
	location := common.CharacterRangeFromCtx(ctx.BaseParserRuleContext)
	location.ChangeBothLines(s.aliasContext.line)

	include := AliasValueInclude{
		AliasValue: AliasValue{
			Location: location,
			Value:    ctx.GetText(),
		},
	}

	rawEntry, _ := s.Parser.Aliases.Get(location.Start.Line)
	entry := rawEntry.(*AliasEntry)

	entry.Values.Values = append(entry.Values.Values, include)

	index := uint32(len(entry.Values.Values) - 1)
	s.aliasContext.currentIncludeIndex = &index
}

func (s *aliasesParserListener) EnterEmail(ctx *parser.EmailContext) {
	location := common.CharacterRangeFromCtx(ctx.BaseParserRuleContext)
	location.ChangeBothLines(s.aliasContext.line)

	email := AliasValueEmail{
		AliasValue: AliasValue{
			Location: location,
			Value:    ctx.GetText(),
		},
	}

	rawEntry, _ := s.Parser.Aliases.Get(location.Start.Line)
	entry := rawEntry.(*AliasEntry)

	entry.Values.Values = append(entry.Values.Values, &email)
}

func (s *aliasesParserListener) EnterError(ctx *parser.ErrorContext) {
	location := common.CharacterRangeFromCtx(ctx.BaseParserRuleContext)
	location.ChangeBothLines(s.aliasContext.line)

	errorValue := AliasValueError{
		AliasValue: AliasValue{
			Location: location,
			Value:    ctx.GetText(),
		},
	}

	rawEntry, _ := s.Parser.Aliases.Get(location.Start.Line)
	entry := rawEntry.(*AliasEntry)

	entry.Values.Values = append(entry.Values.Values, errorValue)

	index := uint32(len(entry.Values.Values) - 1)
	s.aliasContext.currentErrorValueIndex = &index
}

func (s *aliasesParserListener) ExitError(ctx *parser.ErrorContext) {
	s.aliasContext.currentErrorValueIndex = nil
}

// EnterErrorCode is called when production errorCode is entered.
func (s *aliasesParserListener) EnterErrorCode(ctx *parser.ErrorCodeContext) {
	location := common.CharacterRangeFromCtx(ctx.BaseParserRuleContext)
	location.ChangeBothLines(s.aliasContext.line)

	rawEntry, _ := s.Parser.Aliases.Get(location.Start.Line)
	entry := rawEntry.(*AliasEntry)

	values := entry.Values.Values

	rawValue := values[*s.aliasContext.currentErrorValueIndex]
	value := rawValue.(AliasValueError)

	value.Code = &AliasValueErrorCode{
		AliasValue: AliasValue{
			Location: location,
			Value:    ctx.GetText(),
		},
	}

	values[*s.aliasContext.currentErrorValueIndex] = value
}

// EnterErrorMessage is called when production errorMessage is entered.
func (s *aliasesParserListener) EnterErrorMessage(ctx *parser.ErrorMessageContext) {
	location := common.CharacterRangeFromCtx(ctx.BaseParserRuleContext)
	location.ChangeBothLines(s.aliasContext.line)

	rawEntry, _ := s.Parser.Aliases.Get(location.Start.Line)
	entry := rawEntry.(*AliasEntry)

	values := entry.Values.Values

	rawValue := values[*s.aliasContext.currentErrorValueIndex]
	value := rawValue.(AliasValueError)

	value.Message = &AliasValueErrorMessage{
		AliasValue: AliasValue{
			Location: location,
			Value:    ctx.GetText(),
		},
	}

	values[*s.aliasContext.currentErrorValueIndex] = value
}

func createListener(
	parser *AliasesParser,
	line uint32,
) aliasesParserListener {
	return aliasesParserListener{
		Parser: parser,
		aliasContext: aliasesListenerContext{
			line: line,
		},
		Errors: make([]common.LSPError, 0),
	}
}

type errorListener struct {
	*antlr.DefaultErrorListener
	Errors  []common.LSPError
	context aliasesListenerContext
}

func (d *errorListener) SyntaxError(
	recognizer antlr.Recognizer,
	offendingSymbol interface{},
	_ int,
	character int,
	message string,
	error antlr.RecognitionException,
) {
	line := d.context.line
	d.Errors = append(d.Errors, common.LSPError{
		Range: common.CreateSingleCharRange(uint32(line), uint32(character)),
		Err: common.SyntaxError{
			Message: message,
		},
	})
}

func createErrorListener(
	line uint32,
) errorListener {
	return errorListener{
		Errors: make([]common.LSPError, 0),
		context: aliasesListenerContext{
			line: line,
		},
	}
}
