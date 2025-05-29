// Code generated from Fstab.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // Fstab

import "github.com/antlr4-go/antlr/v4"

// BaseFstabListener is a complete listener for a parse tree produced by FstabParser.
type BaseFstabListener struct{}

var _ FstabListener = &BaseFstabListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseFstabListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseFstabListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseFstabListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseFstabListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterEntry is called when production entry is entered.
func (s *BaseFstabListener) EnterEntry(ctx *EntryContext) {}

// ExitEntry is called when production entry is exited.
func (s *BaseFstabListener) ExitEntry(ctx *EntryContext) {}

// EnterSpec is called when production spec is entered.
func (s *BaseFstabListener) EnterSpec(ctx *SpecContext) {}

// ExitSpec is called when production spec is exited.
func (s *BaseFstabListener) ExitSpec(ctx *SpecContext) {}

// EnterMountPoint is called when production mountPoint is entered.
func (s *BaseFstabListener) EnterMountPoint(ctx *MountPointContext) {}

// ExitMountPoint is called when production mountPoint is exited.
func (s *BaseFstabListener) ExitMountPoint(ctx *MountPointContext) {}

// EnterFileSystem is called when production fileSystem is entered.
func (s *BaseFstabListener) EnterFileSystem(ctx *FileSystemContext) {}

// ExitFileSystem is called when production fileSystem is exited.
func (s *BaseFstabListener) ExitFileSystem(ctx *FileSystemContext) {}

// EnterMountOptions is called when production mountOptions is entered.
func (s *BaseFstabListener) EnterMountOptions(ctx *MountOptionsContext) {}

// ExitMountOptions is called when production mountOptions is exited.
func (s *BaseFstabListener) ExitMountOptions(ctx *MountOptionsContext) {}

// EnterFreq is called when production freq is entered.
func (s *BaseFstabListener) EnterFreq(ctx *FreqContext) {}

// ExitFreq is called when production freq is exited.
func (s *BaseFstabListener) ExitFreq(ctx *FreqContext) {}

// EnterPass is called when production pass is entered.
func (s *BaseFstabListener) EnterPass(ctx *PassContext) {}

// ExitPass is called when production pass is exited.
func (s *BaseFstabListener) ExitPass(ctx *PassContext) {}
