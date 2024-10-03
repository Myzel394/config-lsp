package handlers

import (
	sshconfig "config-lsp/handlers/ssh_config"
	"errors"

	protocol "github.com/tliron/glsp/protocol_3_16"
)

func RenameTag(
	params *protocol.RenameParams,
	d *sshconfig.SSHDocument,
	oldName string,
	newName string,
) (*protocol.WorkspaceEdit, error) {
	changes := make([]protocol.TextEdit, 0)

	// tag rename
	info, found := d.Indexes.Tags[oldName]
	if !found {
		return nil, errors.New("Tag could not be found")
	}

	changes = append(changes, protocol.TextEdit{
		Range:   info.EntryValue.ToLSPRange(),
		NewText: newName,
	})

	// Rename all occurrences
	for _, option := range d.Indexes.TagImports[oldName] {
		changes = append(changes, protocol.TextEdit{
			Range:   option.OptionValue.ToLSPRange(),
			NewText: newName,
		})
	}

	return &protocol.WorkspaceEdit{
		Changes: map[protocol.DocumentUri][]protocol.TextEdit{
			params.TextDocument.URI: changes,
		},
	}, nil
}
