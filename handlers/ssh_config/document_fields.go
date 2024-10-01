package sshconfig

import (
	"config-lsp/handlers/ssh_config/ast"
	"config-lsp/handlers/ssh_config/fields"
	"config-lsp/utils"
)

func (d SSHDocument) FindOptionByNameAndBlock(
	option fields.NormalizedOptionName,
	block ast.SSHBlock,
) *ast.AllOptionInfo {
	for _, info := range d.FindOptionsByName(option) {
		if info.Block == block {
			return &info
		}
	}

	return nil
}

func (d SSHDocument) FindOptionsByName(
	option fields.NormalizedOptionName,
) []ast.AllOptionInfo {
	options := make([]ast.AllOptionInfo, 0, 5)

	for _, info := range d.Config.GetAllOptions() {
		if info.Option.Key.Key == option {
			options = append(options, info)
		}
	}

	return options
}

func (d SSHDocument) DoesOptionExist(
	option fields.NormalizedOptionName,
	block ast.SSHBlock,
) bool {
	return d.FindOptionByNameAndBlock(option, block) != nil
}

var matchOption = fields.CreateNormalizedName("Match")

func (d SSHDocument) GetAllMatchBlocks() []*ast.SSHMatchBlock {
	matchBlocks := make([]*ast.SSHMatchBlock, 0, 5)

	options := d.Indexes.AllOptionsPerName[matchOption]
	blocks := utils.KeysOfMap(options)

	for _, block := range blocks {
		matchBlocks = append(matchBlocks, block.(*ast.SSHMatchBlock))
	}

	return matchBlocks
}

var hostOption = fields.CreateNormalizedName("Host")

func (d SSHDocument) GetAllHostBlocks() []*ast.SSHHostBlock {
	hostBlocks := make([]*ast.SSHHostBlock, 0, 5)

	options := d.Indexes.AllOptionsPerName[hostOption]
	blocks := utils.KeysOfMap(options)

	for _, block := range blocks {
		hostBlocks = append(hostBlocks, block.(*ast.SSHHostBlock))
	}

	return hostBlocks
}

// GetAllBlocks returns all blocks in the document
// Note: The blocks are **not** sorted
func (d SSHDocument) GetAllBlocks() []ast.SSHBlock {
	blocks := make([]ast.SSHBlock, 0)

	for _, block := range d.GetAllHostBlocks() {
		blocks = append(blocks, block)
	}

	for _, block := range d.GetAllMatchBlocks() {
		blocks = append(blocks, block)
	}

	return blocks
}
