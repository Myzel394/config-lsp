package docvalues

import (
	"config-lsp/common"
	"config-lsp/utils"
	"fmt"
	"strconv"

	protocol "github.com/tliron/glsp/protocol_3_16"
)

type UIDNotInPasswdErr struct{}

func (e UIDNotInPasswdErr) Error() string {
	return "This UID does not exist"
}

type InvalidUIDError struct{}

func (e InvalidUIDError) Error() string {
	return "This UID is invalid"
}

type UIDNotInRangeError struct{}

func (e UIDNotInRangeError) Error() string {
	return "UIDs must be between 0 and 65535"
}

type UIDValue struct {
	EnforceUsingExisting bool
}

func (v UIDValue) GetTypeDescription() []string {
	return []string{"User ID"}
}

func (v UIDValue) DeprecatedCheckIsValid(value string) []*InvalidValue {
	uid, err := strconv.Atoi(value)

	if err != nil {
		return []*InvalidValue{{
			Err:   InvalidUIDError{},
			Start: 0,
			End:   uint32(len(value)),
		}}
	}

	if uid < 0 || uid > 65535 {
		return []*InvalidValue{{
			Err:   UIDNotInRangeError{},
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
			if info.UID == value {
				found = true
				break
			}
		}

		if !found {
			return []*InvalidValue{{
				Err:   UIDNotInPasswdErr{},
				Start: 0,
				End:   uint32(len(value)),
			}}
		}
	}

	return []*InvalidValue{}
}

var defaultUIDsExplanation = []EnumString{
	{
		InsertText:      "0",
		DescriptionText: "root",
		Documentation:   "The root user",
	},
}

func (v UIDValue) FetchCompletions(value string, cursor common.CursorPosition) []protocol.CompletionItem {
	infos, err := common.FetchPasswdInfo()

	if err != nil {
		return utils.Map(defaultUIDsExplanation, func(enum EnumString) protocol.CompletionItem {
			return enum.ToCompletionItem()
		})
	}

	textFormat := protocol.InsertTextFormatPlainText
	kind := protocol.CompletionItemKindValue
	completions := make([]protocol.CompletionItem, len(infos))

	for index, info := range infos {
		// Find default uids
		var existingUID *EnumString

		for _, defaultUID := range defaultUIDsExplanation {
			if defaultUID.InsertText == info.UID {
				existingUID = &defaultUID
				break
			}
		}

		if existingUID != nil {
			completions[index] = existingUID.ToCompletionItem()
		} else {
			completions[index] = protocol.CompletionItem{
				InsertTextFormat: &textFormat,
				Kind:             &kind,
				InsertText:       &info.UID,
				Documentation:    fmt.Sprintf("User %s; Home: %s", info.Name, info.HomePath),
			}
		}
	}

	return completions
}

func (v UIDValue) DeprecatedFetchHoverInfo(line string, cursor uint32) []string {
	uid, err := strconv.Atoi(line)

	if err != nil {
		return []string{}
	}

	infos, err := common.FetchPasswdInfo()

	if err != nil {
		return []string{}
	}

	for _, info := range infos {
		if info.UID == strconv.Itoa(uid) {
			return []string{
				fmt.Sprintf("User %s; ID: %s:%s; ~:%s", info.Name, info.UID, info.GID, info.HomePath),
			}
		}
	}

	return []string{}
}
