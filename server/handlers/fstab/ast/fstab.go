package ast

import (
	"config-lsp/common"
	commonparser "config-lsp/common/parser"
	"github.com/emirpasic/gods/maps/treemap"
)

type FstabFieldName string

const (
	FstabFieldSpec           FstabFieldName = "spec"
	FstabFieldMountPoint     FstabFieldName = "mountpoint"
	FstabFieldFileSystemType FstabFieldName = "filesystemtype"
	FstabFieldOptions        FstabFieldName = "options"
	FstabFieldFreq           FstabFieldName = "freq"
	FstabFieldFsck           FstabFieldName = "fsck"
)

type FstabField struct {
	common.LocationRange
	Value commonparser.ParsedString
}

type FstabFields struct {
	common.LocationRange
	Spec           *FstabField
	MountPoint     *FstabField
	FilesystemType *FstabField
	Options        *FstabField
	Freq           *FstabField
	Fsck           *FstabField
}

type FstabEntry struct {
	Fields *FstabFields
}

type FstabConfig struct {
	// [uint32]FstabEntry - line number to line mapping
	Entries *treemap.Map

	// [uint32]{} - line number to empty struct for comments
	CommentLines map[uint32]struct{}
}
