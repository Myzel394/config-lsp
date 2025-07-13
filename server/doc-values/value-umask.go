package docvalues

import (
	"config-lsp/common"
	"regexp"

	protocol "github.com/tliron/glsp/protocol_3_16"
)

type UmaskInvalidError struct{}

func (e UmaskInvalidError) Error() string {
	return "This mask is not valid. It must be a 4-digit octal number."
}

type UmaskValue struct{}

func (v UmaskValue) GetTypeDescription() []string {
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

var umaskPattern = regexp.MustCompile("^[0-7]{4}$")

func (v UmaskValue) DeprecatedCheckIsValid(value string) []*InvalidValue {
	if !umaskPattern.MatchString(value) {
		return []*InvalidValue{{
			Err:   UmaskInvalidError{},
			Start: 0,
			End:   uint32(len(value)),
		}}
	}

	return []*InvalidValue{}
}

func (v UmaskValue) FetchCompletions(value string, cursor common.CursorPosition) []protocol.CompletionItem {
	kind := protocol.CompletionItemKindValue

	return []protocol.CompletionItem{
		{
			Label:         "0000",
			Documentation: "Files: rwxrwxrwx Directories: rwxrwxrwx",
			Kind:          &kind,
		},
		{
			Label:         "0022",
			Documentation: "Files: rwxr-xr-x Directories: rwxr-xr-x",
			Kind:          &kind,
		},
		{
			Label:         "0077",
			Documentation: "Files: rwx------ Directories: rwx------",
			Kind:          &kind,
		},
		{
			Label:         "0177",
			Documentation: "Files: rw------- Directories: rw-------",
			Kind:          &kind,
		},
		{
			Label:         "0277",
			Documentation: "Files: r-x------ Directories: r-x------",
		},
		{
			Label:         "0027",
			Documentation: "Files: rwxr-x--- Directories: rwxr-x---",
			Kind:          &kind,
		},
		{
			Label:         "0070",
			Documentation: "Files: rwx---rwx Directories: rwx---rwx",
			Kind:          &kind,
		},
		{
			Label:         "0222",
			Documentation: "Files: r--r--r-- Directories: r--r--r--",
			Kind:          &kind,
		},
		{
			Label:         "0333",
			Documentation: "Files: r--r--r-- Directories: r--r--r--",
			Kind:          &kind,
		},
		{
			Label:         "0444",
			Documentation: "Files: r--r--r-- Directories: r--r--r--",
			Kind:          &kind,
		},
		{
			Label:         "0555",
			Documentation: "Files: r-xr-xr-x Directories: r-xr-xr-x",
			Kind:          &kind,
		},
		{
			Label:         "0666",
			Documentation: "Files: rw-rw-rw- Directories: rw-rw-rw-",
			Kind:          &kind,
		},
		{
			Label:         "0777",
			Documentation: "Files: rwxrwxrwx Directories: rwxrwxrwx",
			Kind:          &kind,
		},
	}
}

func (v UmaskValue) DeprecatedFetchHoverInfo(line string, cursor uint32) []string {
	return []string{}
}
