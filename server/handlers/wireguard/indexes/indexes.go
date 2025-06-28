package indexes

import (
	"config-lsp/parsers/ini"
)

type WGIndexPropertyInfo struct {
	Section  *ini.Section
	Property *ini.Property
}

type WGIndexes struct {
	// map of: section name -> *ini.Section
	SectionsByName map[string][]*ini.Section

	// map of: line number -> WGIndexPropertyInfo
	UnknownProperties map[uint32]WGIndexPropertyInfo

	// map of: line number -> WGIndexPropertyInfo (PreUp / PostUp properties)
	UpProperties map[uint32]WGIndexPropertyInfo
}
