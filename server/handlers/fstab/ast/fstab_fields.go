package ast

import (
	"config-lsp/common"
	docvalues "config-lsp/doc-values"
	"config-lsp/handlers/fstab/fields"
)

// func (c FstabConfig) GetEntry(line uint32) *FstabEntry {
// 	entry, found := c.Entries.Get(line)
//
// 	if !found {
// 		return nil
// 	}
//
// 	return entry.(*FstabEntry)
// }

func (e FstabEntry) GetFieldAtPosition(position common.Position) FstabFieldName {
	if e.Fields.Spec == nil || (e.Fields.Spec.ContainsPosition(position)) {
		return FstabFieldSpec
	}

	if e.Fields.MountPoint == nil || (e.Fields.MountPoint.ContainsPosition(position)) {
		return FstabFieldMountPoint
	}

	if e.Fields.FilesystemType == nil || (e.Fields.FilesystemType.ContainsPosition(position)) {
		return FstabFieldFileSystemType
	}

	if e.Fields.Options == nil || (e.Fields.Options.ContainsPosition(position)) {
		return FstabFieldOptions
	}

	if e.Fields.Freq == nil || (e.Fields.Freq.ContainsPosition(position)) {
		return FstabFieldFreq
	}

	return FstabFieldPass
}

func (e FstabEntry) GetMountOptionsField() docvalues.DeprecatedValue {
	fileSystemType := e.Fields.FilesystemType.Value.Value

	var optionsField docvalues.DeprecatedValue

	if foundField, found := fields.MountOptionsMapField[fileSystemType]; found {
		optionsField = foundField
	} else {
		optionsField = fields.DefaultMountOptionsField
	}

	return optionsField
}
