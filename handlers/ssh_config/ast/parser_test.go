package ast

import (
	"config-lsp/utils"
	"testing"
)

func TestSSHConfigParserExample(
	t *testing.T,
) {
	input := utils.Dedent(`
HostName 1.2.3.4
User root
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

	if !(firstEntry.Value.Value == "HostName 1.2.3.4" &&
		firstEntry.LocationRange.Start.Line == 0 &&
		firstEntry.LocationRange.End.Line == 0 &&
		firstEntry.LocationRange.Start.Character == 0 &&
		firstEntry.LocationRange.End.Character == 16 &&
		firstEntry.Key.Value.Value == "HostName" &&
		firstEntry.Key.LocationRange.Start.Character == 0 &&
		firstEntry.Key.LocationRange.End.Character == 8 &&
		firstEntry.OptionValue.Value.Value == "1.2.3.4" &&
		firstEntry.OptionValue.LocationRange.Start.Character == 9 &&
		firstEntry.OptionValue.LocationRange.End.Character == 16) {
		t.Errorf("Expected first entry to be HostName 1.2.3.4, but got: %v", firstEntry)
	}

	rawSecondEntry, _ := p.Options.Get(uint32(1))
	secondEntry := rawSecondEntry.(*SSHOption)

	if !(secondEntry.Value.Value == "User root" &&
		secondEntry.LocationRange.Start.Line == 1 &&
		secondEntry.LocationRange.End.Line == 1 &&
		secondEntry.LocationRange.Start.Character == 0 &&
		secondEntry.LocationRange.End.Character == 9 &&
		secondEntry.Key.Value.Value == "User" &&
		secondEntry.Key.LocationRange.Start.Character == 0 &&
		secondEntry.Key.LocationRange.End.Character == 4 &&
		secondEntry.OptionValue.Value.Value == "root" &&
		secondEntry.OptionValue.LocationRange.Start.Character == 5 &&
		secondEntry.OptionValue.LocationRange.End.Character == 9) {
		t.Errorf("Expected second entry to be User root, but got: %v", secondEntry)
	}
}

func TestMatchSimpleBlock(
	t *testing.T,
) {
	input := utils.Dedent(`
Hostname 1.2.3.4

Match originalhost "192.168.0.1"
	User root
`)
	p := NewSSHConfig()
	errors := p.Parse(input)

	if len(errors) != 0 {
		t.Fatalf("Expected no errors, got %v", errors)
	}

	if !(p.Options.Size() == 2 &&
		len(utils.KeysOfMap(p.CommentLines)) == 0) {
		t.Errorf("Expected 2 option and no comment lines, but got: %v entries, %v comment lines", p.Options.Size(), len(p.CommentLines))
	}

	rawFirstEntry, _ := p.Options.Get(uint32(0))
	firstEntry := rawFirstEntry.(*SSHOption)

	if !(firstEntry.Value.Value == "Hostname 1.2.3.4" &&
		firstEntry.LocationRange.Start.Line == 0 &&
		firstEntry.LocationRange.End.Line == 0 &&
		firstEntry.LocationRange.Start.Character == 0 &&
		firstEntry.LocationRange.End.Character == 16 &&
		firstEntry.Key.Value.Value == "Hostname" &&
		firstEntry.Key.LocationRange.Start.Character == 0 &&
		firstEntry.Key.LocationRange.End.Character == 8 &&
		firstEntry.OptionValue.Value.Value == "1.2.3.4" &&
		firstEntry.OptionValue.LocationRange.Start.Character == 9 &&
		firstEntry.OptionValue.LocationRange.End.Character == 16) {
		t.Errorf("Expected first entry to be Hostname 1.2.3.4, but got: %v", firstEntry)
	}

	rawSecondEntry, _ := p.Options.Get(uint32(2))
	secondEntry := rawSecondEntry.(*SSHMatchBlock)

	if !(secondEntry.Options.Size() == 1 &&
		secondEntry.LocationRange.Start.Line == 2 &&
		secondEntry.LocationRange.End.Line == 3 &&
		secondEntry.LocationRange.Start.Character == 0 &&
		secondEntry.LocationRange.End.Character == 10 &&
		secondEntry.MatchOption.OptionValue.Value.Raw == "originalhost \"192.168.0.1\"" &&
		secondEntry.MatchOption.OptionValue.LocationRange.Start.Character == 6 &&
		secondEntry.MatchOption.OptionValue.LocationRange.End.Character == 32) {
		t.Errorf("Expected second entry to be Match originalhost \"192.168.0.1\", but got: %v; options amount: %d", secondEntry, secondEntry.Options.Size())
	}

	rawThirdEntry, _ := secondEntry.Options.Get(uint32(3))
	thirdEntry := rawThirdEntry.(*SSHOption)
	if !(thirdEntry.Value.Raw == "\tUser root" &&
		thirdEntry.LocationRange.Start.Line == 3 &&
		thirdEntry.LocationRange.End.Line == 3 &&
		thirdEntry.LocationRange.Start.Character == 0 &&
		thirdEntry.LocationRange.End.Character == 10 &&
		thirdEntry.Key.Value.Value == "User" &&
		thirdEntry.Key.LocationRange.Start.Character == 1 &&
		thirdEntry.Key.LocationRange.End.Character == 5 &&
		thirdEntry.OptionValue.Value.Value == "root" &&
		thirdEntry.OptionValue.LocationRange.Start.Character == 6 &&
		thirdEntry.OptionValue.LocationRange.End.Character == 10) {
		t.Errorf("Expected third entry to be User root, but got: %v", thirdEntry)
	}
}
