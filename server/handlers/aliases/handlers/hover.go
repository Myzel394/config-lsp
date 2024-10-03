package handlers

import (
	"config-lsp/handlers/aliases/ast"
	"config-lsp/handlers/aliases/fields"
	"config-lsp/handlers/aliases/indexes"
	"config-lsp/utils"
	"fmt"
	"strings"
)

// Get hover information for an alias entry
// Expects `entry` to contain at least a key
func GetAliasHoverInfo(
	i indexes.AliasesIndexes,
	entry ast.AliasEntry,
) string {
	header := []string{
		fmt.Sprintf("Emails targeted for `%s` will be passed to:", entry.Key.Value),
		"",
	}

	var forwards []string

	if entry.Values == nil {
		forwards = []string{
			"No forwards configured",
		}
	} else {
		if len(entry.Values.Values) == 1 {
			forwards = []string{
				GetAliasValueHoverInfo(
					i,
					entry.Values.Values[0],
				),
			}
		} else {
			forwards = utils.Map(
				entry.Values.Values,
				func(value ast.AliasValueInterface) string {
					return fmt.Sprintf(
						"* %s",
						GetAliasValueHoverInfo(
							i,
							value,
						),
					)
				},
			)
		}
	}

	content := append(header, forwards...)
	return strings.Join(
		content,
		"\n",
	)
}

func GetAliasValueHoverInfo(
	i indexes.AliasesIndexes,
	value ast.AliasValueInterface,
) string {
	switch value.(type) {
	case ast.AliasValueUser:
		return fmt.Sprintf("User: **%s**", value.GetAliasValue().Value)
	case ast.AliasValueEmail:
		return fmt.Sprintf("Email: **%s**", value.GetAliasValue().Value)
	case ast.AliasValueInclude:
		includeValue := value.(ast.AliasValueInclude)
		return fmt.Sprintf("Included file: `%s`", string(includeValue.Path.Path))
	case ast.AliasValueFile:
		fileValue := value.(ast.AliasValueFile)
		return fmt.Sprintf("File: Email will be written to `%s`", string(fileValue.Path))
	case ast.AliasValueCommand:
		commandValue := value.(ast.AliasValueCommand)
		return fmt.Sprintf("Command: Will be passed as stdin to `%s`", commandValue.Command)
	case ast.AliasValueError:
		errorValue := value.(ast.AliasValueError)

		if errorValue.Code == nil || errorValue.Message == nil {
			return "Error: An error will show up"
		}

		return fmt.Sprintf(
			"Error: An error will show up; code: **%s** (%s), message: '%s'",
			errorValue.Code.Value,
			getErrorCodeInfo(errorValue.Code.ErrorCodeAsInt()),
			errorValue.Message.Value,
		)
	}

	panic("Unknown value type")
}

func GetAliasValueTypeInfo(
	value ast.AliasValueInterface,
) []string {
	switch value.(type) {
	case ast.AliasValueUser:
		return []string{
			"### User",
			fields.UserDeclaration,
			"",
			fields.UserField.Documentation,
		}
	case ast.AliasValueEmail:
		return []string{
			"### Email",
			fields.EmailDeclaration,
			"",
			fields.EmailField.Documentation,
		}
	case ast.AliasValueInclude:
		return []string{
			"### Include",
			fields.IncludeDeclaration,
			"",
			fields.IncludeField.Documentation,
		}
	case ast.AliasValueFile:
		return []string{
			"### File",
			fields.PathDeclaration,
			"",
			fields.PathField.Documentation,
		}
	case ast.AliasValueCommand:
		return []string{
			"### Command",
			fields.CommandDeclaration,
			"",
			fields.CommandField.Documentation,
		}
	case ast.AliasValueError:
		return []string{
			"### Error",
			fields.ErrorDeclaration,
			"",
			fields.ErrorMessageField.Documentation,
		}
	}

	panic("Unknown value type")
}

func getErrorCodeInfo(
	code uint16,
) string {
	if code >= 400 && code <= 499 {
		return "4XX: TempFail"
	}

	if code >= 500 && code <= 599 {
		return "5XX: PermFail"
	}

	return "Unknown code"
}
