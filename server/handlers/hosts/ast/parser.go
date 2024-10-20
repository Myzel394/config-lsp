package ast

import (
	"config-lsp/common"
	parser2 "config-lsp/handlers/hosts/ast/parser"
	"config-lsp/utils"
	"regexp"

	"github.com/antlr4-go/antlr/v4"
	"github.com/emirpasic/gods/maps/treemap"

	gods "github.com/emirpasic/gods/utils"
)

func (p *HostsParser) Clear() {
	p.Tree = HostsTree{
		Entries: treemap.NewWith(gods.UInt32Comparator),
	}
	p.CommentLines = make(map[uint32]struct{})
}

var commentPattern = regexp.MustCompile(`^\s*#.*$`)
var emptyPattern = regexp.MustCompile(`^\s*$`)

func (p *HostsParser) parseStatement(
	line uint32,
	input string,
) []common.LSPError {
	stream := antlr.NewInputStream(input)

	errorListener := createErrorListener(line)
	lexer := parser2.NewHostsLexer(stream)
	lexer.RemoveErrorListeners()
	lexer.AddErrorListener(&errorListener)

	errors := errorListener.Errors

	errorListener = createErrorListener(line)
	tokenStream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	antlrParser := parser2.NewHostsParser(tokenStream)
	antlrParser.RemoveErrorListeners()
	antlrParser.AddErrorListener(&errorListener)

	listener := createHostsFileListener(p, line)
	antlr.ParseTreeWalkerDefault.Walk(
		&listener,
		antlrParser.LineStatement(),
	)

	errors = append(errors, listener.Errors...)
	errors = append(errors, errorListener.Errors...)

	return errors
}

func (p *HostsParser) Parse(input string) []common.LSPError {
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

func NewHostsParser() HostsParser {
	p := HostsParser{}
	p.Clear()

	return p
}
