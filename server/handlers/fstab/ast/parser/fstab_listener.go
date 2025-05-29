// Code generated from Fstab.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // Fstab

import "github.com/antlr4-go/antlr/v4"

// FstabListener is a complete listener for a parse tree produced by FstabParser.
type FstabListener interface {
	antlr.ParseTreeListener

	// EnterEntry is called when entering the entry production.
	EnterEntry(c *EntryContext)

	// EnterSpec is called when entering the spec production.
	EnterSpec(c *SpecContext)

	// EnterMountPoint is called when entering the mountPoint production.
	EnterMountPoint(c *MountPointContext)

	// EnterFileSystem is called when entering the fileSystem production.
	EnterFileSystem(c *FileSystemContext)

	// EnterMountOptions is called when entering the mountOptions production.
	EnterMountOptions(c *MountOptionsContext)

	// EnterFreq is called when entering the freq production.
	EnterFreq(c *FreqContext)

	// EnterPass is called when entering the pass production.
	EnterPass(c *PassContext)

	// ExitEntry is called when exiting the entry production.
	ExitEntry(c *EntryContext)

	// ExitSpec is called when exiting the spec production.
	ExitSpec(c *SpecContext)

	// ExitMountPoint is called when exiting the mountPoint production.
	ExitMountPoint(c *MountPointContext)

	// ExitFileSystem is called when exiting the fileSystem production.
	ExitFileSystem(c *FileSystemContext)

	// ExitMountOptions is called when exiting the mountOptions production.
	ExitMountOptions(c *MountOptionsContext)

	// ExitFreq is called when exiting the freq production.
	ExitFreq(c *FreqContext)

	// ExitPass is called when exiting the pass production.
	ExitPass(c *PassContext)
}
