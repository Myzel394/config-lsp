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

var includeOption = fields.CreateNormalizedName("Include")

func CreateIndexes(config ast.SSHDConfig) (*SSHDIndexes, []common.LSPError) {
	errs := make([]common.LSPError, 0)
	indexes := &SSHDIndexes{
		AllOptionsPerName: make(map[fields.NormalizedOptionName](map[*ast.SSHDMatchBlock]([]*ast.SSHDOption))),
		Includes:          make(map[uint32]*SSHDIndexIncludeLine),
		UnknownOptions:    make(map[uint32]ast.SSHDOptionInfo),
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

			errs = append(errs, addOption(indexes, matchBlock.MatchOption, matchBlock)...)

			it := matchBlock.Options.Iterator()
			for it.Next() {
				option := it.Value().(*ast.SSHDOption)

				errs = append(errs, addOption(indexes, option, matchBlock)...)
			}
		}
	}

	// Add Includes
	for matchBlock, options := range indexes.AllOptionsPerName[includeOption] {
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
	key := option.Key.Key

	if optionsMap, found := i.AllOptionsPerName[key]; found {
		if options, found := optionsMap[matchBlock]; found {
			if _, duplicatesAllowed := fields.AllowedDuplicateOptions[key]; !duplicatesAllowed {
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
				i.AllOptionsPerName[key][matchBlock] = append(
					i.AllOptionsPerName[key][matchBlock],
					option,
				)
			}
		} else {
			i.AllOptionsPerName[key][matchBlock] = []*ast.SSHDOption{
				option,
			}
		}
	} else {
		i.AllOptionsPerName[key] = map[*ast.SSHDMatchBlock]([]*ast.SSHDOption){
			matchBlock: {
				option,
			},
		}
	}

	return errs
}
