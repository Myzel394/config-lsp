package openssh

import (
	"config-lsp/common"
	"regexp"
)

func createOpenSSHConfigParser() common.SimpleConfigParser {
	return common.SimpleConfigParser{
		Lines: make(map[string]common.SimpleConfigLine),
		Options: common.SimpleConfigOptions{
			Separator:        *regexp.MustCompile(`(?m)^\s*(?P<OptionName>\w+)(?P<Separator>\s*)(?P<Value>.*)\s*$`),
			IgnorePattern:    *regexp.MustCompile(`^(?:#|\s*$)`),
			IdealSeparator:   " ",
			AvailableOptions: &Options,
		},
	}
}

var Parser = createOpenSSHConfigParser()
