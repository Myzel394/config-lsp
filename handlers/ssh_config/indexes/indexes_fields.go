package indexes

import (
	"config-lsp/handlers/ssh_config/ast"
	"config-lsp/handlers/ssh_config/fields"
)

func (u SSHIndexIgnoredUnknowns) GetIgnoredForLine(line uint32) map[fields.NormalizedOptionName]struct{} {
	if line >= u.OptionValue.Start.Line {
		return u.IgnoredOptions
	}

	return nil
}

func (i SSHIndexes) CanOptionBeIgnored(
	option *ast.SSHOption,
	block ast.SSHBlock,
) bool {
	ignores, found := i.IgnoredOptions[block]

	if !found {
		return false
	}

	ignoredOptions := ignores.GetIgnoredForLine(option.Start.Line)

	_, found = ignoredOptions[option.Key.Key]

	return found
}
