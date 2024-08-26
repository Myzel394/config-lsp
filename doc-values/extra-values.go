package docvalues

import (
	"config-lsp/utils"
	"os"
	"regexp"
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
		FetchValue: func(context CustomValueContext) Value {
			infos, err := fetchPasswdInfo()

			if err != nil {
				return StringValue{}
			}

			enumValues := EnumValue{
				EnforceValues: enforceValues,
				Values: utils.Map(infos, func(info passwdInfo) EnumString {
					return CreateEnumString(info.Name)
				}),
			}

			if separatorForMultiple == "" {
				return enumValues
			} else {
				return ArrayValue{
					DuplicatesExtractor: &SimpleDuplicatesExtractor,
					SubValue:            enumValues,
					Separator:           separatorForMultiple,
				}
			}
		},
	}
}

type groupInfo struct {
	Name string
	GID  string
}

var _cachedGroupInfo []groupInfo

func fetchGroupInfo() ([]groupInfo, error) {
	if len(_cachedGroupInfo) > 0 {
		return _cachedGroupInfo, nil
	}

	readBytes, err := os.ReadFile("/etc/group")

	if err != nil {
		return []groupInfo{}, err
	}

	lines := strings.Split(string(readBytes), "\n")
	infos := make([]groupInfo, 0)

	for _, line := range lines {
		splitted := strings.Split(line, ":")

		if len(splitted) < 3 {
			continue
		}

		info := groupInfo{
			Name: splitted[0],
			GID:  splitted[2],
		}

		infos = append(infos, info)
	}

	_cachedGroupInfo = infos

	return infos, nil
}

func GroupValue(separatorForMultiple string, enforceValues bool) Value {
	return CustomValue{
		FetchValue: func(context CustomValueContext) Value {
			infos, err := fetchGroupInfo()

			if err != nil {
				return StringValue{}
			}

			enumValues := EnumValue{
				EnforceValues: enforceValues,
				Values: utils.Map(infos, func(info groupInfo) EnumString {
					return CreateEnumString(info.Name)
				}),
			}

			if separatorForMultiple == "" {
				return enumValues
			} else {
				return ArrayValue{
					DuplicatesExtractor: &SimpleDuplicatesExtractor,
					SubValue:            enumValues,
					Separator:           separatorForMultiple,
				}
			}
		},
	}
}

func PositiveNumberValue() Value {
	zero := 0
	return NumberValue{
		Min: &zero,
	}
}

func MaskValue() Value {
	min := 0
	max := 777
	return NumberValue{Min: &min, Max: &max}
}

func SingleEnumValue(value string) EnumValue {
	return EnumValue{
		EnforceValues: true,
		Values: []EnumString{
			CreateEnumString(value),
		},
	}
}

func DomainValue() Value {
	return RegexValue{
		Regex: *regexp.MustCompile(`^.+?\..+$`),
	}
}
