package fetchers

import (
	"config-lsp/common"
	"config-lsp/handlers/aliases/indexes"
	"fmt"
	"strings"
)

type aliasesUser struct {
	DefinedOnLine uint32
}

type User struct {
	PasswdInfo *common.PasswdInfo

	AliasInfo *aliasesUser
}

func (u User) Documentation() string {
	if u.PasswdInfo != nil {
		return strings.Join(
			[]string{
				fmt.Sprintf("%s (%s:%s)", u.PasswdInfo.Name, u.PasswdInfo.UID, u.PasswdInfo.GID),
				fmt.Sprintf("Home: `%s`", u.PasswdInfo.HomePath),
			},
			"\n",
		)
	}

	if u.AliasInfo != nil {
		return fmt.Sprintf("Defined on line %d", u.AliasInfo.DefinedOnLine+1)
	}

	return ""
}

// Returns a map of [username]user
// The username is normalized
func GetAvailableUserValues(
	i *indexes.AliasesIndexes,
) map[string]User {
	users := make(map[string]User)

	passwdUsers, err := common.FetchPasswdInfo()

	if err == nil {
		for _, info := range passwdUsers {
			key := indexes.NormalizeKey(info.Name)
			users[key] = User{
				PasswdInfo: &info,
			}
		}
	}

	if i != nil && i.Keys != nil {
		for name, key := range i.Keys {
			// Indexes keys are already normalized
			users[name] = User{
				AliasInfo: &aliasesUser{
					DefinedOnLine: key.Location.Start.Line,
				},
			}
		}
	}

	return users
}
