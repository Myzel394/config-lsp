package docvalues

import (
	"config-lsp/common"
	"config-lsp/utils"
	"errors"
	"strings"

	protocol "github.com/tliron/glsp/protocol_3_16"
)

type PathType uint8

const (
	PathTypeFile      PathType = 1
	PathTypeDirectory PathType = 2
)

type PathValue struct {
	IsOptional   bool
	RequiredType PathType
}

func (v PathValue) GetTypeDescription() []string {
	hints := make([]string, 0)

	switch v.RequiredType {
	case PathTypeFile:
		hints = append(hints, "File")
	case PathTypeDirectory:
		hints = append(hints, "Directory")
	}

	if v.IsOptional {
		hints = append(hints, "Optional")
	}

	return []string{strings.Join(hints, ", ")}
}

func (v PathValue) DeprecatedCheckIsValid(value string) []*InvalidValue {
	if !utils.DoesPathExist(value) {
		if v.IsOptional {
			return nil
		} else {
			return []*InvalidValue{{
				Err:   errors.New("This path does not exist"),
				Start: 0,
				End:   uint32(len(value)),
			}}
		}
	}

	fileExpected := (v.RequiredType & PathTypeFile) == PathTypeFile
	directoryExpected := (v.RequiredType & PathTypeDirectory) == PathTypeDirectory

	isValid := true

	// If file is expected
	if fileExpected {
		// and exists
		isValid = isValid && utils.IsPathFile(value)
		// file not expected
	} else {
		// and should not exist
		isValid = isValid && !utils.IsPathFile(value)
	}

	// if directory
	if directoryExpected {
		// and exists
		isValid = isValid && utils.IsPathDirectory(value)
		// directory not expected
	} else {
		// and should not exist
		isValid = isValid && !utils.IsPathDirectory(value)
	}

	if isValid {
		return nil
	}

	if fileExpected && directoryExpected {
		return []*InvalidValue{{
			Err:   errors.New("This must be either a file or a directory"),
			Start: 0,
			End:   uint32(len(value)),
		}}
	}
	if fileExpected {
		return []*InvalidValue{{
			Err:   errors.New("This must be a file"),
			Start: 0,
			End:   uint32(len(value)),
		}}
	}
	if directoryExpected {
		return []*InvalidValue{{
			Err:   errors.New("This must be a directory"),
			Start: 0,
			End:   uint32(len(value)),
		}}
	}

	return []*InvalidValue{{
		Err:   errors.New("This path is invalid"),
		Start: 0,
		End:   uint32(len(value)),
	}}
}

func (v PathValue) FetchCompletions(value string, cursor common.CursorPosition) []protocol.CompletionItem {
	return []protocol.CompletionItem{}
}

func (v PathValue) DeprecatedFetchHoverInfo(line string, cursor uint32) []string {
	return []string{}
}
