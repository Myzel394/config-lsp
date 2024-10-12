// Code generated from Config.g4 by ANTLR 4.13.0. DO NOT EDIT.

package parser // Config

import "github.com/antlr4-go/antlr/v4"

// ConfigListener is a complete listener for a parse tree produced by ConfigParser.
type ConfigListener interface {
	antlr.ParseTreeListener

	// EnterLineStatement is called when entering the lineStatement production.
	EnterLineStatement(c *LineStatementContext)

	// EnterEntry is called when entering the entry production.
	EnterEntry(c *EntryContext)

	// EnterLeadingComment is called when entering the leadingComment production.
	EnterLeadingComment(c *LeadingCommentContext)

	// EnterKey is called when entering the key production.
	EnterKey(c *KeyContext)

	// EnterSeparator is called when entering the separator production.
	EnterSeparator(c *SeparatorContext)

	// EnterValue is called when entering the value production.
	EnterValue(c *ValueContext)

	// EnterString is called when entering the string production.
	EnterString(c *StringContext)

	// EnterCommentSymbol is called when entering the commentSymbol production.
	EnterCommentSymbol(c *CommentSymbolContext)

	// ExitLineStatement is called when exiting the lineStatement production.
	ExitLineStatement(c *LineStatementContext)

	// ExitEntry is called when exiting the entry production.
	ExitEntry(c *EntryContext)

	// ExitLeadingComment is called when exiting the leadingComment production.
	ExitLeadingComment(c *LeadingCommentContext)

	// ExitKey is called when exiting the key production.
	ExitKey(c *KeyContext)

	// ExitSeparator is called when exiting the separator production.
	ExitSeparator(c *SeparatorContext)

	// ExitValue is called when exiting the value production.
	ExitValue(c *ValueContext)

	// ExitString is called when exiting the string production.
	ExitString(c *StringContext)

	// ExitCommentSymbol is called when exiting the commentSymbol production.
	ExitCommentSymbol(c *CommentSymbolContext)
}
