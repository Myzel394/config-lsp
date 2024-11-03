package ast

import (
	"config-lsp/common"
	"config-lsp/utils"
	"testing"
)

func TestExample1(
	t *testing.T,
) {
	input := utils.Dedent(`
LABEL=test /mnt/test ext4 defaults 0 0
LABEL=example /mnt/example fat32 defaults 0 2
`)
	c := NewFstabConfig()

	errors := c.Parse(input)

	if len(errors) > 0 {
		t.Fatalf("Expected no errors, got %v", errors)
	}

	if c.Entries.Size() != 2 {
		t.Fatalf("Expected 2 entry, got %d", c.Entries.Size())
	}

	rawFirstEntry, _ := c.Entries.Get(uint32(0))
	firstEntry := rawFirstEntry.(*FstabEntry)
	if !(firstEntry.Fields.Spec.Value.Value == "LABEL=test" && firstEntry.Fields.MountPoint.Value.Value == "/mnt/test" && firstEntry.Fields.FilesystemType.Value.Value == "ext4" && firstEntry.Fields.Options.Value.Value == "defaults" && firstEntry.Fields.Freq.Value.Value == "0" && firstEntry.Fields.Fsck.Value.Value == "0") {
		t.Fatalf("Expected entry to be LABEL=test /mnt/test ext4 defaults 0 0, got %v", firstEntry)
	}

	if !(firstEntry.Fields.Spec.LocationRange.Start.Line == 0 && firstEntry.Fields.Spec.LocationRange.Start.Character == 0) {
		t.Errorf("Expected spec start to be 0:0, got %v", firstEntry.Fields.Spec.LocationRange.Start)
	}

	if !(firstEntry.Fields.Spec.LocationRange.End.Line == 0 && firstEntry.Fields.Spec.LocationRange.End.Character == 10) {
		t.Errorf("Expected spec end to be 0:10, got %v", firstEntry.Fields.Spec.LocationRange.End)
	}

	if !(firstEntry.Fields.MountPoint.LocationRange.Start.Line == 0 && firstEntry.Fields.MountPoint.LocationRange.Start.Character == 11) {
		t.Errorf("Expected mountpoint start to be 0:11, got %v", firstEntry.Fields.MountPoint.LocationRange.Start)
	}

	if !(firstEntry.Fields.MountPoint.LocationRange.End.Line == 0 && firstEntry.Fields.MountPoint.LocationRange.End.Character == 20) {
		t.Errorf("Expected mountpoint end to be 0:20, got %v", firstEntry.Fields.MountPoint.LocationRange.End)
	}

	if !(firstEntry.Fields.FilesystemType.LocationRange.Start.Line == 0 && firstEntry.Fields.FilesystemType.LocationRange.Start.Character == 21) {
		t.Errorf("Expected filesystemtype start to be 0:21, got %v", firstEntry.Fields.FilesystemType.LocationRange.Start)
	}

	if !(firstEntry.Fields.FilesystemType.LocationRange.End.Line == 0 && firstEntry.Fields.FilesystemType.LocationRange.End.Character == 25) {
		t.Errorf("Expected filesystemtype end to be 0:25, got %v", firstEntry.Fields.FilesystemType.LocationRange.End)
	}

	if !(firstEntry.Fields.Options.LocationRange.Start.Line == 0 && firstEntry.Fields.Options.LocationRange.Start.Character == 26) {
		t.Errorf("Expected options start to be 0:26, got %v", firstEntry.Fields.Options.LocationRange.Start)
	}

	if !(firstEntry.Fields.Options.LocationRange.End.Line == 0 && firstEntry.Fields.Options.LocationRange.End.Character == 34) {
		t.Errorf("Expected options end to be 0:34, got %v", firstEntry.Fields.Options.LocationRange.End)
	}

	if !(firstEntry.Fields.Freq.LocationRange.Start.Line == 0 && firstEntry.Fields.Freq.LocationRange.Start.Character == 35) {
		t.Errorf("Expected freq start to be 0:35, got %v", firstEntry.Fields.Freq.LocationRange.Start)
	}

	if !(firstEntry.Fields.Freq.LocationRange.End.Line == 0 && firstEntry.Fields.Freq.LocationRange.End.Character == 36) {
		t.Errorf("Expected freq end to be 0:36, got %v", firstEntry.Fields.Freq.LocationRange.End)
	}

	if !(firstEntry.Fields.Fsck.LocationRange.Start.Line == 0 && firstEntry.Fields.Fsck.LocationRange.Start.Character == 37) {
		t.Errorf("Expected pass start to be 0:37, got %v", firstEntry.Fields.Fsck.LocationRange.Start)
	}

	field := firstEntry.GetFieldAtPosition(common.IndexPosition(0))
	if !(field == FstabFieldSpec) {
		t.Errorf("Expected field to be spec, got %v", field)
	}

	field = firstEntry.GetFieldAtPosition(common.IndexPosition(11))
	if !(field == FstabFieldMountPoint) {
		t.Errorf("Expected field to be mountpoint, got %v", field)
	}

	field = firstEntry.GetFieldAtPosition(common.IndexPosition(33))
	if !(field == FstabFieldOptions) {
		t.Errorf("Expected field to be spec, got %v", field)
	}

	field = firstEntry.GetFieldAtPosition(common.IndexPosition(35))
	if !(field == FstabFieldFreq) {
		t.Errorf("Expected field to be freq, got %v", field)
	}

	rawSecondEntry, _ := c.Entries.Get(uint32(1))
	secondEntry := rawSecondEntry.(*FstabEntry)
	if !(secondEntry.Fields.Start.Line == 1) {
		t.Errorf("Expected start line to be 1, got %d", secondEntry.Fields.Start.Line)
	}
}

func TestIncompleteExample(
	t *testing.T,
) {
	input := utils.Dedent(`
LABEL=test /mnt/test ext4 defaults 
`)
	c := NewFstabConfig()

	errors := c.Parse(input)

	if len(errors) != 0 {
		t.Fatalf("Expected no errors, got %v", errors)
	}

	rawFirstEntry, _ := c.Entries.Get(uint32(0))
	firstEntry := rawFirstEntry.(*FstabEntry)

	if !(firstEntry.Fields.Spec.Value.Raw == "LABEL=test" && firstEntry.Fields.MountPoint.Value.Raw == "/mnt/test" && firstEntry.Fields.FilesystemType.Value.Raw == "ext4" && firstEntry.Fields.Options.Value.Raw == "defaults") {
		t.Fatalf("Expected entry to be LABEL=test /mnt/test ext4 defaults, got %v", firstEntry)
	}

	field := firstEntry.GetFieldAtPosition(common.IndexPosition(0))
	if !(field == FstabFieldSpec) {
		t.Errorf("Expected field to be spec, got %v", field)
	}

	field = firstEntry.GetFieldAtPosition(common.IndexPosition(11))
	if !(field == FstabFieldMountPoint) {
		t.Errorf("Expected field to be mountpoint, got %v", field)
	}

	field = firstEntry.GetFieldAtPosition(common.IndexPosition(33))
	if !(field == FstabFieldOptions) {
		t.Errorf("Expected field to be spec, got %v", field)
	}

	field = firstEntry.GetFieldAtPosition(common.IndexPosition(35))
	if !(field == FstabFieldFreq) {
		t.Errorf("Expected field to be freq, got %v", field)
	}
}
