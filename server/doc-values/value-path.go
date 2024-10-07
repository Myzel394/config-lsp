package docvalues

import (
	"config-lsp/utils"
	protocol "github.com/tliron/glsp/protocol_3_16"
	"strings"
)

type PathDoesNotExistError struct{}

func (e PathDoesNotExistError) Error() string {
	return "This path does not exist"
}

type PathInvalidError struct{}

func (e PathInvalidError) Error() string {
	return "This path is invalid"
}

type PathType uint8

const (
	PathTypeExistenceOptional PathType = 0
	PathTypeFile              PathType = 1
	PathTypeDirectory         PathType = 2
)

type PathValue struct {
	RequiredType PathType
}

func (v PathValue) GetTypeDescription() []string {
	hints := make([]string, 0)

	switch v.RequiredType {
	case PathTypeExistenceOptional:
		hints = append(hints, "Optional")
		break
	case PathTypeFile:
		hints = append(hints, "File")
	case PathTypeDirectory:
		hints = append(hints, "Directory")
	}

	return []string{strings.Join(hints, ", ")}
}

func (v PathValue) DeprecatedCheckIsValid(value string) []*InvalidValue {
	if !utils.DoesPathExist(value) {
		return []*InvalidValue{{
			Err:   PathDoesNotExistError{},
			Start: 0,
			End:   uint32(len(value)),
		}}
	}

	isValid := false

	if (v.RequiredType & PathTypeFile) == PathTypeFile {
		isValid = isValid && utils.IsPathFile(value)
	}

	if (v.RequiredType & PathTypeDirectory) == PathTypeDirectory {
		isValid = isValid && utils.IsPathDirectory(value)
	}

	if isValid {
		return nil
	}

	return []*InvalidValue{{
		Err:   PathInvalidError{},
		Start: 0,
		End:   uint32(len(value)),
	},
	}
}

func (v PathValue) DeprecatedFetchCompletions(line string, cursor uint32) []protocol.CompletionItem {
	return []protocol.CompletionItem{}
}

func (v PathValue) DeprecatedFetchHoverInfo(line string, cursor uint32) []string {
	return []string{}
}
