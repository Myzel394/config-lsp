package indexes

import "config-lsp/parsers/ini"

type BTCIndexPropertyInfo struct {
	Section  *ini.Section
	Property *ini.Property
}

type BTCIndexes struct {
	// map of: section name -> ini.Section
	SectionsByName map[string]*ini.Section
	// map of: line number -> WGIndexPropertyInfo
	UnknownProperties map[uint32]BTCIndexPropertyInfo
}
