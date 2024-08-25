package analyzer

import (
	"config-lsp/common"
	"config-lsp/handlers/hosts/tree"
	"errors"
)

func analyzeEntriesAreValid(
	parser tree.HostsParser,
) []common.LSPError {
	err := make([]common.LSPError, 0)

	for lineNumber, entry := range parser.Tree.Entries {
		if entry.IPAddress == nil {
			err = append(err, common.LSPError{
				Range: common.CreateFullLineRange(lineNumber),
				Err:   errors.New("IP Address is required"),
			})
			continue
		}

		if entry.Hostname == nil {
			err = append(err, common.LSPError{
				Range: common.CreateFullLineRange(lineNumber),
				Err:   errors.New("Hostname is required"),
			})
			continue
		}
	}

	return err
}
