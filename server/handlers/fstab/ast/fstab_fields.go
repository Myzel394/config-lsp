package ast

import (
	"config-lsp/common"
	docvalues "config-lsp/doc-values"
	"config-lsp/handlers/fstab/fields"
	"config-lsp/utils"
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

// Create a mount options field for the entry
func (e FstabEntry) FetchMountOptionsField(includeDefaults bool) docvalues.DeprecatedValue {
	if e.Fields.FilesystemType == nil {
		return nil
	}

	option, found := fields.MountOptionsMapField[e.Fields.FilesystemType.Value.Value]

	if !found {
		return nil
	}

	var enums []docvalues.EnumString
	var assignable map[docvalues.EnumString]docvalues.DeprecatedValue

	if includeDefaults {
		enums = append(option.Enums, fields.DefaultOptions...)
		assignable = utils.MergeMaps(option.Assignable, fields.DefaultAssignOptions)
	} else {
		enums = option.Enums
		assignable = option.Assignable
	}

	return &docvalues.ArrayValue{
		Separator:           ",",
		DuplicatesExtractor: &fields.MountOptionsExtractor,
		SubValue: docvalues.OrValue{
			Values: []docvalues.DeprecatedValue{
				docvalues.KeyEnumAssignmentValue{
					Values:          assignable,
					ValueIsOptional: false,
					Separator:       "=",
				},
				docvalues.EnumValue{
					EnforceValues: true,
					Values:        enums,
				},
			},
		},
	}
}
