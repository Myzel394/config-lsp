// Code generated from Config.g4 by ANTLR 4.13.0. DO NOT EDIT.

package parser // Config

import "github.com/antlr4-go/antlr/v4"

// BaseConfigListener is a complete listener for a parse tree produced by ConfigParser.
type BaseConfigListener struct{}

var _ ConfigListener = &BaseConfigListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseConfigListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseConfigListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseConfigListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseConfigListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterLineStatement is called when production lineStatement is entered.
func (s *BaseConfigListener) EnterLineStatement(ctx *LineStatementContext) {}

// ExitLineStatement is called when production lineStatement is exited.
func (s *BaseConfigListener) ExitLineStatement(ctx *LineStatementContext) {}

// EnterEntry is called when production entry is entered.
func (s *BaseConfigListener) EnterEntry(ctx *EntryContext) {}

// ExitEntry is called when production entry is exited.
func (s *BaseConfigListener) ExitEntry(ctx *EntryContext) {}

// EnterLeadingComment is called when production leadingComment is entered.
func (s *BaseConfigListener) EnterLeadingComment(ctx *LeadingCommentContext) {}

// ExitLeadingComment is called when production leadingComment is exited.
func (s *BaseConfigListener) ExitLeadingComment(ctx *LeadingCommentContext) {}

// EnterKey is called when production key is entered.
func (s *BaseConfigListener) EnterKey(ctx *KeyContext) {}

// ExitKey is called when production key is exited.
func (s *BaseConfigListener) ExitKey(ctx *KeyContext) {}

// EnterSeparator is called when production separator is entered.
func (s *BaseConfigListener) EnterSeparator(ctx *SeparatorContext) {}

// ExitSeparator is called when production separator is exited.
func (s *BaseConfigListener) ExitSeparator(ctx *SeparatorContext) {}

// EnterValue is called when production value is entered.
func (s *BaseConfigListener) EnterValue(ctx *ValueContext) {}

// ExitValue is called when production value is exited.
func (s *BaseConfigListener) ExitValue(ctx *ValueContext) {}

// EnterString is called when production string is entered.
func (s *BaseConfigListener) EnterString(ctx *StringContext) {}

// ExitString is called when production string is exited.
func (s *BaseConfigListener) ExitString(ctx *StringContext) {}
