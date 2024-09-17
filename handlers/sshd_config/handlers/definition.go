package handlers

import (
	"config-lsp/handlers/sshd_config/indexes"
	"config-lsp/utils"
	"fmt"
	"slices"

	protocol "github.com/tliron/glsp/protocol_3_16"
)

func GetIncludeOptionLocation(
	include *indexes.SSHIndexIncludeLine,
	cursor uint32,
) []protocol.Location {
	index, found := slices.BinarySearchFunc(
		include.Values,
		cursor,
		func(i *indexes.SSHIndexIncludeValue, cursor uint32) int {
			if cursor < i.Start.Character {
				return -1
			}

			if cursor > i.End.Character {
				return 1
			}

			return 0
		},
	)

	if !found {
		return nil
	}

	path := include.Values[index]
	println("paths", fmt.Sprintf("%v", path.Paths))

	return utils.Map(
		path.Paths,
		func(path indexes.ValidPath) protocol.Location {
			return protocol.Location{
				URI: path.AsURI(),
			}
		},
	)
}