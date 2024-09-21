package handlers

import (
	sshdconfig "config-lsp/handlers/sshd_config"
	"config-lsp/handlers/sshd_config/ast"

	protocol "github.com/tliron/glsp/protocol_3_16"
)

func FormatDocument(
	d *sshdconfig.SSHDocument,
	textRange protocol.Range,
	options protocol.FormattingOptions,
) ([]protocol.TextEdit, error) {
	edits := make([]protocol.TextEdit, 0)

	it := d.Config.Options.Iterator()
	for it.Next() {
		line := it.Key().(uint32)
		entry := it.Value().(ast.SSHDEntry)

		if !(line >= textRange.Start.Line && line <= textRange.End.Line) {
			continue
		}

		switch entry.(type) {
		case *ast.SSHDOption:
			option := entry.(*ast.SSHDOption)
			edits = append(edits, formatSSHDOption(option, options)...)
		case *ast.SSHDMatchBlock:
			matchBlock := entry.(*ast.SSHDMatchBlock)

			edits = formatSSHDMatchBlock(matchBlock, options)
		}
	}

	return edits, nil
}
