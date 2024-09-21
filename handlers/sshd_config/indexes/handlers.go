package indexes

import (
	"config-lsp/common"
	"config-lsp/handlers/sshd_config/ast"
	"config-lsp/handlers/sshd_config/fields"
	"errors"
	"fmt"
	"regexp"
)

var whitespacePattern = regexp.MustCompile(`\S+`)

func CreateIndexes(config ast.SSHDConfig) (*SSHDIndexes, []common.LSPError) {
	errs := make([]common.LSPError, 0)
	indexes := &SSHDIndexes{
		AllOptionsPerName: make(map[string](map[*ast.SSHDMatchBlock]([]*ast.SSHDOption))),
		Includes:          make(map[uint32]*SSHDIndexIncludeLine),
	}

	it := config.Options.Iterator()
	for it.Next() {
		entry := it.Value().(ast.SSHDEntry)

		switch entry.(type) {
		case *ast.SSHDOption:
			option := entry.(*ast.SSHDOption)

			errs = append(errs, addOption(indexes, option, nil)...)
		case *ast.SSHDMatchBlock:
			matchBlock := entry.(*ast.SSHDMatchBlock)

			errs = append(errs, addOption(indexes, matchBlock.MatchEntry, matchBlock)...)

			it := matchBlock.Options.Iterator()
			for it.Next() {
				option := it.Value().(*ast.SSHDOption)

				errs = append(errs, addOption(indexes, option, matchBlock)...)
			}
		}
	}

	// Add Includes
	for matchBlock, options := range indexes.AllOptionsPerName["Include"] {
		includeOption := options[0]
		rawValue := includeOption.OptionValue.Value.Value
		pathIndexes := whitespacePattern.FindAllStringIndex(rawValue, -1)
		paths := make([]*SSHDIndexIncludeValue, 0)

		for _, pathIndex := range pathIndexes {
			startIndex := pathIndex[0]
			endIndex := pathIndex[1]

			rawPath := rawValue[startIndex:endIndex]

			offset := includeOption.OptionValue.Start.Character
			path := SSHDIndexIncludeValue{
				LocationRange: common.LocationRange{
					Start: common.Location{
						Line:      includeOption.Start.Line,
						Character: uint32(startIndex) + offset,
					},
					End: common.Location{
						Line:      includeOption.Start.Line,
						Character: uint32(endIndex) + offset,
					},
				},
				Value: rawPath,
				Paths: make([]ValidPath, 0),
			}

			paths = append(paths, &path)
		}

		indexes.Includes[includeOption.Start.Line] = &SSHDIndexIncludeLine{
			Values:     paths,
			Option:     includeOption,
			MatchBlock: matchBlock,
		}
	}

	return indexes, errs
}

func addOption(
	i *SSHDIndexes,
	option *ast.SSHDOption,
	matchBlock *ast.SSHDMatchBlock,
) []common.LSPError {
	var errs []common.LSPError

	if optionsMap, found := i.AllOptionsPerName[option.Key.Key]; found {
		if options, found := optionsMap[matchBlock]; found {
			if _, duplicatesAllowed := fields.AllowedDuplicateOptions[option.Key.Key]; !duplicatesAllowed {
				firstDefinedOption := options[0]
				errs = append(errs, common.LSPError{
					Range: option.Key.LocationRange,
					Err: errors.New(fmt.Sprintf(
						"Option '%s' has already been defined on line %d",
						option.Key.Key,
						firstDefinedOption.Start.Line+1,
					)),
				})
			} else {
				i.AllOptionsPerName[option.Key.Key][matchBlock] = append(
					i.AllOptionsPerName[option.Key.Key][matchBlock],
					option,
				)
			}
		} else {
			i.AllOptionsPerName[option.Key.Key][matchBlock] = []*ast.SSHDOption{
				option,
			}
		}
	} else {
		i.AllOptionsPerName[option.Key.Key] = map[*ast.SSHDMatchBlock]([]*ast.SSHDOption){
			matchBlock: {
				option,
			},
		}
	}

	return errs
}
