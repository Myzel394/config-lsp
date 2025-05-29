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

	// EnterEntrySingle is called when entering the entrySingle production.
	EnterEntrySingle(c *EntrySingleContext)

	// EnterEntryWithValue is called when entering the entryWithValue production.
	EnterEntryWithValue(c *EntryWithValueContext)

	// EnterSeparator is called when entering the separator production.
	EnterSeparator(c *SeparatorContext)

	// EnterValues is called when entering the values production.
	EnterValues(c *ValuesContext)

	// EnterValue is called when entering the value production.
	EnterValue(c *ValueContext)

	// EnterCriteriaSingle is called when entering the criteriaSingle production.
	EnterCriteriaSingle(c *CriteriaSingleContext)

	// EnterCriteriaWithValue is called when entering the criteriaWithValue production.
	EnterCriteriaWithValue(c *CriteriaWithValueContext)

	// EnterString is called when entering the string production.
	EnterString(c *StringContext)

	// ExitRoot is called when exiting the root production.
	ExitRoot(c *RootContext)

	// ExitMatchEntry is called when exiting the matchEntry production.
	ExitMatchEntry(c *MatchEntryContext)

	// ExitEntrySingle is called when exiting the entrySingle production.
	ExitEntrySingle(c *EntrySingleContext)

	// ExitEntryWithValue is called when exiting the entryWithValue production.
	ExitEntryWithValue(c *EntryWithValueContext)

	// ExitSeparator is called when exiting the separator production.
	ExitSeparator(c *SeparatorContext)

	// ExitValues is called when exiting the values production.
	ExitValues(c *ValuesContext)

	// ExitValue is called when exiting the value production.
	ExitValue(c *ValueContext)

	// ExitCriteriaSingle is called when exiting the criteriaSingle production.
	ExitCriteriaSingle(c *CriteriaSingleContext)

	// ExitCriteriaWithValue is called when exiting the criteriaWithValue production.
	ExitCriteriaWithValue(c *CriteriaWithValueContext)

	// ExitString is called when exiting the string production.
	ExitString(c *StringContext)
}
