package indexes

import (
	"config-lsp/common"
	"config-lsp/handlers/wireguard/ast"
	"config-lsp/handlers/wireguard/fields"
	"config-lsp/parsers/ini"
)

var preUp = fields.CreateNormalizedName("PreUp")
var postUp = fields.CreateNormalizedName("PostUp")
var preDown = fields.CreateNormalizedName("PreDown")
var postDown = fields.CreateNormalizedName("PostDown")

func CreateIndexes(config *ast.WGConfig) (*WGIndexes, []common.LSPError) {
	indexes := &WGIndexes{
		SectionsByName:    make(map[string][]*ini.Section),
		UnknownProperties: make(map[uint32]WGIndexPropertyInfo),
		AsymmetricRules:   make(map[*ini.Section]AsymmetricRules),
	}

	// Use the WGSections from the config
	for _, section := range config.Sections {
		sectionName := section.Header.Name

		indexes.SectionsByName[sectionName] = append(
			indexes.SectionsByName[sectionName],
			section,
		)

		preCount := 0
		postCount := 0

		it := section.Properties.Iterator()
		for it.Next() {
			// lineNumber := it.Key().(uint32)
			property := it.Value().(*ini.Property)
			normalizedName := fields.CreateNormalizedName(property.Key.Name)

			switch normalizedName {
			case preUp:
				preCount++
			case postUp:
				postCount++
			case preDown:
				preCount--
			case postDown:
				postCount--
			}
		}

		indexes.AsymmetricRules[section] = AsymmetricRules{
			PreMissing:  preCount != 0,
			PostMissing: postCount != 0,
		}
	}

	return indexes, nil
}
