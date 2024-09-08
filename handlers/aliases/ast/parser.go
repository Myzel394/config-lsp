package ast

import (
	"config-lsp/common"
	"config-lsp/handlers/aliases/parser"
	"config-lsp/utils"
	"regexp"

	"github.com/antlr4-go/antlr/v4"
	"github.com/emirpasic/gods/maps/treemap"
	gods "github.com/emirpasic/gods/utils"
)

func NewAliasesParser() AliasesParser {
	p := AliasesParser{}
	p.Clear()

	return p
}

func (p *AliasesParser) Clear() {
	p.CommentLines = make(map[uint32]struct{})
	p.Aliases = treemap.NewWith(gods.UInt32Comparator)
}

var commentPattern = regexp.MustCompile(`^\s*#.*$`)
var emptyPattern = regexp.MustCompile(`^\s*$`)

func (p *AliasesParser) Parse(input string) []common.LSPError {
	errors := make([]common.LSPError, 0)
	lines := utils.SplitIntoLines(input)

	for rawLineNumber, line := range lines {
		lineNumber := uint32(rawLineNumber)

		if commentPattern.MatchString(line) {
			p.CommentLines[lineNumber] = struct{}{}
			continue
		}

		if emptyPattern.MatchString(line) {
			continue
		}

		errors = append(
			errors,
			p.parseStatement(lineNumber, line)...,
		)
	}

	return errors
}

func (p *AliasesParser) parseStatement(
	line uint32,
	input string,
) []common.LSPError {
	stream := antlr.NewInputStream(input)

	lexerErrorListener := createErrorListener(line)
	lexer := parser.NewAliasesLexer(stream)
	lexer.RemoveErrorListeners()
	lexer.AddErrorListener(&lexerErrorListener)

	tokenStream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	parserErrorListener := createErrorListener(line)
	antlrParser := parser.NewAliasesParser(tokenStream)
	antlrParser.RemoveErrorListeners()
	antlrParser.AddErrorListener(&parserErrorListener)

	listener := createListener(p, line)
	antlr.ParseTreeWalkerDefault.Walk(
		&listener,
		antlrParser.LineStatement(),
	)

	errors := lexerErrorListener.Errors
	errors = append(errors, parserErrorListener.Errors...)

	return errors
}
