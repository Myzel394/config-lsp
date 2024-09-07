package analyzer

import (
	"config-lsp/common"
	"config-lsp/handlers/aliases/ast"
	"config-lsp/handlers/aliases/indexes"
	"config-lsp/utils"
	"errors"
	"fmt"
)

var valueHandlerMap = map[string]func(
	rawValue []ast.AliasValueInterface,
) []common.LSPError{
	"AliasValueUser":    analyzeValueUser,
	"AliasValueEmail":   analyzeValueEmail,
	"AliasValueCommand": analyzeValueCommand,
	"AliasValueFile":    analyzeValueFile,
	"AliasValueInclude": analyzeValueInclude,
	"AliasValueError":   analyzeValueError,
}

func analyzeContainsNoDoubleValues(
	p ast.AliasesParser,
) []common.LSPError {
	errors := make([]common.LSPError, 0)

	it := p.Aliases.Iterator()

	for it.Next() {
		entry := it.Value().(*ast.AliasEntry)

		valuesPerType := utils.Group(
			entry.Values.Values,
			func(entry ast.AliasValueInterface) string {
				return entry.GetStructName()
			},
		)

		for valueType, values := range valuesPerType {
			handler := valueHandlerMap[valueType]

			newErrors := handler(values)
			errors = append(errors, newErrors...)
		}
	}

	return errors
}

func analyzeValueUser(
	rawValues []ast.AliasValueInterface,
) []common.LSPError {
	users := make(map[string]struct{})
	errs := make([]common.LSPError, 0)

	// Simple double value check
	for _, rawValue := range rawValues {
		value := rawValue.(ast.AliasValueUser)
		key := indexes.NormalizeKey(value.Value)

		if _, found := users[key]; found {
			errs = append(errs, common.LSPError{
				Range: value.Location,
				Err:   errors.New(fmt.Sprintf("User '%s' is defined multiple times", key)),
			})
		} else {
			users[key] = struct{}{}
		}

	}

	return errs
}

func analyzeValueEmail(
	rawValues []ast.AliasValueInterface,
) []common.LSPError {
	emails := make(map[string]struct{})
	errs := make([]common.LSPError, 0)

	for _, rawValue := range rawValues {
		value := rawValue.(ast.AliasValueEmail)

		// Simple double value check
		if _, found := emails[value.Value]; found {
			errs = append(errs, common.LSPError{
				Range: value.Location,
				Err:   errors.New(fmt.Sprintf("Email '%s' is defined multiple times", value.Value)),
			})
		} else {
			emails[value.Value] = struct{}{}
		}
	}

	return errs
}

func analyzeValueCommand(
	rawValues []ast.AliasValueInterface,
) []common.LSPError {
	commands := make(map[string]struct{})
	errs := make([]common.LSPError, 0)

	for _, rawValue := range rawValues {
		value := rawValue.(ast.AliasValueCommand)
		command := value.Command

		// Simple double value check
		if _, found := commands[command]; found {
			errs = append(errs, common.LSPError{
				Range: value.Location,
				Err:   errors.New(fmt.Sprintf("Command '%s' is defined multiple times", command)),
			})
		} else {
			commands[command] = struct{}{}
		}
	}

	return errs
}

func analyzeValueFile(
	rawValues []ast.AliasValueInterface,
) []common.LSPError {
	files := make(map[string]struct{})
	errs := make([]common.LSPError, 0)

	for _, rawValue := range rawValues {
		value := rawValue.(ast.AliasValueFile)
		path := string(value.Path)

		// Simple double value check
		if _, found := files[path]; found {
			errs = append(errs, common.LSPError{
				Range: value.Location,
				Err:   errors.New(fmt.Sprintf("File '%s' is defined multiple times", path)),
			})
		} else {
			files[path] = struct{}{}
		}
	}

	return errs
}

func analyzeValueInclude(
	rawValues []ast.AliasValueInterface,
) []common.LSPError {
	files := make(map[string]struct{})
	errs := make([]common.LSPError, 0)

	for _, rawValue := range rawValues {
		value := rawValue.(ast.AliasValueInclude)
		path := string(value.Path.Path)

		// Simple double value check
		if _, found := files[path]; found {
			errs = append(errs, common.LSPError{
				Range: value.Location,
				Err:   errors.New(fmt.Sprintf("Inclusion '%s' is included multiple times", path)),
			})
		} else {
			files[path] = struct{}{}
		}
	}

	return errs
}

func analyzeValueError(
	rawValues []ast.AliasValueInterface,
) []common.LSPError {
	codes := make(map[uint16]struct{})
	errs := make([]common.LSPError, 0)

	for _, rawValue := range rawValues {
		value := rawValue.(ast.AliasValueError)
		code := value.Code.ErrorCodeAsInt()

		// Simple double value check
		if _, found := codes[code]; found {
			errs = append(errs, common.LSPError{
				Range: value.Location,
				Err:   errors.New(fmt.Sprintf("Error code '%d' is defined multiple times", code)),
			})
		} else {
			codes[code] = struct{}{}
		}
	}

	return errs
}
