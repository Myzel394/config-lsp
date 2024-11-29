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

// LABEL=test  ext4 defaults 0 0
func (e FstabEntry) GetFieldAtPosition(position common.Position) FstabFieldName {
	// No fields defined, empty line
	if e.Fields.Spec == nil && e.Fields.MountPoint == nil && e.Fields.FilesystemType == nil && e.Fields.Options == nil && e.Fields.Freq == nil && e.Fields.Fsck == nil {
		return FstabFieldSpec
	}

	// First, try if out of the existing fields the user wants to edit one of them

	if e.Fields.Spec != nil && e.Fields.Spec.ContainsPosition(position) {
		return FstabFieldSpec
	}
	if e.Fields.MountPoint != nil && e.Fields.MountPoint.ContainsPosition(position) {
		return FstabFieldMountPoint
	}
	if e.Fields.FilesystemType != nil && e.Fields.FilesystemType.ContainsPosition(position) {
		return FstabFieldFileSystemType
	}
	if e.Fields.Options != nil && e.Fields.Options.ContainsPosition(position) {
		return FstabFieldOptions
	}
	if e.Fields.Freq != nil && e.Fields.Freq.ContainsPosition(position) {
		return FstabFieldFreq
	}
	if e.Fields.Fsck != nil && e.Fields.Fsck.ContainsPosition(position) {
		return FstabFieldFsck
	}

	// Okay let's try to fetch the field by assuming the user is typing from left to right normally

	if e.Fields.Spec != nil && e.Fields.Spec.IsPositionAfterEnd(position) && (e.Fields.MountPoint == nil || e.Fields.MountPoint.IsPositionBeforeEnd(position)) {
		return FstabFieldMountPoint
	}

	if e.Fields.MountPoint != nil && e.Fields.MountPoint.IsPositionAfterEnd(position) && (e.Fields.FilesystemType == nil || e.Fields.FilesystemType.IsPositionBeforeEnd(position)) {
		return FstabFieldFileSystemType
	}

	if e.Fields.FilesystemType != nil && e.Fields.FilesystemType.IsPositionAfterEnd(position) && (e.Fields.Options == nil || e.Fields.Options.IsPositionBeforeEnd(position)) {
		return FstabFieldOptions
	}

	if e.Fields.Options != nil && e.Fields.Options.IsPositionAfterEnd(position) && (e.Fields.Freq == nil || e.Fields.Freq.IsPositionBeforeEnd(position)) {
		return FstabFieldFreq
	}

	if e.Fields.Freq != nil && e.Fields.Freq.IsPositionAfterEnd(position) && (e.Fields.Fsck == nil || e.Fields.Fsck.IsPositionBeforeEnd(position)) {
		return FstabFieldFsck
	}

	// Okay shit no idea, let's just give whatever is missing

	if e.Fields.Spec == nil {
		return FstabFieldSpec
	}

	if e.Fields.MountPoint == nil {
		return FstabFieldMountPoint
	}

	if e.Fields.FilesystemType == nil {
		return FstabFieldFileSystemType
	}

	if e.Fields.Options == nil {
		return FstabFieldOptions
	}

	if e.Fields.Freq == nil {
		return FstabFieldFreq
	}

	return FstabFieldFsck
}

// LABEL=test /mnt/test btrfs subvol=backup,fat=32 [0] [0]
func (e FstabEntry) getCursorIndex() uint8 {
	definedAmount := e.getDefinedFieldsAmount()

	switch definedAmount {
	case 5:

	}

	return 0
}

func (e FstabEntry) getDefinedFieldsAmount() uint8 {
	var definedAmount uint8 = 0

	if e.Fields.Spec != nil {
		definedAmount++
	}
	if e.Fields.MountPoint != nil {
		definedAmount++
	}
	if e.Fields.FilesystemType != nil {
		definedAmount++
	}
	if e.Fields.Options != nil {
		definedAmount++
	}
	if e.Fields.Freq != nil {
		definedAmount++
	}
	if e.Fields.Fsck != nil {
		definedAmount++
	}

	return definedAmount
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
