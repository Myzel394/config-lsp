// Code generated from Aliases.g4 by ANTLR 4.13.0. DO NOT EDIT.

package parser // Aliases

import "github.com/antlr4-go/antlr/v4"

// AliasesListener is a complete listener for a parse tree produced by AliasesParser.
type AliasesListener interface {
	antlr.ParseTreeListener

	// EnterLineStatement is called when entering the lineStatement production.
	EnterLineStatement(c *LineStatementContext)

	// EnterEntry is called when entering the entry production.
	EnterEntry(c *EntryContext)

	// EnterSeparator is called when entering the separator production.
	EnterSeparator(c *SeparatorContext)

	// EnterKey is called when entering the key production.
	EnterKey(c *KeyContext)

	// EnterValues is called when entering the values production.
	EnterValues(c *ValuesContext)

	// EnterValue is called when entering the value production.
	EnterValue(c *ValueContext)

	// EnterUser is called when entering the user production.
	EnterUser(c *UserContext)

	// EnterFile is called when entering the file production.
	EnterFile(c *FileContext)

	// EnterCommand is called when entering the command production.
	EnterCommand(c *CommandContext)

	// EnterInclude is called when entering the include production.
	EnterInclude(c *IncludeContext)

	// EnterComment is called when entering the comment production.
	EnterComment(c *CommentContext)

	// EnterEmail is called when entering the email production.
	EnterEmail(c *EmailContext)

	// EnterError is called when entering the error production.
	EnterError(c *ErrorContext)

	// EnterErrorStatus is called when entering the errorStatus production.
	EnterErrorStatus(c *ErrorStatusContext)

	// EnterErrorCode is called when entering the errorCode production.
	EnterErrorCode(c *ErrorCodeContext)

	// EnterErrorMessage is called when entering the errorMessage production.
	EnterErrorMessage(c *ErrorMessageContext)

	// ExitLineStatement is called when exiting the lineStatement production.
	ExitLineStatement(c *LineStatementContext)

	// ExitEntry is called when exiting the entry production.
	ExitEntry(c *EntryContext)

	// ExitSeparator is called when exiting the separator production.
	ExitSeparator(c *SeparatorContext)

	// ExitKey is called when exiting the key production.
	ExitKey(c *KeyContext)

	// ExitValues is called when exiting the values production.
	ExitValues(c *ValuesContext)

	// ExitValue is called when exiting the value production.
	ExitValue(c *ValueContext)

	// ExitUser is called when exiting the user production.
	ExitUser(c *UserContext)

	// ExitFile is called when exiting the file production.
	ExitFile(c *FileContext)

	// ExitCommand is called when exiting the command production.
	ExitCommand(c *CommandContext)

	// ExitInclude is called when exiting the include production.
	ExitInclude(c *IncludeContext)

	// ExitComment is called when exiting the comment production.
	ExitComment(c *CommentContext)

	// ExitEmail is called when exiting the email production.
	ExitEmail(c *EmailContext)

	// ExitError is called when exiting the error production.
	ExitError(c *ErrorContext)

	// ExitErrorStatus is called when exiting the errorStatus production.
	ExitErrorStatus(c *ErrorStatusContext)

	// ExitErrorCode is called when exiting the errorCode production.
	ExitErrorCode(c *ErrorCodeContext)

	// ExitErrorMessage is called when exiting the errorMessage production.
	ExitErrorMessage(c *ErrorMessageContext)
}
