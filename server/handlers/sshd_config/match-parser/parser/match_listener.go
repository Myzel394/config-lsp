// Code generated from Match.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // Match

import "github.com/antlr4-go/antlr/v4"

// MatchListener is a complete listener for a parse tree produced by MatchParser.
type MatchListener interface {
	antlr.ParseTreeListener

	// EnterRoot is called when entering the root production.
	EnterRoot(c *RootContext)

	// EnterMatchEntry is called when entering the matchEntry production.
	EnterMatchEntry(c *MatchEntryContext)

	// EnterSeparator is called when entering the separator production.
	EnterSeparator(c *SeparatorContext)

	// EnterCriteria is called when entering the criteria production.
	EnterCriteria(c *CriteriaContext)

	// EnterValues is called when entering the values production.
	EnterValues(c *ValuesContext)

	// EnterValue is called when entering the value production.
	EnterValue(c *ValueContext)

	// ExitRoot is called when exiting the root production.
	ExitRoot(c *RootContext)

	// ExitMatchEntry is called when exiting the matchEntry production.
	ExitMatchEntry(c *MatchEntryContext)

	// ExitSeparator is called when exiting the separator production.
	ExitSeparator(c *SeparatorContext)

	// ExitCriteria is called when exiting the criteria production.
	ExitCriteria(c *CriteriaContext)

	// ExitValues is called when exiting the values production.
	ExitValues(c *ValuesContext)

	// ExitValue is called when exiting the value production.
	ExitValue(c *ValueContext)
}
