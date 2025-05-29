// Code generated from Match.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // Match

import "github.com/antlr4-go/antlr/v4"

// BaseMatchListener is a complete listener for a parse tree produced by MatchParser.
type BaseMatchListener struct{}

var _ MatchListener = &BaseMatchListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseMatchListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseMatchListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseMatchListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseMatchListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterRoot is called when production root is entered.
func (s *BaseMatchListener) EnterRoot(ctx *RootContext) {}

// ExitRoot is called when production root is exited.
func (s *BaseMatchListener) ExitRoot(ctx *RootContext) {}

// EnterMatchEntry is called when production matchEntry is entered.
func (s *BaseMatchListener) EnterMatchEntry(ctx *MatchEntryContext) {}

// ExitMatchEntry is called when production matchEntry is exited.
func (s *BaseMatchListener) ExitMatchEntry(ctx *MatchEntryContext) {}

// EnterEntrySingle is called when production entrySingle is entered.
func (s *BaseMatchListener) EnterEntrySingle(ctx *EntrySingleContext) {}

// ExitEntrySingle is called when production entrySingle is exited.
func (s *BaseMatchListener) ExitEntrySingle(ctx *EntrySingleContext) {}

// EnterEntryWithValue is called when production entryWithValue is entered.
func (s *BaseMatchListener) EnterEntryWithValue(ctx *EntryWithValueContext) {}

// ExitEntryWithValue is called when production entryWithValue is exited.
func (s *BaseMatchListener) ExitEntryWithValue(ctx *EntryWithValueContext) {}

// EnterSeparator is called when production separator is entered.
func (s *BaseMatchListener) EnterSeparator(ctx *SeparatorContext) {}

// ExitSeparator is called when production separator is exited.
func (s *BaseMatchListener) ExitSeparator(ctx *SeparatorContext) {}

// EnterValues is called when production values is entered.
func (s *BaseMatchListener) EnterValues(ctx *ValuesContext) {}

// ExitValues is called when production values is exited.
func (s *BaseMatchListener) ExitValues(ctx *ValuesContext) {}

// EnterValue is called when production value is entered.
func (s *BaseMatchListener) EnterValue(ctx *ValueContext) {}

// ExitValue is called when production value is exited.
func (s *BaseMatchListener) ExitValue(ctx *ValueContext) {}

// EnterCriteriaSingle is called when production criteriaSingle is entered.
func (s *BaseMatchListener) EnterCriteriaSingle(ctx *CriteriaSingleContext) {}

// ExitCriteriaSingle is called when production criteriaSingle is exited.
func (s *BaseMatchListener) ExitCriteriaSingle(ctx *CriteriaSingleContext) {}

// EnterCriteriaWithValue is called when production criteriaWithValue is entered.
func (s *BaseMatchListener) EnterCriteriaWithValue(ctx *CriteriaWithValueContext) {}

// ExitCriteriaWithValue is called when production criteriaWithValue is exited.
func (s *BaseMatchListener) ExitCriteriaWithValue(ctx *CriteriaWithValueContext) {}

// EnterString is called when production string is entered.
func (s *BaseMatchListener) EnterString(ctx *StringContext) {}

// ExitString is called when production string is exited.
func (s *BaseMatchListener) ExitString(ctx *StringContext) {}
