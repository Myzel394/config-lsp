package indexes

import "config-lsp/handlers/wireguard/ast"

type WGIndexPropertyInfo struct {
	Section  *ast.WGSection
	Property *ast.WGProperty
}

type WGIndexes struct {
	// map of: section name -> *WGSection
	SectionsByName map[string][]*ast.WGSection

	// map of: line number -> *WGIndexPropertyInfo
	UnknownProperties map[uint32]WGIndexPropertyInfo
}
