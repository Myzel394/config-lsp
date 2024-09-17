package indexes

import (
	"config-lsp/common"
	"config-lsp/handlers/sshd_config/ast"
	"config-lsp/handlers/sshd_config/fields"
	"errors"
	"fmt"
	"regexp"
)

type ValidPath string

func (v ValidPath) AsURI() string {
	return "file://" + string(v)
}

// SSHIndexIncludeValue Used to store the individual includes
// An `Include` statement can have multiple paths,
// each [SSHIndexIncludeValue] represents a single entered path.
// Note that an entered path can represent multiple real paths, as
// the path can contain wildcards.
// All true paths are stored in the [Paths] field.
type SSHIndexIncludeValue struct {
	common.LocationRange
	Value string

	// Actual valid paths, these will be set by the analyzer
	Paths []ValidPath
}

type SSHIndexIncludeLine struct {
	Values     []*SSHIndexIncludeValue
	Option     *ast.SSHDOption
	MatchBlock *ast.SSHDMatchBlock
}

type SSHIndexes struct {
	// This is a map of `Option name` to a list of options with that name
	AllOptionsPerName map[*ast.SSHDMatchBlock](map[string]([]*ast.SSHDOption))

	Includes map[uint32]*SSHIndexIncludeLine
}

func (i SSHIndexes) GetAllOptionsForName(name string) map[*ast.SSHDMatchBlock][]*ast.SSHDOption {
	allOptions := make(map[*ast.SSHDMatchBlock][]*ast.SSHDOption)

	for matchBlock, options := range i.AllOptionsPerName {
		if opts, found := options[name]; found {
			allOptions[matchBlock] = opts
		}
	}

	return allOptions
}

var whitespacePattern = regexp.MustCompile(`\S+`)

func CreateIndexes(config ast.SSHDConfig) (*SSHIndexes, []common.LSPError) {
	errs := make([]common.LSPError, 0)
	indexes := &SSHIndexes{
		AllOptionsPerName: make(map[*ast.SSHDMatchBlock](map[string]([]*ast.SSHDOption))),
		Includes:          make(map[uint32]*SSHIndexIncludeLine),
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
	for matchBlock, options := range indexes.GetAllOptionsForName("Include") {
		includeOption := options[0]
		rawValue := includeOption.OptionValue.Value.Value
		pathIndexes := whitespacePattern.FindAllStringIndex(rawValue, -1)
		paths := make([]*SSHIndexIncludeValue, 0)

		for _, pathIndex := range pathIndexes {
			startIndex := pathIndex[0]
			endIndex := pathIndex[1]

			rawPath := rawValue[startIndex:endIndex]

			offset := includeOption.OptionValue.Start.Character
			path := SSHIndexIncludeValue{
				LocationRange: common.LocationRange{
					Start: common.Location{
						Line:      includeOption.Start.Line,
						Character: uint32(startIndex) + offset,
					},
					End: common.Location{
						Line:      includeOption.Start.Line,
						Character: uint32(endIndex) + offset - 1,
					},
				},
				Value: rawPath,
				Paths: make([]ValidPath, 0),
			}

			paths = append(paths, &path)
		}

		indexes.Includes[includeOption.Start.Line] = &SSHIndexIncludeLine{
			Values:     paths,
			Option:     includeOption,
			MatchBlock: matchBlock,
		}
	}

	return indexes, errs
}

func addOption(
	i *SSHIndexes,
	option *ast.SSHDOption,
	matchBlock *ast.SSHDMatchBlock,
) []common.LSPError {
	var errs []common.LSPError

	if optionsMap, found := i.AllOptionsPerName[matchBlock]; found {
		if options, found := optionsMap[option.Key.Key]; found {
			if _, duplicatesAllowed := fields.AllowedDuplicateOptions[option.Key.Key]; !duplicatesAllowed {
				firstDefinedOption := options[0]
				errs = append(errs, common.LSPError{
					Range: option.Key.LocationRange,
					Err: errors.New(fmt.Sprintf(
						"Option '%s' has already been defined on line %d",
						option.Key.Key,
						firstDefinedOption.Start.Line,
					)),
				})
			} else {
				i.AllOptionsPerName[matchBlock][option.Key.Key] = append(
					i.AllOptionsPerName[matchBlock][option.Key.Key],
					option,
				)
			}
		} else {
			i.AllOptionsPerName[matchBlock][option.Key.Key] = []*ast.SSHDOption{
				option,
			}
		}
	} else {
		i.AllOptionsPerName[matchBlock] = map[string]([]*ast.SSHDOption){
			option.Key.Key: {
				option,
			},
		}
	}

	return errs
}
