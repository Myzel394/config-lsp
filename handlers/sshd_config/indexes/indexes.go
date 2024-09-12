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

type SSHIndexEntry struct {
	Option     string
	MatchBlock *ast.SSHMatchBlock
}

type SSHIndexes struct {
	EntriesPerKey map[SSHIndexEntry][]*ast.SSHOption
}

func CreateIndexes(config ast.SSHConfig) (*SSHIndexes, []common.LSPError) {
	errs := make([]common.LSPError, 0)
	indexes := &SSHIndexes{
		EntriesPerKey: make(map[SSHIndexEntry][]*ast.SSHOption),
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

	indexEntry := SSHIndexEntry{
		Option:     option.Key.Value,
		MatchBlock: matchBlock,
	}

	if existingEntry, found := i.EntriesPerKey[indexEntry]; found {
		if _, found := allowedDoubleOptions[option.Key.Value]; found {
			// Double value, but doubles are allowed
			i.EntriesPerKey[indexEntry] = append(existingEntry, option)
		} else {
			errs = append(errs, common.LSPError{
				Range: option.Key.LocationRange,
				Err:   errors.New(fmt.Sprintf("Option %s is already defined on line %d", option.Key.Value, existingEntry[0].Start.Line+1)),
			})
		}
	} else {
		i.EntriesPerKey[indexEntry] = []*ast.SSHOption{option}
	}

	return errs
}
