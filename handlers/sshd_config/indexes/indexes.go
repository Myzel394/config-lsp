package indexes

import (
	"config-lsp/common"
	"config-lsp/handlers/sshd_config/ast"
	"errors"
	"fmt"
)

var allowedDoubleOptions = map[string]struct{}{
	"AllowGroups":   {},
	"AllowUsers":    {},
	"DenyGroups":    {},
	"DenyUsers":     {},
	"ListenAddress": {},
	"Match":         {},
	"Port":          {},
}

type SSHIndexKey struct {
	Option     string
	MatchBlock *ast.SSHMatchBlock
}

type SSHIndexAllOption struct {
	MatchBlock *ast.SSHMatchBlock
	Option   *ast.SSHOption
}

type SSHIndexes struct {
	// Contains a map of `Option name + MatchBlock` to a list of options with that name
	// This means an option may be specified inside a match block, and to get this
	// option you need to know the match block it was specified in
	// If you want to get all options for a specific name, you can use the `AllOptionsPerName` field
	OptionsPerRelativeKey map[SSHIndexKey][]*ast.SSHOption

	// This is a map of `Option name` to a list of options with that name
	AllOptionsPerName map[string][]*SSHIndexAllOption
}

func CreateIndexes(config ast.SSHConfig) (*SSHIndexes, []common.LSPError) {
	errs := make([]common.LSPError, 0)
	indexes := &SSHIndexes{
		OptionsPerRelativeKey: make(map[SSHIndexKey][]*ast.SSHOption),
		AllOptionsPerName:     make(map[string][]*SSHIndexAllOption),
	}

	it := config.Options.Iterator()

	for it.Next() {
		entry := it.Value().(ast.SSHEntry)

		switch entry.(type) {
		case *ast.SSHOption:
			option := entry.(*ast.SSHOption)

			errs = append(errs, addOption(indexes, option, nil)...)
		case *ast.SSHMatchBlock:
			matchBlock := entry.(*ast.SSHMatchBlock)

			it := matchBlock.Options.Iterator()

			for it.Next() {
				option := it.Value().(*ast.SSHOption)

				errs = append(errs, addOption(indexes, option, matchBlock)...)
			}
		}
	}

	return indexes, errs
}

func addOption(
	i *SSHIndexes,
	option *ast.SSHOption,
	matchBlock *ast.SSHMatchBlock,
) []common.LSPError {
	var errs []common.LSPError

	indexEntry := SSHIndexKey{
		Option:     option.Key.Value,
		MatchBlock: matchBlock,
	}

	if existingEntry, found := i.OptionsPerRelativeKey[indexEntry]; found {
		if _, found := allowedDoubleOptions[option.Key.Value]; found {
			// Double value, but doubles are allowed
			i.OptionsPerRelativeKey[indexEntry] = append(existingEntry, option)
		} else {
			errs = append(errs, common.LSPError{
				Range: option.Key.LocationRange,
				Err:   errors.New(fmt.Sprintf("Option %s is already defined on line %d", option.Key.Value, existingEntry[0].Start.Line+1)),
			})
		}
	} else {
		i.OptionsPerRelativeKey[indexEntry] = []*ast.SSHOption{option}
	}


	if existingEntry, found := i.AllOptionsPerName[option.Key.Value]; found {
		i.AllOptionsPerName[option.Key.Value] = append(existingEntry, &SSHIndexAllOption{
			MatchBlock: matchBlock,
			Option:    option,
		})
	} else {
		i.AllOptionsPerName[option.Key.Value] = []*SSHIndexAllOption{
			{
				MatchBlock: matchBlock,
				Option:   option,
			},
		}
	}

	return errs
}
