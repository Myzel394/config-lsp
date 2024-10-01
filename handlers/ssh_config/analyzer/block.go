package analyzer

import (
	"config-lsp/common"
	sshconfig "config-lsp/handlers/ssh_config"
	"errors"
)

func analyzeBlocks(
	d *sshconfig.SSHDocument,
) []common.LSPError {
	errs := make([]common.LSPError, 0)

	for _, block := range d.GetAllBlocks() {
		if block.GetOptions().Size() == 0 {
			errs = append(errs, common.LSPError{
				Range: block.GetEntryOption().LocationRange,
				Err:   errors.New("This block is empty"),
			})
		}
	}

	return errs
}
