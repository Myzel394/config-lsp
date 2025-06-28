package indexes

import (
	"config-lsp/common"
	"config-lsp/handlers/wireguard/ast"
	"config-lsp/parsers/ini"
)

func CreateIndexes(config *ast.WGConfig) (*WGIndexes, []common.LSPError) {
	errs := make([]common.LSPError, 0)
	indexes := &WGIndexes{
		SectionsByName:    make(map[string][]*ini.Section),
		UnknownProperties: make(map[uint32]WGIndexPropertyInfo),
	}

	// Use the WGSections from the config
	for _, section := range config.Sections {
		sectionName := section.Header.Name

		indexes.SectionsByName[sectionName] = append(
			indexes.SectionsByName[sectionName],
			section,
		)
	}

	return indexes, errs
}
