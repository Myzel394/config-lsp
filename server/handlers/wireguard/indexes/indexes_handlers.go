package indexes

import (
	"config-lsp/common"
	"config-lsp/handlers/wireguard/ast"
)

func CreateIndexes(config *ast.WGConfig) (*WGIndexes, []common.LSPError) {
	errs := make([]common.LSPError, 0)
	indexes := &WGIndexes{
		SectionsByName:    make(map[string][]*ast.WGSection),
		UnknownProperties: make(map[uint32]WGIndexPropertyInfo),
	}

	for _, section := range config.Sections {
		indexes.SectionsByName[section.Header.Name] = append(
			indexes.SectionsByName[section.Header.Name],
			section,
		)
	}

	return indexes, errs
}
