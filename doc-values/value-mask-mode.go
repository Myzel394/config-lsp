package docvalues

import (
	"fmt"
	"regexp"

	protocol "github.com/tliron/glsp/protocol_3_16"
)

type MaskModeInvalidError struct{}

func (e MaskModeInvalidError) Error() string {
	return "This mask is not valid. It must be a 4-digit octal number."
}

type MaskModeValue struct{}

func (v MaskModeValue) GetTypeDescription() []string {
	return []string{
		"File mode mask",
		"4-digit octal number",
		"Example: 0000",
		"1st digit: setuid, setgid, sticky",
		"2nd digit: user permissions",
		"3rd digit: group permissions",
		"4th digit: other permissions",
	}
}

var maskModePattern = regexp.MustCompile("^[0-7]{4}$")

func (v MaskModeValue) CheckIsValid(value string) []*InvalidValue {
	if !maskModePattern.MatchString(value) {
		return []*InvalidValue{{
			Err:   MaskModeInvalidError{},
			Start: 0,
			End:   uint32(len(value)),
		}}
	}

	return []*InvalidValue{}
}

func (v MaskModeValue) FetchCompletions(line string, cursor uint32) []protocol.CompletionItem {
	kind := protocol.CompletionItemKindValue

	perm0644 := "0644"
	perm0755 := "0755"
	perm0777 := "0777"
	perm0400 := "0400"
	perm0600 := "0600"
	perm0666 := "0666"
	perm0700 := "0700"
	perm0444 := "0444"
	perm0111 := "0111"
	perm0555 := "0555"

	return []protocol.CompletionItem{
		{
			Label:         "0000",
			Documentation: "No permissions",
			Kind:          &kind,
		},
		{
			Label:         "0644 (-rw-r--r--)",
			InsertText:    &perm0644,
			Documentation: "Read/write for owner, read for others",
			Kind:          &kind,
		},
		{
			Label:         "0755 (-rwxr-xr-x)",
			InsertText:    &perm0755,
			Documentation: "Read/write/execute for owner, read/execute for others",
			Kind:          &kind,
		},
		{
			Label:         "0777 (-rwxrwxrwx)",
			InsertText:    &perm0777,
			Documentation: "Read/write/execute for all",
			Kind:          &kind,
		},
		{
			Label:         "0400 (-r--------)",
			InsertText:    &perm0400,
			Documentation: "Read for owner",
			Kind:          &kind,
		},
		{
			Label:         "0600 (-rw-------)",
			InsertText:    &perm0600,
			Documentation: "Read/write for owner",
			Kind:          &kind,
		},
		{
			Label:         "0666 (-rw-rw-rw-)",
			InsertText:    &perm0666,
			Documentation: "Read/write for all",
			Kind:          &kind,
		},
		{
			Label:         "0700 (-rwx------)",
			InsertText:    &perm0700,
			Documentation: "Read/write/execute for owner",
			Kind:          &kind,
		},
		{
			Label:         "0444 (-r--r--r--)",
			InsertText:    &perm0444,
			Documentation: "Read for all",
			Kind:          &kind,
		},
		{
			Label:         "0111 (-x--x--x)",
			InsertText:    &perm0111,
			Documentation: "Execute for all",
			Kind:          &kind,
		},
		{
			Label:         "0555 (-r-xr-xr-x)",
			InsertText:    &perm0555,
			Documentation: "Read/execute for all",
			Kind:          &kind,
		},
	}
}

func getMaskRepresentation(digit uint8) string {
	switch digit {
	case 0:
		return "---"
	case 1:
		return "--x"
	case 2:
		return "-w-"
	case 3:
		return "-wx"
	case 4:
		return "r--"
	case 5:
		return "r-x"
	case 6:
		return "rw-"
	case 7:
		return "rwx"
	}

	return ""
}

func (v MaskModeValue) FetchHoverInfo(line string, cursor uint32) []string {
	if !maskModePattern.MatchString(line) {
		return []string{}
	}

	mask := line

	firstDigit := uint8(mask[1] - 48)
	secondDigit := uint8(mask[2] - 48)
	thridDigit := uint8(mask[3] - 48)

	representation := getMaskRepresentation(firstDigit) + getMaskRepresentation(secondDigit) + getMaskRepresentation(thridDigit)

	return []string{
		fmt.Sprintf("%s (%s)", mask, representation),
	}
}
