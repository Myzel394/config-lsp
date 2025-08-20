package indexes

import (
	"config-lsp/parsers/ini"
)

type WGIndexPropertyInfo struct {
	Section  *ini.Section
	Property *ini.Property
}

type AsymmetricRules struct {
	PreMissing  bool
	PostMissing bool
}

type WGIndexes struct {
	// map of: section name -> *ini.Section
	SectionsByName map[string][]*ini.Section

	// map of: line number -> WGIndexPropertyInfo
	UnknownProperties map[uint32]WGIndexPropertyInfo

	// Lists which properties are asymmetric (so that means they are missing a pre or post)
	AsymmetricRules map[*ini.Section]AsymmetricRules
}
