package ast

import (
	"config-lsp/utils"
	"testing"
)

func TestSimpleParserExample(
	t *testing.T,
) {
	input := utils.Dedent(`
PermitRootLogin no
PasswordAuthentication yes
`)
	p := NewSSHConfig()
	errors := p.Parse(input)

	if len(errors) != 0 {
		t.Fatalf("Expected no errors, got %v", errors)
	}

	if !(p.Options.Size() == 2 &&
		len(utils.KeysOfMap(p.CommentLines)) == 0) {
		t.Errorf("Expected 2 options and no comment lines, but got: %v, %v", p.Options, p.CommentLines)
	}

	rawFirstEntry, _ := p.Options.Get(uint32(0))
	firstEntry := rawFirstEntry.(*SSHOption)

	if !(firstEntry.Value == "PermitRootLogin no" &&
		firstEntry.LocationRange.Start.Line == 0 &&
		firstEntry.LocationRange.End.Line == 0 &&
		firstEntry.LocationRange.Start.Character == 0 &&
		firstEntry.LocationRange.End.Character == 18 &&
		firstEntry.Key.Value == "PermitRootLogin" &&
		firstEntry.Key.LocationRange.Start.Character == 0 &&
		firstEntry.Key.LocationRange.End.Character == 15 &&
		firstEntry.OptionValue.Value == "no" &&
		firstEntry.OptionValue.LocationRange.Start.Character == 16 &&
		firstEntry.OptionValue.LocationRange.End.Character == 18) {
		t.Errorf("Expected first entry to be PermitRootLogin no, but got: %v", firstEntry)
	}

	rawSecondEntry, _ := p.Options.Get(uint32(1))
	secondEntry := rawSecondEntry.(*SSHOption)

	if !(secondEntry.Value == "PasswordAuthentication yes" &&
		secondEntry.LocationRange.Start.Line == 1 &&
		secondEntry.LocationRange.End.Line == 1 &&
		secondEntry.LocationRange.Start.Character == 0 &&
		secondEntry.LocationRange.End.Character == 26 &&
		secondEntry.Key.Value == "PasswordAuthentication" &&
		secondEntry.Key.LocationRange.Start.Character == 0 &&
		secondEntry.Key.LocationRange.End.Character == 22 &&
		secondEntry.OptionValue.Value == "yes" &&
		secondEntry.OptionValue.LocationRange.Start.Character == 23 &&
		secondEntry.OptionValue.LocationRange.End.Character == 26) {
		t.Errorf("Expected second entry to be PasswordAuthentication yes, but got: %v", secondEntry)
	}
}
