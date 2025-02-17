// Contains structs that are used as utilities, but are
// not used for the AST itself
package ast

type SSHDOptionInfo struct {
	MatchBlock *SSHDMatchBlock
	Option     *SSHDOption
}
