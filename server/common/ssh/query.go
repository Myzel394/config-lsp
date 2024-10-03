package ssh

import (
	"config-lsp/doc-values"
	"config-lsp/utils"
	"os/exec"
	"strings"
)

var _cachedQueries map[string][]docvalues.EnumString = make(map[string][]docvalues.EnumString)

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
