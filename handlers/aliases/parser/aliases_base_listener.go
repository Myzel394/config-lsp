// Code generated from Aliases.g4 by ANTLR 4.13.0. DO NOT EDIT.

package parser // Aliases

import "github.com/antlr4-go/antlr/v4"

// BaseAliasesListener is a complete listener for a parse tree produced by AliasesParser.
type BaseAliasesListener struct{}

var _ AliasesListener = &BaseAliasesListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseAliasesListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseAliasesListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseAliasesListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseAliasesListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterLineStatement is called when production lineStatement is entered.
func (s *BaseAliasesListener) EnterLineStatement(ctx *LineStatementContext) {}

// ExitLineStatement is called when production lineStatement is exited.
func (s *BaseAliasesListener) ExitLineStatement(ctx *LineStatementContext) {}

// EnterEntry is called when production entry is entered.
func (s *BaseAliasesListener) EnterEntry(ctx *EntryContext) {}

// ExitEntry is called when production entry is exited.
func (s *BaseAliasesListener) ExitEntry(ctx *EntryContext) {}

// EnterSeparator is called when production separator is entered.
func (s *BaseAliasesListener) EnterSeparator(ctx *SeparatorContext) {}

// ExitSeparator is called when production separator is exited.
func (s *BaseAliasesListener) ExitSeparator(ctx *SeparatorContext) {}

// EnterKey is called when production key is entered.
func (s *BaseAliasesListener) EnterKey(ctx *KeyContext) {}

// ExitKey is called when production key is exited.
func (s *BaseAliasesListener) ExitKey(ctx *KeyContext) {}

// EnterValues is called when production values is entered.
func (s *BaseAliasesListener) EnterValues(ctx *ValuesContext) {}

// ExitValues is called when production values is exited.
func (s *BaseAliasesListener) ExitValues(ctx *ValuesContext) {}

// EnterValue is called when production value is entered.
func (s *BaseAliasesListener) EnterValue(ctx *ValueContext) {}

// ExitValue is called when production value is exited.
func (s *BaseAliasesListener) ExitValue(ctx *ValueContext) {}

// EnterUser is called when production user is entered.
func (s *BaseAliasesListener) EnterUser(ctx *UserContext) {}

// ExitUser is called when production user is exited.
func (s *BaseAliasesListener) ExitUser(ctx *UserContext) {}

// EnterFile is called when production file is entered.
func (s *BaseAliasesListener) EnterFile(ctx *FileContext) {}

// ExitFile is called when production file is exited.
func (s *BaseAliasesListener) ExitFile(ctx *FileContext) {}

// EnterCommand is called when production command is entered.
func (s *BaseAliasesListener) EnterCommand(ctx *CommandContext) {}

// ExitCommand is called when production command is exited.
func (s *BaseAliasesListener) ExitCommand(ctx *CommandContext) {}

// EnterInclude is called when production include is entered.
func (s *BaseAliasesListener) EnterInclude(ctx *IncludeContext) {}

// ExitInclude is called when production include is exited.
func (s *BaseAliasesListener) ExitInclude(ctx *IncludeContext) {}

// EnterComment is called when production comment is entered.
func (s *BaseAliasesListener) EnterComment(ctx *CommentContext) {}

// ExitComment is called when production comment is exited.
func (s *BaseAliasesListener) ExitComment(ctx *CommentContext) {}

// EnterEmail is called when production email is entered.
func (s *BaseAliasesListener) EnterEmail(ctx *EmailContext) {}

// ExitEmail is called when production email is exited.
func (s *BaseAliasesListener) ExitEmail(ctx *EmailContext) {}

// EnterError is called when production error is entered.
func (s *BaseAliasesListener) EnterError(ctx *ErrorContext) {}

// ExitError is called when production error is exited.
func (s *BaseAliasesListener) ExitError(ctx *ErrorContext) {}

// EnterErrorStatus is called when production errorStatus is entered.
func (s *BaseAliasesListener) EnterErrorStatus(ctx *ErrorStatusContext) {}

// ExitErrorStatus is called when production errorStatus is exited.
func (s *BaseAliasesListener) ExitErrorStatus(ctx *ErrorStatusContext) {}

// EnterErrorCode is called when production errorCode is entered.
func (s *BaseAliasesListener) EnterErrorCode(ctx *ErrorCodeContext) {}

// ExitErrorCode is called when production errorCode is exited.
func (s *BaseAliasesListener) ExitErrorCode(ctx *ErrorCodeContext) {}

// EnterErrorMessage is called when production errorMessage is entered.
func (s *BaseAliasesListener) EnterErrorMessage(ctx *ErrorMessageContext) {}

// ExitErrorMessage is called when production errorMessage is exited.
func (s *BaseAliasesListener) ExitErrorMessage(ctx *ErrorMessageContext) {}
