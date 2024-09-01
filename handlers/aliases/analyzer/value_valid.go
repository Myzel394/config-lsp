package analyzer

import (
	"config-lsp/common"
	"config-lsp/handlers/aliases"
	"config-lsp/handlers/aliases/ast"
	"config-lsp/handlers/aliases/fetchers"
	"config-lsp/handlers/aliases/indexes"
	"fmt"
	"net/mail"
	"path"
	"strconv"

	ers "errors"
)

func analyzeValuesAreValid(
	d *aliases.AliasesDocument,
) []common.LSPError {
	errors := make([]common.LSPError, 0)

	it := d.Parser.Aliases.Iterator()

	for it.Next() {
		entry := it.Value().(*ast.AliasEntry)

		if entry.Key == nil {
			errors = append(errors, common.LSPError{
				Range: entry.Location,
				Err:   ers.New("An alias is required"),
			})

			continue
		}

		if entry.Separator == nil {
			errors = append(errors, common.LSPError{
				Range: entry.Location,
				Err:   ers.New("A ':' is required as a separator"),
			})

			continue
		}

		if entry.Values == nil || len(entry.Values.Values) == 0 {
			errors = append(errors, common.LSPError{
				Range: entry.Location,
				Err:   ers.New("A value is required"),
			})

			continue
		}

		for _, value := range entry.Values.Values {
			newErrors := checkValue(d.Indexes, value)

			errors = append(errors, newErrors...)
		}
	}

	return errors
}

func checkValue(
	i *indexes.AliasesIndexes,
	value ast.AliasValueInterface,
) []common.LSPError {
	switch value.(type) {
	case ast.AliasValueUser:
		aliasValue := value.(ast.AliasValueUser)

		users := fetchers.GetAvailableUserValues(i)

		if _, found := users[aliasValue.Value]; !found {
			return []common.LSPError{{
				Range: aliasValue.Location,
				Err:   ers.New(fmt.Sprintf("User '%s' not found", aliasValue.Value)),
			}}
		}
	case ast.AliasValueEmail:
		emailValue := value.(ast.AliasValueEmail)

		if _, error := mail.ParseAddress(emailValue.Value); error != nil {
			return []common.LSPError{{
				Range: emailValue.Location,
				Err:   ers.New(fmt.Sprintf("This does not seem to be a valid email: %s", error.Error())),
			}}
		}
	case ast.AliasValueFile:
		fileValue := value.(ast.AliasValueFile)

		// I'm not sure if the path really needs to be absolute
		// The docs say:
		// "Append messages to file, specified by its absolute pathname."
		//
		if !path.IsAbs(fileValue.Value) {
			return []common.LSPError{{
				Range: fileValue.Location,
				Err:   ers.New("This path must be absolute"),
			}}
		}
	case ast.AliasValueError:
		errorValue := value.(ast.AliasValueError)

		if errorValue.Code == nil {
			return []common.LSPError{{
				Range: errorValue.Location,
				Err:   ers.New("An error code in the form of 4XX or 5XX is required"),
			}}
		}

		errorCode, err := strconv.Atoi(errorValue.Code.Value)

		if err != nil || (errorCode < 400 || errorCode > 599) {
			return []common.LSPError{{
				Range: errorValue.Code.Location,
				Err:   ers.New("This error code is invalid. It must be in the form of 4XX or 5XX"),
			}}
		}

		if errorValue.Message == nil || errorValue.Message.Value == "" {
			return []common.LSPError{{
				Range: errorValue.Location,
				Err:   ers.New("An error message is required"),
			}}
		}
	}
	return nil
}
