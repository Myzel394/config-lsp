package common

import (
	"os"
	"strings"
)

type passwdInfo struct {
	Name     string
	UID      string
	GID      string
	HomePath string
}

var _cachedPasswdInfo []passwdInfo

func fetchPasswdInfo() ([]passwdInfo, error) {
	if len(_cachedPasswdInfo) > 0 {
		return _cachedPasswdInfo, nil
	}

	readBytes, err := os.ReadFile("/etc/passwd")

	if err != nil {
		return []passwdInfo{}, err
	}

	lines := strings.Split(string(readBytes), "\n")
	infos := make([]passwdInfo, 0)

	for _, line := range lines {
		splitted := strings.Split(line, ":")

		if len(splitted) < 6 {
			continue
		}

		info := passwdInfo{
			Name:     splitted[0],
			UID:      splitted[2],
			GID:      splitted[3],
			HomePath: splitted[5],
		}

		infos = append(infos, info)
	}

	_cachedPasswdInfo = infos

	return infos, nil
}

// UserValue returns a Value that fetches user names from /etc/passwd
// if `separatorForMultiple` is not empty, it will return an ArrayValue
func UserValue(separatorForMultiple string, enforceValues bool) Value {
	return CustomValue{
		FetchValue: func() Value {
			infos, err := fetchPasswdInfo()

			if err != nil {
				return StringValue{}
			}

			enumValues := EnumValue{
				EnforceValues: enforceValues,
				Values: Map(infos, func(info passwdInfo) string {
					return info.Name
				}),
			}

			if separatorForMultiple == "" {
				return enumValues
			} else {
				return ArrayValue{
					DuplicatesExtractor: &SimpleDuplicatesExtractor,
					SubValue:        enumValues,
					Separator:       separatorForMultiple,
				}
			}
		},
	}
}
