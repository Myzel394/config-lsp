package ini

import (
	"config-lsp/common"
	"github.com/emirpasic/gods/maps/treemap"
)

// PropertyKey represents a key in an INI property
type PropertyKey struct {
	common.LocationRange
	Name string
}

// PropertyValue represents a value in an INI property
type PropertyValue struct {
	common.LocationRange
	Value string
}

// PropertySeparator represents the separator in an INI property
type PropertySeparator struct {
	common.LocationRange
}

// Property represents a key-value pair in an INI document
type Property struct {
	common.LocationRange
	RawValue string

	Key       PropertyKey
	Separator *PropertySeparator
	Value     *PropertyValue
}

// Header represents a section header in an INI document
type Header struct {
	common.LocationRange
	Name string
}

// Section represents a section in an INI document
type Section struct {
	common.LocationRange
	// if `nil` = this is not a real section, but a placeholder for an empty section
	Header *Header
	// [uint32]*Property: line number -> *Property
	Properties *treemap.Map
}

// Config represents a complete INI document
type Config struct {
	Sections []*Section
	// Used to identify where not to show diagnostics
	CommentLines map[uint32]struct{}

	XParseConfig INIParseConfig // Configuration for parsing INI files
}

type INIParseConfig struct {
	// Allow setting properties outside of sections
	AllowRootProperties bool
}
