package ast

import (
	"config-lsp/utils"
	"testing"
)

func TestValidOneSectionExample(t *testing.T) {
	input := utils.Dedent(`
[core]
	repositoryformatversion = 0
	filemode = true
	bare = false
`)

	config := NewGitConfig()
	errors := config.Parse(input)

	if len(errors) > 0 {
		t.Fatalf("Expected no errors, got %v", errors)
	}

	if len(config.Sections) != 1 {
		t.Fatalf("Expected 1 section, got %d", len(config.Sections))
	}

	section := config.Sections[0]

	firstOption := section.Entries[0]
	if !(firstOption.Key.Value.Value == "repositoryformatversion" && firstOption.Value.Value == "0") {
		t.Errorf("Expected repositoryformatversion, got %s", firstOption.Key.Value.Value)
	}

	secondOption := section.Entries[1]
	if !(secondOption.Key.Value.Value == "filemode" && secondOption.Value.Value == "true") {
		t.Errorf("Expected filemode, got %s", secondOption.Key.Value.Value)
	}

	thirdOption := section.Entries[2]
	if !(thirdOption.Key.Value.Value == "bare" && thirdOption.Value.Value == "false") {
		t.Errorf("Expected bare, got %s", thirdOption.Key.Value.Value)
	}

	foundSection, foundEntry := config.FindOption(1)

	if !(foundSection == section && foundEntry == firstOption) {
		t.Errorf("Expected first option, got %s", foundEntry.Key.Value.Value)
	}

	foundSection, foundEntry = config.FindOption(0)

	if !(foundSection == section && foundEntry == nil) {
		t.Errorf("Expected nil, got %s", foundEntry)
	}
}

func TestComplexExample(t *testing.T) {
	input := utils.Dedent(`
[core]
    repositoryformatversion = 0 # This contains 4 spaces instead of a tab at the beginning
	filemode = true # Hello World

	# Hello World
	bare = false ; Trailing comment

# Test
[remote "origin"]
	url = https://github.com/github/repo.git
	fetch = +refs/heads/*:refs/remotes/origin/*

[alias]
	ours = "!f() { git checkout --ours $@ && git add $@; }; f" ; This is a comment
; Hello
`)

	config := NewGitConfig()
	errors := config.Parse(input)

	if len(errors) > 0 {
		t.Fatalf("Expected no errors, got %v", errors)
	}

	if len(config.Sections) != 3 {
		t.Fatalf("Expected 3 sections, got %d", len(config.Sections))
	}

	section := config.Sections[0]

	if !(section.Title.Title == "core") {
		t.Errorf("Expected core, got %s", section.Title.Title)
	}

	if !(len(section.Entries) == 3) {
		t.Errorf("Expected 3 entries, got %d", len(section.Entries))
	}

	firstOption := section.Entries[0]
	if !(firstOption.Key.Value.Value == "repositoryformatversion" && firstOption.Value.Value == "0" && firstOption.Start.Line == 1 && firstOption.End.Line == 1 && firstOption.Start.Character == 4 && firstOption.End.Character == 31) {
		t.Errorf("Expected 0, got %s", firstOption)
	}

	section = config.Sections[1]

	if !(section.Title.Title == `remote "origin"`) {
		t.Errorf("Expected remote \"origin\", got %s", section.Title.Title)
	}

	if !(len(section.Entries) == 2) {
		t.Errorf("Expected 2 entries, got %d", len(section.Entries))
	}

	section = config.Sections[2]

	if !(section.Title.Title == "alias") {
		t.Errorf("Expected alias, got %s", section.Title.Title)
	}

	secondOption := section.Entries[0]
	if !(secondOption.Key.Value.Value == "ours" && secondOption.Value.Value == "!f() { git checkout --ours $@ && git add $@; }; f" && secondOption.Start.Character == 1 && secondOption.End.Character == 59) {
		t.Errorf("Expected ours, got %s", secondOption.Key.Value.Value)
	}
}

func TestMissingSectionHeader(t *testing.T) {
	input := utils.Dedent(`
repositoryformatversion = 0
filemode = true
bare = false

[core]
	bare = false
`)

	config := NewGitConfig()
	errors := config.Parse(input)

	if len(errors) != 1 {
		t.Fatalf("Expected 1 error, got %d", len(errors))
	}
}

func TestLeadingLine(t *testing.T) {
	input := utils.Dedent(`
[core]
	command = git \
		commit \
		-m "Hello World"
`)

	config := NewGitConfig()
	errors := config.Parse(input)

	if len(errors) > 0 {
		t.Fatalf("Expected no errors, got %v", errors)
	}

	if len(config.Sections) != 1 {
		t.Fatalf("Expected 1 section, got %d", len(config.Sections))
	}

	section := config.Sections[0]

	if !(section.Title.Title == "core") {
		t.Errorf("Expected core, got %s", section.Title.Title)
	}

	firstOption := section.Entries[0]
	if !(firstOption.Key.Value.Value == "command" && firstOption.Value.Value == "git commit -m Hello World") {
		t.Errorf("Expected command, got %s", firstOption.Key.Value.Value)
	}

	if !(firstOption.Value.Raw.Parts[0].Text == "git " && firstOption.Value.Raw.Parts[1].Text == "\t\tcommit " && firstOption.Value.Raw.Parts[2].Text == "\t\t-m \"Hello World\"") {
		t.Errorf("Expected command, got %s", firstOption.Value.Raw.Parts)
	}
}