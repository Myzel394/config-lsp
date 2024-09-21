package matchparser

import (
	"config-lsp/common"
	parser2 "config-lsp/common/parsers/openssh-match-parser/parser"
	"github.com/antlr4-go/antlr/v4"
)

func NewMatch() *Match {
	match := new(Match)
	match.Clear()

	return match
}

func (m *Match) Clear() {
	m.Entries = make([]*MatchEntry, 0)
}

func (m *Match) Parse(
	input string,
	line uint32,
	startCharacter uint32,
) []common.LSPError {
	context := createMatchListenerContext(line, startCharacter)

	stream := antlr.NewInputStream(input)

	lexerErrorListener := createErrorListener(context.line)
	lexer := parser2.NewMatchLexer(stream)
	lexer.RemoveErrorListeners()
	lexer.AddErrorListener(&lexerErrorListener)

	tokenStream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	parserErrorListener := createErrorListener(context.line)
	antlrParser := parser2.NewMatchParser(tokenStream)
	antlrParser.RemoveErrorListeners()
	antlrParser.AddErrorListener(&parserErrorListener)

	listener := createListener(m, context)
	antlr.ParseTreeWalkerDefault.Walk(
		&listener,
		antlrParser.Root(),
	)

	errors := lexerErrorListener.Errors
	errors = append(errors, parserErrorListener.Errors...)
	errors = append(errors, listener.Errors...)

	return errors

}
