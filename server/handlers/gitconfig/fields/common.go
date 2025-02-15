package fields

import "strings"

type NormalizedOptionName string

func CreateNormalizedName(s string) NormalizedOptionName {
	return NormalizedOptionName(strings.ToLower(s))
}
