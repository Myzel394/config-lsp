package indexes

import (
	"config-lsp/common"
	"config-lsp/handlers/sshd_config/ast"
)

type ValidPath string

func (v ValidPath) AsURI() string {
	return "file://" + string(v)
}

// SSHDIndexIncludeValue Used to store the individual includes
// An `Include` statement can have multiple paths,
// each [SSHDIndexIncludeValue] represents a single entered path.
// Note that an entered path can represent multiple real paths, as
// the path can contain wildcards.
// All true paths are stored in the [Paths] field.
type SSHDIndexIncludeValue struct {
	common.LocationRange
	Value string

	// Actual valid paths, these will be set by the analyzer
	Paths []ValidPath
}

type SSHDIndexIncludeLine struct {
	Values     []*SSHDIndexIncludeValue
	Option     *ast.SSHDOption
	MatchBlock *ast.SSHDMatchBlock
}

type SSHDIndexes struct {
	// This is a map of `Option name` to a list of options with that name
	AllOptionsPerName map[string](map[*ast.SSHDMatchBlock]([]*ast.SSHDOption))

	Includes map[uint32]*SSHDIndexIncludeLine
}
