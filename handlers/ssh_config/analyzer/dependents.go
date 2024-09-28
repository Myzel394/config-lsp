package analyzer

import (
	"config-lsp/common"
	sshconfig "config-lsp/handlers/ssh_config"
	"config-lsp/handlers/ssh_config/ast"
	"config-lsp/handlers/ssh_config/fields"
	"errors"
	"fmt"
)

func analyzeDependents(
	d *sshconfig.SSHDocument,
) []common.LSPError {
	errs := make([]common.LSPError, 0)

	for _, option := range d.Config.GetAllOptions() {
		errs = append(errs, checkIsDependent(d, option.Option.Key, option.Block)...)
	}

	return errs
}

func checkIsDependent(
	d *sshconfig.SSHDocument,
	key *ast.SSHKey,
	block ast.SSHBlock,
) []common.LSPError {
	errs := make([]common.LSPError, 0)

	dependentOptions, found := fields.DependentFields[key.Key]

	if !found {
		return errs
	}

	for _, dependentOption := range dependentOptions {
		if opts, found := d.Indexes.AllOptionsPerName[dependentOption]; found {
			_, existsInBlock := opts[block]
			_, existsInGlobal := opts[nil]

			if existsInBlock || existsInGlobal {
				continue
			}
		}

		errs = append(errs, common.LSPError{
			Range: key.LocationRange,
			Err:   errors.New(fmt.Sprintf("Option '%s' requires option '%s' to be present", key.Key, dependentOption)),
		})
	}

	return errs
}

