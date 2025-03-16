package fields

import "strings"

type NormalizedName string

func CreateNormalizedName(s string) NormalizedName {
	return NormalizedName(strings.ToLower(s))
}
