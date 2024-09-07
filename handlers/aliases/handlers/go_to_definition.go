package handlers

import (
	"config-lsp/common"
	"config-lsp/handlers/aliases/ast"
	"config-lsp/handlers/aliases/indexes"
	"config-lsp/utils"

	protocol "github.com/tliron/glsp/protocol_3_16"
)

func GetDefinitionLocationForValue(
	i indexes.AliasesIndexes,
	value ast.AliasValueInterface,
	params *protocol.DefinitionParams,
) []protocol.Location {
	switch value.(type) {
	case ast.AliasValueUser:
		userValue := value.(ast.AliasValueUser)

		// Own defined alias
		if entry, found := i.Keys[indexes.NormalizeKey(userValue.Value)]; found {
			return []protocol.Location{
				{
					URI:   params.TextDocument.URI,
					Range: entry.Location.ToLSPRange(),
				},
			}
		}

		// System user
		systemUsers, _ := getSystemUserMap()
		if user, found := systemUsers[userValue.Value]; found {
			return []protocol.Location{
				{
					URI: "file:///etc/passwd",
					Range: protocol.Range{
						Start: protocol.Position{
							Line:      user.Line,
							Character: 0,
						},
						End: protocol.Position{
							Line:      user.Line,
							Character: uint32(len(user.Name)),
						},
					},
				},
			}
		}
	case ast.AliasValueFile:
		fileValue := value.(ast.AliasValueFile)
		path := string(fileValue.Path)

		if utils.DoesPathExist(path) {
			return []protocol.Location{
				{
					URI: "file://" + path,
				},
			}
		}
	case ast.AliasValueInclude:
		includeValue := value.(ast.AliasValueInclude)

		if includeValue.Path != nil {
			path := string(includeValue.Path.Path)

			if utils.DoesPathExist(path) {
				return []protocol.Location{
					{
						URI: "file://" + path,
					},
				}
			}
		}
	}

	return nil
}

func getSystemUserMap() (map[string]common.PasswdInfo, error) {
	users, err := common.FetchPasswdInfo()

	if err != nil {
		return nil, err
	}

	userMap := make(map[string]common.PasswdInfo)

	for _, user := range users {
		userMap[user.Name] = user
	}

	return userMap, nil
}
