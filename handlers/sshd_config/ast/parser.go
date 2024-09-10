package ast

import (
	"config-lsp/common"
	"config-lsp/handlers/sshd_config/ast/parser"
	"config-lsp/utils"
	"regexp"

	"github.com/antlr4-go/antlr/v4"
	"github.com/emirpasic/gods/maps/treemap"

	gods "github.com/emirpasic/gods/utils"
)

func NewSSHConfig() *SSHConfig {
	config := &SSHConfig{}
	config.Clear()

	return config
}

func (c *SSHConfig) Clear() {
	c.Options = treemap.NewWith(gods.UInt32Comparator)
	c.CommentLines = make(map[uint32]struct{})
}

var commentPattern = regexp.MustCompile(`^\s*#.*$`)
var emptyPattern = regexp.MustCompile(`^\s*$`)

func (c *SSHConfig) Parse(input string) []common.LSPError {
	errors := make([]common.LSPError, 0)
	lines := utils.SplitIntoLines(input)
	context := createSSHListenerContext()

	for rawLineNumber, line := range lines {
		context.line = uint32(rawLineNumber)

		if commentPattern.MatchString(line) {
			c.CommentLines[context.line] = struct{}{}
			continue
		}

		if emptyPattern.MatchString(line) {
			continue
		}

		errors = append(
			errors,
			c.parseStatement(context, line)...,
		)
	}

	return errors
}

func (c *SSHConfig) parseStatement(
	context *sshListenerContext,
	input string,
) []common.LSPError {
	stream := antlr.NewInputStream(input)

	lexerErrorListener := createErrorListener(context.line)
	lexer := parser.NewConfigLexer(stream)
	lexer.RemoveErrorListeners()
	lexer.AddErrorListener(&lexerErrorListener)

	tokenStream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	parserErrorListener := createErrorListener(context.line)
	antlrParser := parser.NewConfigParser(tokenStream)
	antlrParser.RemoveErrorListeners()
	antlrParser.AddErrorListener(&parserErrorListener)

	listener := createListener(c, context)
	antlr.ParseTreeWalkerDefault.Walk(
		&listener,
		antlrParser.LineStatement(),
	)

	errors := lexerErrorListener.Errors
	errors = append(errors, parserErrorListener.Errors...)
	errors = append(errors, listener.Errors...)

	return errors
}
