package indexes

import (
	"config-lsp/common"
	"config-lsp/handlers/ssh_config/ast"
	"config-lsp/handlers/ssh_config/fields"
	matchparser "config-lsp/handlers/ssh_config/match-parser"
)

type ValidPath string

func (v ValidPath) AsURI() string {
	return "file://" + string(v)
}

type SSHIndexIncludeValue struct {
	common.LocationRange
	Value string

	// Actual valid paths, these will be set by the analyzer
	Paths []ValidPath
}

type SSHIndexIncludeLine struct {
	Values []*SSHIndexIncludeValue
	Option *ast.SSHOption
	Block  ast.SSHBlock
}

type SSHIndexIgnoredUnknownInfo struct {
	common.LocationRange
}

type SSHIndexIgnoredUnknowns struct {
	OptionValue    *ast.SSHOption
	IgnoredOptions map[fields.NormalizedOptionName]SSHIndexIgnoredUnknownInfo
}

type SSHIndexTagInfo struct {
	EntryValue *matchparser.MatchValue
	Block      *ast.SSHMatchBlock
}

type SSHIndexes struct {
	AllOptionsPerName map[fields.NormalizedOptionName](map[ast.SSHBlock]([]*ast.SSHOption))

	Includes map[uint32]*SSHIndexIncludeLine

	BlockRanges map[uint32]ast.SSHBlock

	// Map of <block|nil (for global)> to a list of ignored options
	IgnoredOptions map[ast.SSHBlock]SSHIndexIgnoredUnknowns

	// This is used for code actions.
	// This stores a list of unknown option, so that we can provide
	// a code action to add them to a "IgnoreUnknown" option
	// This is a map of <line> to <option>
	UnknownOptions map[uint32]ast.AllOptionInfo

	// Map of available tags
	Tags map[string]SSHIndexTagInfo
	// Map of <Tag name> to <option per block>
	TagImports map[string](map[ast.SSHBlock]*ast.SSHOption)
}
