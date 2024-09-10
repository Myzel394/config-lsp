package fields

import (
	docvalues "config-lsp/doc-values"
	"config-lsp/utils"
	"os/exec"
	"regexp"
	"strings"
)

var isJustDigitsPattern = regexp.MustCompile(`^\d+$`)

var booleanEnumValue = docvalues.EnumValue{
	EnforceValues: true,
	Values: []docvalues.EnumString{
		docvalues.CreateEnumString("yes"),
		docvalues.CreateEnumString("no"),
	},
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

var channelTimeoutExtractor = docvalues.ExtractKeyDuplicatesExtractor("=")
var setEnvExtractor = docvalues.ExtractKeyDuplicatesExtractor("=")

func prefixPlusMinusCaret(values []docvalues.EnumString) docvalues.PrefixWithMeaningValue {
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

var _cachedQueries map[string][]docvalues.EnumString = make(map[string][]docvalues.EnumString)

func queryValues(query string) ([]string, error) {
	cmd := exec.Command("ssh", "-Q", query)

	output, err := cmd.Output()

	if err != nil {
		return []string{}, err
	}

	return strings.Split(string(output), "\n"), nil
}

func queryOpenSSHOptions(
	query string,
) ([]docvalues.EnumString, error) {
	var availableQueries []docvalues.EnumString
	key := query

	if _cachedQueries[key] != nil && len(_cachedQueries[key]) > 0 {
		return _cachedQueries[key], nil
	} else {
		availableRawQueries, err := queryValues(query)
		availableQueries = utils.Map(availableRawQueries, docvalues.CreateEnumString)

		if err != nil {
			return []docvalues.EnumString{}, err
		}

		_cachedQueries[key] = availableQueries
	}

	return availableQueries, nil
}
