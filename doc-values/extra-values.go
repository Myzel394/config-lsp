package docvalues

import (
	"config-lsp/common"
	"config-lsp/utils"
	"regexp"
)

// UserValue returns a Value that fetches user names from /etc/passwd
// if `separatorForMultiple` is not empty, it will return an ArrayValue
func UserValue(separatorForMultiple string, enforceValues bool) Value {
	return CustomValue{
		FetchValue: func(context CustomValueContext) Value {
			infos, err := common.FetchPasswdInfo()

			if err != nil {
				return StringValue{}
			}

			enumValues := EnumValue{
				EnforceValues: enforceValues,
				Values: utils.Map(infos, func(info common.PasswdInfo) EnumString {
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

func GroupValue(separatorForMultiple string, enforceValues bool) Value {
	return CustomValue{
		FetchValue: func(context CustomValueContext) Value {
			infos, err := common.FetchGroupInfo()

			if err != nil {
				return StringValue{}
			}

			enumValues := EnumValue{
				EnforceValues: enforceValues,
				Values: utils.Map(infos, func(info common.GroupInfo) EnumString {
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
