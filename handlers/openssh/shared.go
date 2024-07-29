package openssh

import (
	"config-lsp/common"
	"regexp"
)

func createOpenSSHConfigParser() common.SimpleConfigParser {
	pattern, err := regexp.Compile(`^(?:#|\s*$)`)

	if err != nil {
		panic(err)
	}

	return common.SimpleConfigParser{
		Lines: make(map[string]common.SimpleConfigLine),
		Options: common.SimpleConfigOptions{
			Separator:        " ",
			IgnorePattern:    *pattern,
			AvailableOptions: &Options,
		},
	}
}

var Parser = createOpenSSHConfigParser()
