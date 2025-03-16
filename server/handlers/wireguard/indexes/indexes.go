package indexes

import "config-lsp/handlers/wireguard/ast"

type WGIndexes struct {
	// map of: section name -> WGSection
	SectionsByName map[string][]*ast.WGSection
}
