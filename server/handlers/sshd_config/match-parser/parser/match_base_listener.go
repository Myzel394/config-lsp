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

// EnterSeparator is called when production separator is entered.
func (s *BaseMatchListener) EnterSeparator(ctx *SeparatorContext) {}

// ExitSeparator is called when production separator is exited.
func (s *BaseMatchListener) ExitSeparator(ctx *SeparatorContext) {}

// EnterCriteria is called when production criteria is entered.
func (s *BaseMatchListener) EnterCriteria(ctx *CriteriaContext) {}

// ExitCriteria is called when production criteria is exited.
func (s *BaseMatchListener) ExitCriteria(ctx *CriteriaContext) {}

// EnterValues is called when production values is entered.
func (s *BaseMatchListener) EnterValues(ctx *ValuesContext) {}

// ExitValues is called when production values is exited.
func (s *BaseMatchListener) ExitValues(ctx *ValuesContext) {}

// EnterValue is called when production value is entered.
func (s *BaseMatchListener) EnterValue(ctx *ValueContext) {}

// ExitValue is called when production value is exited.
func (s *BaseMatchListener) ExitValue(ctx *ValueContext) {}
