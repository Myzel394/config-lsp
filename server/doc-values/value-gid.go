package docvalues

import (
	"config-lsp/common"
	"config-lsp/utils"
	"fmt"
	"strconv"

	protocol "github.com/tliron/glsp/protocol_3_16"
)

type GIDNotInPasswdErr struct{}

func (e GIDNotInPasswdErr) Error() string {
	return "This UID does not exist"
}

type InvalidGIDError struct{}

func (e InvalidGIDError) Error() string {
	return "This UID is invalid"
}

type GIDNotInRangeError struct{}

func (e GIDNotInRangeError) Error() string {
	return "UIDs must be between 0 and 65535"
}

type GIDValue struct {
	EnforceUsingExisting bool
}

func (v GIDValue) GetTypeDescription() []string {
	return []string{"Group ID"}
}

func (v GIDValue) DeprecatedCheckIsValid(value string) []*InvalidValue {
	uid, err := strconv.Atoi(value)

	if err != nil {
		return []*InvalidValue{{
			Err:   InvalidGIDError{},
			Start: 0,
			End:   uint32(len(value)),
		}}
	}

	if uid < 0 || uid > 65535 {
		return []*InvalidValue{{
			Err:   GIDNotInRangeError{},
			Start: 0,
			End:   uint32(len(value)),
		}}
	}

	if v.EnforceUsingExisting {
		infos, err := common.FetchPasswdInfo()

		if err != nil {
			return []*InvalidValue{}
		}

		found := false

		for _, info := range infos {
			if info.GID == value {
				found = true
				break
			}
		}

		if !found {
			return []*InvalidValue{{
				Err:   GIDNotInPasswdErr{},
				Start: 0,
				End:   uint32(len(value)),
			}}
		}
	}

	return []*InvalidValue{}
}

var defaultGIDsExplanation = []EnumString{
	{
		InsertText:      "0",
		DescriptionText: "root",
		Documentation:   "The group of the root user",
	},
}

func (v GIDValue) FetchCompletions(value string, cursor common.CursorPosition) []protocol.CompletionItem {
	infos, err := common.FetchGroupInfo()

	if err != nil {
		return utils.Map(defaultUIDsExplanation, func(enum EnumString) protocol.CompletionItem {
			return enum.ToCompletionItem()
		})
	}

	textFormat := protocol.InsertTextFormatPlainText
	kind := protocol.CompletionItemKindValue
	completions := make([]protocol.CompletionItem, len(infos))

	for index, info := range infos {
		// Find default gids
		var existingGID *EnumString

		for _, defaultGID := range defaultUIDsExplanation {
			if defaultGID.InsertText == info.GID {
				existingGID = &defaultGID
				break
			}
		}

		if existingGID != nil {
			completions[index] = existingGID.ToCompletionItem()
		} else {
			completions[index] = protocol.CompletionItem{
				InsertTextFormat: &textFormat,
				Kind:             &kind,
				InsertText:       &info.GID,
				Documentation:    info.Name,
			}
		}
	}

	return completions
}

func (v GIDValue) DeprecatedFetchHoverInfo(line string, cursor uint32) []string {
	uid, err := strconv.Atoi(line)

	if err != nil {
		return []string{}
	}

	infos, err := common.FetchGroupInfo()

	if err != nil {
		return []string{}
	}

	for _, info := range infos {
		if info.GID == strconv.Itoa(uid) {
			return []string{
				fmt.Sprintf("Group %s; GID: %s", info.Name, info.GID),
			}
		}
	}

	return []string{}
}
