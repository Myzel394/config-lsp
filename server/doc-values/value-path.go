package docvalues

import (
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

	fileRequired := (v.RequiredType & PathTypeFile) == PathTypeFile
	directoryRequired := (v.RequiredType & PathTypeDirectory) == PathTypeDirectory

	isValid := true

	// If file is expected
	if fileRequired {
		// and exists
		isValid = isValid && utils.IsPathFile(value)
		// file not expected
	} else {
		// and should not exist
		isValid = isValid && !utils.IsPathFile(value)
	}

	// if directory
	if directoryRequired {
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

	if fileRequired && directoryRequired {
		return []*InvalidValue{{
			Err:   errors.New("This must be either a file or a directory"),
			Start: 0,
			End:   uint32(len(value)),
		}}
	}
	if fileRequired {
		return []*InvalidValue{{
			Err:   errors.New("This must be a file"),
			Start: 0,
			End:   uint32(len(value)),
		}}
	}
	if directoryRequired {
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

func (v PathValue) DeprecatedFetchCompletions(line string, cursor uint32) []protocol.CompletionItem {
	return []protocol.CompletionItem{}
}

func (v PathValue) DeprecatedFetchHoverInfo(line string, cursor uint32) []string {
	return []string{}
}
