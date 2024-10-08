package ast

// func (c FstabConfig) GetEntry(line uint32) *FstabEntry {
// 	entry, found := c.Entries.Get(line)
//
// 	if !found {
// 		return nil
// 	}
//
// 	return entry.(*FstabEntry)
// }

func (e FstabEntry) GetFieldAtPosition(cursor uint32) FstabFieldName {
	if e.Fields.Spec == nil || (cursor >= e.Fields.Spec.Start.Character && cursor <= e.Fields.Spec.End.Character) {
		return FstabFieldSpec
	}

	if e.Fields.MountPoint == nil || (cursor >= e.Fields.MountPoint.Start.Character && cursor <= e.Fields.MountPoint.End.Character) {
		return FstabFieldMountPoint
	}

	if e.Fields.FilesystemType == nil || (cursor >= e.Fields.FilesystemType.Start.Character && cursor <= e.Fields.FilesystemType.End.Character) {
		return FstabFieldFileSystemType
	}

	if e.Fields.Options == nil || (cursor >= e.Fields.Options.Start.Character && cursor <= e.Fields.Options.End.Character) {
		return FstabFieldOptions
	}

	if e.Fields.Freq == nil || (cursor >= e.Fields.Freq.Start.Character && cursor <= e.Fields.Freq.End.Character) {
		return FstabFieldFreq
	}

	return FstabFieldPass
}
