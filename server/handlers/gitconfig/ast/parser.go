package ast

import (
	"config-lsp/common"
	"config-lsp/handlers/gitconfig/ast/parser"
	"errors"
	"regexp"
	"strings"

	"github.com/antlr4-go/antlr/v4"
	"github.com/emirpasic/gods/maps/treemap"

	gods "github.com/emirpasic/gods/utils"
	protocol "github.com/tliron/glsp/protocol_3_16"
)

func NewGitConfig() *GitConfig {
	config := &GitConfig{}
	config.Clear()

	return config
}

var commentPattern = regexp.MustCompile(`^\s*[#;]`)
var emptyLinePattern = regexp.MustCompile(`^\s*$`)
var headerPattern = regexp.MustCompile(`^\s*\[`)

func (c *GitConfig) Parse(input string) []common.LSPError {
	errs := make([]common.LSPError, 0)
	lines := common.SplitIntoVirtualLines(input)
	context := createListenerContext()

	for _, virtualLine := range lines {
		lineNumber := uint32(virtualLine.Parts[0].Start.Line)
		line := virtualLine.GetText()
		context.line = lineNumber
		context.virtualLine = virtualLine

		if emptyLinePattern.MatchString(line) {
			continue
		}

		if commentPattern.MatchString(line) {
			c.CommentLines[lineNumber] = struct{}{}
			continue
		}

		if headerPattern.MatchString(line) {
			c.parseHeader(context, line)

			continue
		}

		if context.currentSection == nil {
			if !context.isWaitingForNextSection {
				context.isWaitingForNextSection = true

				errs = append(errs, common.LSPError{
					Range: common.LocationRange{
						Start: common.Location{
							Line:      lineNumber,
							Character: 0,
						},
						End: common.Location{
							Line:      lineNumber,
							Character: uint32(len(line)),
						},
					},
					Err: errors.New(`This section is missing a header (e.g. "[section]")`),
				})
			}

			continue
		}

		context.isWaitingForNextSection = false

		errs = append(
			errs,
			c.parseStatement(context, line)...,
		)
	}

	return errs
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

	location := common.LocationRange{
		Start: common.Location{
			Line:      context.line,
			Character: uint32(leftBracketIndex),
		},
		End: common.Location{
			Line:      context.line,
			Character: uint32(rightBracketIndex + 1),
		},
	}
	context.currentSection = &GitSection{
		LocationRange: location,
		Title: &GitSectionHeader{
			LocationRange: location,
			Title:         input[leftBracketIndex+1 : rightBracketIndex],
		},
		Entries: treemap.NewWith(gods.UInt32Comparator),
	}
	c.Sections = append(c.Sections, context.currentSection)

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
