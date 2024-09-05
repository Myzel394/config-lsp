package handlers

import (
	"config-lsp/handlers/aliases/ast"
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
			"`user`",
			"",
			"A user on the host machine. The user must have a valid entry in the passwd(5) database file.",
		}
	case ast.AliasValueEmail:
		return []string{
			"### Email",
			"`user-part@domain-part`",
			"",
			"An email address in RFC 5322 format. If an address extension is appended to the user-part, it is first compared for an exact match.  It is then stripped so that an address such as user+ext@example.com will only use the part that precedes ‘+’ as a key.",
		}
	case ast.AliasValueInclude:
		return []string{
			"### Include",
			"`include:/path/to/file`",
			"",
			"Include any definitions in file as alias entries. The format of the file is identical to this one.",
		}
	case ast.AliasValueFile:
		return []string{
			"### File",
			"`/path/to/file`",
			"",
			"Append messages to file, specified by its absolute pathname.",
		}
	case ast.AliasValueCommand:
		return []string{
			"### Command",
			"`|command`",
			"",
			"Pipe the message to command on its standard input. The command is run under the privileges of the daemon's unprivileged account.",
		}
	case ast.AliasValueError:
		return []string{
			"### Error",
			"`error:code message`",
			"",
			"A status code and message to return. The code must be 3 digits, starting 4XX (TempFail) or 5XX (PermFail). The message must be present and can be freely chosen.",
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
