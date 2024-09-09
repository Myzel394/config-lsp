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

	// EnterKey is called when entering the key production.
	EnterKey(c *KeyContext)

	// EnterValue is called when entering the value production.
	EnterValue(c *ValueContext)

	// EnterLeadingComment is called when entering the leadingComment production.
	EnterLeadingComment(c *LeadingCommentContext)

	// ExitLineStatement is called when exiting the lineStatement production.
	ExitLineStatement(c *LineStatementContext)

	// ExitEntry is called when exiting the entry production.
	ExitEntry(c *EntryContext)

	// ExitKey is called when exiting the key production.
	ExitKey(c *KeyContext)

	// ExitValue is called when exiting the value production.
	ExitValue(c *ValueContext)

	// ExitLeadingComment is called when exiting the leadingComment production.
	ExitLeadingComment(c *LeadingCommentContext)
}
