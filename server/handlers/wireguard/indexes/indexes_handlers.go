package indexes

import (
	"config-lsp/common"
	"config-lsp/handlers/wireguard/ast"
	"config-lsp/handlers/wireguard/fields"
	"config-lsp/parsers/ini"
)

var upPropertyNames = map[fields.NormalizedName]struct{}{
	fields.CreateNormalizedName("PreUp"):  {},
	fields.CreateNormalizedName("PostUp"): {},
}

func CreateIndexes(config *ast.WGConfig) (*WGIndexes, []common.LSPError) {
	indexes := &WGIndexes{
		SectionsByName:    make(map[string][]*ini.Section),
		UnknownProperties: make(map[uint32]WGIndexPropertyInfo),
		UpProperties:      make(map[uint32]WGIndexPropertyInfo),
	}

	// Use the WGSections from the config
	for _, section := range config.Sections {
		sectionName := section.Header.Name

		indexes.SectionsByName[sectionName] = append(
			indexes.SectionsByName[sectionName],
			section,
		)

		it := section.Properties.Iterator()
		for it.Next() {
			lineNumber := it.Key().(uint32)
			property := it.Value().(*ini.Property)
			normalizedName := fields.CreateNormalizedName(property.Key.Name)

			if _, found := upPropertyNames[normalizedName]; found {
				indexes.UpProperties[lineNumber] = WGIndexPropertyInfo{
					Section:  section,
					Property: property,
				}
			}
		}
	}

	return indexes, nil
}
