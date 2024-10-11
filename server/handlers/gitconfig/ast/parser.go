package ast

import (
	"config-lsp/common"
	"config-lsp/utils"
	"regexp"
	"strings"

	"github.com/antlr4-go/antlr/v4"
	protocol "github.com/tliron/glsp/protocol_3_16"
)

func NewGitConfig() *GitConfig {
	config := &GitConfig{}
	config.Clear()

	return config
}

var commentPattern = regexp.MustCompile(`^\s*#`)
var emptyLinePattern = regexp.MustCompile(`^\s*$`)
var headerPattern = regexp.MustCompile(`^\s*\[`)

func (c *GitConfig) Parse(input string) []common.LSPError {
	errors := make([]common.LSPError, 0)
	lines := utils.SplitIntoLines(input)
	context := createListenerContext()

	for rawLineNumber, line := range lines {
		lineNumber := uint32(rawLineNumber)

		if emptyLinePattern.MatchString(line) {
			continue
		}

		if commentPattern.MatchString(line) {
			c.CommentLines[lineNumber] = struct{}{}
			continue
		}

		if headerPattern.MatchString(line) {
			c.parseHeader(context, line)
		}

		errors = append(
			errors,
			c.parseStatement(context, line)...,
		)
	}

	return errors
}

func (c *GitConfig) parseHeader(
	context *gitconfigListenerContext,
	input string,
) []protocol.Diagnostic {
	leftBracketIndex := strings.Index(input, "[")
	rightBracketIndex := strings.Index(input, "]")

	if rightBracketIndex == -1 {
		return []protocol.Diagnostic{
			{
				Range: protocol.Range{
					Start: protocol.Position{
						Line:      context.line,
						Character: 0,
					},
					End: protocol.Position{
						Line:      context.line,
						Character: uint32(len(input)),
					},
				},
				Message: `This section title is missing a closing bracket "]"`,
			},
		}
	}

	if leftBracketIndex != 0 {
		return []protocol.Diagnostic{
			{
				Range: protocol.Range{
					Start: protocol.Position{
						Line:      context.line,
						Character: 0,
					},
					End: protocol.Position{
						Line:      context.line,
						Character: uint32(leftBracketIndex),
					},
				},
				Message: `A section title should not have any characters before the opening bracket "["`,
			},
		}
	}

	if rightBracketIndex != len(input)-1 {
		return []protocol.Diagnostic{
			{
				Range: protocol.Range{
					Start: protocol.Position{
						Line:      context.line,
						Character: uint32(rightBracketIndex),
					},
					End: protocol.Position{
						Line:      context.line,
						Character: uint32(len(input)),
					},
				},
				Message: `A section title should not have any characters after the closing bracket "]"`,
			},
		}
	}

	return nil
}

func (c *GitConfig) parseStatement(
	context *gitconfigListenerContext,
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

