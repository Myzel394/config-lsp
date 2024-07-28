package handlers

import (
	"os/exec"
	"strings"
)

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
