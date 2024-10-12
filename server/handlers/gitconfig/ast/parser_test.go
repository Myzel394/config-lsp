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

	rawFirstOption, _ := section.Entries.Get(uint32(1))
	firstOption := rawFirstOption.(*GitEntry)
	if !(firstOption.Key.Value.Value == "repositoryformatversion" && firstOption.Value.Value == "0") {
		t.Errorf("Expected repositoryformatversion, got %s", firstOption.Key.Value.Value)
	}

	rawSecondOption, _ := section.Entries.Get(uint32(2))
	secondOption := rawSecondOption.(*GitEntry)
	if !(secondOption.Key.Value.Value == "filemode" && secondOption.Value.Value == "true") {
		t.Errorf("Expected filemode, got %s", secondOption.Key.Value.Value)
	}

	rawThirdOption, _ := section.Entries.Get(uint32(3))
	thirdOption := rawThirdOption.(*GitEntry)
	if !(thirdOption.Key.Value.Value == "bare" && thirdOption.Value.Value == "false") {
		t.Errorf("Expected bare, got %s", thirdOption.Key.Value.Value)
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

	if !(section.Entries.Size() == 3) {
		t.Errorf("Expected 3 entries, got %d", section.Entries.Size())
	}

	rawFirstOption, _ := section.Entries.Get(uint32(1))
	firstOption := rawFirstOption.(*GitEntry)
	if !(firstOption.Key.Value.Value == "repositoryformatversion" && firstOption.Value.Value == "0" && firstOption.Start.Line == 1 && firstOption.End.Line == 1 && firstOption.Start.Character == 4 && firstOption.End.Character == 31) {
		t.Errorf("Expected 0, got %s", firstOption)
	}

	_, found := section.Entries.Get(uint32(3))
	if found {
		t.Errorf("Expected no entry at line 3")
	}

	section = config.Sections[1]

	if !(section.Title.Title == `remote "origin"`) {
		t.Errorf("Expected remote \"origin\", got %s", section.Title.Title)
	}

	if !(section.Entries.Size() == 2) {
		t.Errorf("Expected 2 entries, got %d", section.Entries.Size())
	}

	section = config.Sections[2]

	if !(section.Title.Title == "alias") {
		t.Errorf("Expected alias, got %s", section.Title.Title)
	}

	rawSecondOption, _ := section.Entries.Get(uint32(13))
	secondOption := rawSecondOption.(*GitEntry)
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
