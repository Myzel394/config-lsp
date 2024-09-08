package fstab

import (
	fstabdocumentation "config-lsp/handlers/fstab/documentation"
	handlers "config-lsp/handlers/fstab/handlers"
	"config-lsp/handlers/fstab/parser"
	"config-lsp/utils"
	"testing"
)

func TestValidBasicExample(t *testing.T) {
	input := utils.Dedent(`
LABEL=test /mnt/test ext4 defaults 0 0
`)
	p := parser.FstabParser{}
	p.Clear()

	errors := p.ParseFromContent(input)

	if len(errors) > 0 {
		t.Fatal("ParseFromContent failed with error", errors)
	}

	// Get hover for first field
	rawEntry, _ := p.Entries.Get(uint32(0))
	entry := rawEntry.(parser.FstabEntry)

	println("Getting hover info")
	{
		hover, err := handlers.GetHoverInfo(&entry, uint32(0))

		if err != nil {
			t.Fatal("getHoverInfo failed with error", err)
		}

		if hover.Contents != handlers.SpecHoverField.Contents {
			t.Fatal("getHoverInfo failed to return correct hover content. Got:", hover.Contents, "but expected:", handlers.SpecHoverField.Contents)
		}

		// Get hover for second field
		hover, err = handlers.GetHoverInfo(&entry, uint32(11))
		if err != nil {
			t.Fatal("getHoverInfo failed with error", err)
		}

		if hover.Contents != handlers.MountPointHoverField.Contents {
			t.Fatal("getHoverInfo failed to return correct hover content. Got:", hover.Contents, "but expected:", handlers.MountPointHoverField.Contents)
		}

		hover, err = handlers.GetHoverInfo(&entry, uint32(20))

		if err != nil {
			t.Fatal("getHoverInfo failed with error", err)
		}

		if hover.Contents != handlers.MountPointHoverField.Contents {
			t.Fatal("getHoverInfo failed to return correct hover content. Got:", hover.Contents, "but expected:", handlers.MountPointHoverField.Contents)
		}
	}

	println("Getting completions")
	{
		completions, err := handlers.GetCompletion(entry.Line, uint32(0))

		if err != nil {
			t.Fatal("getCompletion failed with error", err)
		}

		if len(completions) != 4 {
			t.Fatal("getCompletion failed to return correct number of completions. Got:", len(completions), "but expected:", 4)
		}

		if !(completions[0].Label == "LABEL" ||
			completions[1].Label == "LABEL" ||
			completions[2].Label == "LABEL" ||
			completions[3].Label == "LABEL") {
			t.Fatal("getCompletion failed to return correct label. Got:", completions[0].Label, "but expected:", "LABEL")
		}
	}

	{
		completions, err := handlers.GetCompletion(entry.Line, uint32(21))

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
		diagnostics := p.AnalyzeValues()

		if len(diagnostics) > 0 {
			t.Fatal("AnalyzeValues failed with error", diagnostics)
		}
	}
}

func TestInvalidOptionsExample(t *testing.T) {
	input := utils.Dedent(`
LABEL=test /mnt/test btrfs subvol=backup,fat=32 0 0
`)
	p := parser.FstabParser{}
	p.Clear()

	errors := p.ParseFromContent(input)

	if len(errors) > 0 {
		t.Fatal("ParseFromContent returned error", errors)
	}

	// Get hover for first field
	println("Checking values")
	{
		diagnostics := p.AnalyzeValues()

		if len(diagnostics) == 0 {
			t.Fatal("AnalyzeValues should have returned error")
		}
	}
}
