package indexes

import (
	"config-lsp/common"
	"config-lsp/handlers/sshd_config/ast"
	"errors"
	"fmt"
	"regexp"
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
	Option     *ast.SSHOption
}

type ValidPath string

func (v ValidPath) AsURI() string {
	return "file://" + string(v)
}

type SSHIndexIncludeValue struct {
	common.LocationRange
	Value string

	// Actual valid paths, these will be set by the analyzer
	Paths []ValidPath
}

type SSHIndexIncludeLine struct {
	Values []*SSHIndexIncludeValue
	Option *SSHIndexAllOption
}

type SSHIndexes struct {
	// Contains a map of `Option name + MatchBlock` to a list of options with that name
	// This means an option may be specified inside a match block, and to get this
	// option you need to know the match block it was specified in
	// If you want to get all options for a specific name, you can use the `AllOptionsPerName` field
	OptionsPerRelativeKey map[SSHIndexKey][]*ast.SSHOption

	// This is a map of `Option name` to a list of options with that name
	AllOptionsPerName map[string]map[uint32]*SSHIndexAllOption

	Includes map[uint32]*SSHIndexIncludeLine
}

var whitespacePattern = regexp.MustCompile(`\S+`)

func CreateIndexes(config ast.SSHConfig) (*SSHIndexes, []common.LSPError) {
	errs := make([]common.LSPError, 0)
	indexes := &SSHIndexes{
		OptionsPerRelativeKey: make(map[SSHIndexKey][]*ast.SSHOption),
		AllOptionsPerName:     make(map[string]map[uint32]*SSHIndexAllOption),
		Includes:              make(map[uint32]*SSHIndexIncludeLine),
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

	// Add Includes
	for _, includeOption := range indexes.AllOptionsPerName["Include"] {
		rawValue := includeOption.Option.OptionValue.Value
		pathIndexes := whitespacePattern.FindAllStringIndex(rawValue, -1)
		paths := make([]*SSHIndexIncludeValue, 0)

		for _, pathIndex := range pathIndexes {
			startIndex := pathIndex[0]
			endIndex := pathIndex[1]

			rawPath := rawValue[startIndex:endIndex]

			offset := includeOption.Option.OptionValue.Start.Character
			path := SSHIndexIncludeValue{
				LocationRange: common.LocationRange{
					Start: common.Location{
						Line:      includeOption.Option.Start.Line,
						Character: uint32(startIndex) + offset,
					},
					End: common.Location{
						Line:      includeOption.Option.Start.Line,
						Character: uint32(endIndex) + offset - 1,
					},
				},
				Value: rawPath,
				Paths: make([]ValidPath, 0),
			}

			paths = append(paths, &path)
		}

		indexes.Includes[includeOption.Option.Start.Line] = &SSHIndexIncludeLine{
			Values: paths,
			Option: includeOption,
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

	if _, found := i.AllOptionsPerName[option.Key.Value]; found {
		i.AllOptionsPerName[option.Key.Value][option.Start.Line] = &SSHIndexAllOption{
			MatchBlock: matchBlock,
			Option:     option,
		}
	} else {
		i.AllOptionsPerName[option.Key.Value] = map[uint32]*SSHIndexAllOption{
			option.Start.Line: {
				MatchBlock: matchBlock,
				Option:     option,
			},
		}
	}

	return errs
}
