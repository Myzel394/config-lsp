package fstab

import (
	fstabdocumentation "config-lsp/handlers/fstab/documentation"
	"config-lsp/utils"
	"testing"
)

var sampleValidBasicExample = `
LABEL=test /mnt/test ext4 defaults 0 0
`
var sampleInvalidOptionsExample = `
LABEL=test /mnt/test btrfs subvol=backup,fat=32 0 0
`

func TestValidBasicExample(t *testing.T) {
	// Arrange
	parser := FstabParser{}

	errors := parser.ParseFromContent(sampleValidBasicExample)

	if len(errors) > 0 {
		t.Fatal("ParseFromContent failed with error", errors)
	}

	// Get hover for first field
	entry := parser.entries[0]

	println("Getting hover info")
	{
		hover, err := getHoverInfo(&entry, uint32(0))

		if err != nil {
			t.Fatal("getHoverInfo failed with error", err)
		}

		if hover.Contents != SpecHoverField.Contents {
			t.Fatal("getHoverInfo failed to return correct hover content. Got:", hover.Contents, "but expected:", SpecHoverField.Contents)
		}

		// Get hover for second field
		hover, err = getHoverInfo(&entry, uint32(11))
		if err != nil {
			t.Fatal("getHoverInfo failed with error", err)
		}

		if hover.Contents != MountPointHoverField.Contents {
			t.Fatal("getHoverInfo failed to return correct hover content. Got:", hover.Contents, "but expected:", MountPointHoverField.Contents)
		}

		hover, err = getHoverInfo(&entry, uint32(20))

		if err != nil {
			t.Fatal("getHoverInfo failed with error", err)
		}

		if hover.Contents != MountPointHoverField.Contents {
			t.Fatal("getHoverInfo failed to return correct hover content. Got:", hover.Contents, "but expected:", MountPointHoverField.Contents)
		}
	}

	println("Getting completions")
	{
		completions, err := getCompletion(entry.Line, uint32(0))

		if err != nil {
			t.Fatal("getCompletion failed with error", err)
		}

		if len(completions) != 4 {
			t.Fatal("getCompletion failed to return correct number of completions. Got:", len(completions), "but expected:", 4)
		}

		if completions[0].Label != "UUID" && completions[0].Label != "PARTUID" {
			t.Fatal("getCompletion failed to return correct label. Got:", completions[0].Label, "but expected:", "UUID")
		}
	}

	{
		completions, err := getCompletion(entry.Line, uint32(21))

		if err != nil {
			t.Fatal("getCompletion failed with error", err)
		}

		expectedLength := len(utils.KeysOfMap(fstabdocumentation.MountOptionsMapField))
		if len(completions) != expectedLength {
			t.Fatal("getCompletion failed to return correct number of completions. Got:", len(completions), "but expected:", expectedLength)
		}
	}

	println("Checking values")
	{
		diagnostics := parser.AnalyzeValues()

		if len(diagnostics) > 0 {
			t.Fatal("AnalyzeValues failed with error", diagnostics)
		}
	}
}

func TestInvalidOptionsExample(t *testing.T) {
	// Arrange
	parser := FstabParser{}

	errors := parser.ParseFromContent(sampleInvalidOptionsExample)

	if len(errors) > 0 {
		t.Fatal("ParseFromContent returned error", errors)
	}

	// Get hover for first field
	println("Checking values")
	{
		diagnostics := parser.AnalyzeValues()

		if len(diagnostics) == 0 {
			t.Fatal("AnalyzeValues should have returned error")
		}
	}
}
