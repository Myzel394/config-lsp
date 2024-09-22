package ast

import "config-lsp/common"


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
	Values     []*SSHIndexIncludeValue
	Option     *SSHOption
	Block *SSHBlock
}

type SSHIndexes struct {
	Includes []*SSHIndexIncludeLine
}

func NewSSHIndexes() *SSHIndexes {
	return &SSHIndexes{
		Includes: make([]*SSHIndexIncludeLine, 0),
	}
}

