package handlers

import (
	"config-lsp/common"
	"config-lsp/handlers/ssh_config/indexes"
	"config-lsp/utils"
	"slices"

	protocol "github.com/tliron/glsp/protocol_3_16"
)

func GetIncludeOptionLocation(
	include *indexes.SSHIndexIncludeLine,
	index common.IndexPosition,
) []protocol.Location {
	foundIndex, found := slices.BinarySearchFunc(
		include.Values,
		index,
		func(current *indexes.SSHIndexIncludeValue, target common.IndexPosition) int {
			if current.IsPositionAfterEnd(target) {
				return -1
			}

			if current.IsPositionBeforeStart(target) {
				return 1
			}

			return 0
		},
	)

	if !found {
		return nil
	}

	path := include.Values[foundIndex]

	return utils.Map(path.Paths, func(path indexes.ValidPath) protocol.Location {
		return protocol.Location{
			URI: path.AsURI(),
		}
	})
}
