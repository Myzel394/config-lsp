package openssh

import (
	docvalues "config-lsp/doc-values"
	"os/exec"
	"strings"
)

var BooleanEnumValue = docvalues.EnumValue{
	EnforceValues: true,
	Values:        []string{"yes", "no"},
}

var plusMinuxCaretPrefixes = []docvalues.Prefix{
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

var ChannelTimeoutExtractor = docvalues.ExtractKeyDuplicatesExtractor("=")
var SetEnvExtractor = docvalues.ExtractKeyDuplicatesExtractor("=")

func PrefixPlusMinusCaret(values []string) docvalues.PrefixWithMeaningValue {
	return docvalues.PrefixWithMeaningValue{
		Prefixes: []docvalues.Prefix{
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
		SubValue: docvalues.ArrayValue{
			Separator:           ",",
			DuplicatesExtractor: &docvalues.SimpleDuplicatesExtractor,
			SubValue: docvalues.EnumValue{
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
