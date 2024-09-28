package sshconfig

import "config-lsp/handlers/ssh_config/ast"

func (d SSHDocument) FindOptionByNameAndBlock(
	name string,
	block ast.SSHBlock,
) *ast.AllOptionInfo {
	for _, info := range d.FindOptionsByName(name) {
		if info.Block == block {
			return &info
		}
	}

	return nil
}

func (d SSHDocument) FindOptionsByName(
	name string,
) []ast.AllOptionInfo {
	options := make([]ast.AllOptionInfo, 0, 5)

	for _, info := range d.Config.GetAllOptions() {
		if info.Option.Key.Key == name {
			options = append(options, info)
		}
	}

	return options
}

func (d SSHDocument) DoesOptionExist(
	name string,
	block ast.SSHBlock,
) bool {
	return d.FindOptionByNameAndBlock(name, block) != nil
}

