package indexes

import (
	"config-lsp/common"
	"config-lsp/handlers/ssh_config/ast"
	"config-lsp/handlers/ssh_config/fields"
	matchparser "config-lsp/handlers/ssh_config/match-parser"
	"errors"
	"fmt"
	"regexp"
	"strings"
)

var includePattern = regexp.MustCompile(`".+?"|[^\s]+`)

func NewSSHIndexes() *SSHIndexes {
	return &SSHIndexes{
		AllOptionsPerName: make(map[fields.NormalizedOptionName](map[ast.SSHBlock]([]*ast.SSHOption))),
		Includes:          make(map[uint32]*SSHIndexIncludeLine),
		IgnoredOptions:    make(map[ast.SSHBlock]SSHIndexIgnoredUnknowns),
		UnknownOptions:    make(map[uint32]ast.AllOptionInfo),
		Tags:              make(map[string]SSHIndexTagInfo),
		TagImports:        make(map[string](map[ast.SSHBlock]*ast.SSHOption)),
	}
}

var includeOption = fields.CreateNormalizedName("Include")
var matchOption = fields.CreateNormalizedName("Match")
var tagOption = fields.CreateNormalizedName("Tag")

func CreateIndexes(config ast.SSHConfig) (*SSHIndexes, []common.LSPError) {
	errs := make([]common.LSPError, 0)
	indexes := NewSSHIndexes()

	it := config.Options.Iterator()
	for it.Next() {
		entry := it.Value().(ast.SSHEntry)

		switch entry.GetType() {
		case ast.SSHTypeOption:
			errs = append(errs, addOption(indexes, entry.GetOption(), nil)...)
		case ast.SSHTypeMatch:
			fallthrough
		case ast.SSHTypeHost:
			block := entry.(ast.SSHBlock)

			errs = append(errs, addOption(indexes, entry.GetOption(), block)...)

			it := block.GetOptions().Iterator()
			for it.Next() {
				option := it.Value().(*ast.SSHOption)

				errs = append(errs, addOption(indexes, option, block)...)
			}
		}
	}

	// Add Includes
	for block, options := range indexes.AllOptionsPerName[includeOption] {
		includeOption := options[0]

		// We want to parse quotes manually
		rawValue := includeOption.OptionValue.Value.Raw
		pathIndexes := includePattern.FindAllStringIndex(rawValue, -1)
		paths := make([]*SSHIndexIncludeValue, 0)

		for _, pathIndex := range pathIndexes {
			startIndex := pathIndex[0]
			endIndex := pathIndex[1]

			rawPath := strings.ReplaceAll(
				rawValue[startIndex:endIndex],
				`"`,
				"",
			)

			offset := includeOption.OptionValue.Start.Character
			path := SSHIndexIncludeValue{
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

		indexes.Includes[includeOption.Start.Line] = &SSHIndexIncludeLine{
			Values: paths,
			Option: includeOption,
			Block:  block,
		}
	}

	// Add tags
	for rawBlock := range indexes.AllOptionsPerName[matchOption] {
		block := rawBlock.(*ast.SSHMatchBlock)
		for _, entry := range block.MatchValue.Entries {
			if entry.Criteria.Type != matchparser.MatchCriteriaTypeTagged || entry.Value.Value == "" {
				continue
			}

			for _, value := range entry.Values.Values {
				name := value.Value.Value

				if existingBlock, found := indexes.Tags[name]; found {
					// Tag already exists
					errs = append(errs, common.LSPError{
						Range: entry.LocationRange,
						Err:   fmt.Errorf("Tag %s has already been defined on line %d", name, existingBlock.Block.Start.Line+1),
					})

					continue
				}

				// Add tag
				indexes.Tags[name] = SSHIndexTagInfo{
					EntryValue: value,
					Block:      block,
				}
			}
		}
	}

	// Add Tag imports
	for block, options := range indexes.AllOptionsPerName[tagOption] {
		for _, option := range options {
			if option.OptionValue == nil || option.OptionValue.Value.Value == "" {
				continue
			}

			tagName := option.OptionValue.Value.Value

			if tagImport, found := indexes.TagImports[tagName]; found {
				if definedOption, found := tagImport[block]; found {
					errs = append(errs, common.LSPError{
						Range: option.OptionValue.LocationRange,
						Err:   fmt.Errorf("Tag %s has already been imported on line %d", tagName, definedOption.Start.Line+1),
					})

					continue
				} else {
					tagImport[block] = option
				}
			} else {
				indexes.TagImports[tagName] = map[ast.SSHBlock]*ast.SSHOption{
					block: option,
				}
			}
		}
	}

	return indexes, errs
}

var ignoreUnknownOption = fields.CreateNormalizedName("IgnoreUnknown")

func addOption(
	i *SSHIndexes,
	option *ast.SSHOption,
	block ast.SSHBlock,
) []common.LSPError {
	var errs []common.LSPError

	if optionsMap, found := i.AllOptionsPerName[option.Key.Key]; found {
		if options, found := optionsMap[block]; found {
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
				i.AllOptionsPerName[option.Key.Key][block] = append(
					i.AllOptionsPerName[option.Key.Key][block],
					option,
				)
			}
		} else {
			i.AllOptionsPerName[option.Key.Key][block] = []*ast.SSHOption{
				option,
			}
		}
	} else {
		i.AllOptionsPerName[option.Key.Key] = map[ast.SSHBlock]([]*ast.SSHOption){
			block: {
				option,
			},
		}
	}

	ignoredOptions, found := i.AllOptionsPerName[ignoreUnknownOption]

	if found {
		for block, options := range ignoredOptions {
			// Only using first options as multiple `IgnoreUnknown`s are not allowed anyway
			addIgnoredOption(i, options[0], block)
		}
	}

	return errs
}

var ignoredValuesPattern = regexp.MustCompile(`\S+`)

func addIgnoredOption(
	i *SSHIndexes,
	option *ast.SSHOption,
	block ast.SSHBlock,
) {
	rawIgnored := option.OptionValue.Value.Value
	ignoredAsSlice := ignoredValuesPattern.FindAllStringIndex(rawIgnored, -1)
	ignored := make(map[fields.NormalizedOptionName]SSHIndexIgnoredUnknownInfo, 0)

	for _, ignoreInfo := range ignoredAsSlice {
		start := ignoreInfo[0]
		end := ignoreInfo[1]
		name := rawIgnored[start:end]

		ignored[fields.CreateNormalizedName(name)] = SSHIndexIgnoredUnknownInfo{
			LocationRange: common.LocationRange{
				Start: common.Location{
					Line:      option.OptionValue.Start.Line,
					Character: option.OptionValue.Start.Character + uint32(start),
				},
				End: common.Location{
					Line:      option.OptionValue.End.Line,
					Character: option.OptionValue.Start.Character + uint32(end),
				},
			},
		}
	}

	i.IgnoredOptions[block] = SSHIndexIgnoredUnknowns{
		OptionValue:    option,
		IgnoredOptions: ignored,
	}
}
