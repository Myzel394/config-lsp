package openssh

import (
	"config-lsp/common"
	"os/exec"
	"strings"
)

var BooleanEnumValue = common.EnumValue{
	EnforceValues: true,
	Values:        []string{"yes", "no"},
}

var plusMinuxCaretPrefixes = []common.Prefix{
	{
		Prefix:  "+",
		Meaning: "Append to the default set",
	},
	{
		Prefix:  "-",
		Meaning: "Remove from the default set",
	},
	{
		Prefix:  "^",
		Meaning: "Place at the head of the default set",
	},
}

var ChannelTimeoutExtractor = common.ExtractKeyDuplicatesExtractor("=")

func PrefixPlusMinusCaret(values []string) common.PrefixWithMeaningValue {
	return common.PrefixWithMeaningValue{
		Prefixes: []common.Prefix{
			{
				Prefix:  "+",
				Meaning: "Append to the default set",
			},
			{
				Prefix:  "-",
				Meaning: "Remove from the default set",
			},
			{
				Prefix:  "^",
				Meaning: "Place at the head of the default set",
			},
		},
		SubValue: common.ArrayValue{
			Separator:       ",",
			DuplicatesExtractor: &common.SimpleDuplicatesExtractor,
			SubValue: common.EnumValue{
				Values: values,
			},
		},
	}
}

var _cachedQueries map[string][]string = make(map[string][]string)

func queryValues(query string) ([]string, error) {
	cmd := exec.Command("ssh", "-Q", query)

	output, err := cmd.Output()

	if err != nil {
		return []string{}, err
	}

	return strings.Split(string(output), "\n"), nil
}

func QueryOpenSSHOptions(
	query string,
) ([]string, error) {
	var availableQueries []string
	key := query

	if _cachedQueries[key] != nil && len(_cachedQueries[key]) > 0 {
		return _cachedQueries[key], nil
	} else {
		availableQueries, err := queryValues(query)

		if err != nil {
			return []string{}, err
		}

		_cachedQueries[key] = availableQueries
	}

	return availableQueries, nil
}

