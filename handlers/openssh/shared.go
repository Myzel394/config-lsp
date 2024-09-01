package openssh

import (
	"regexp"
)

func createOpenSSHConfigParser() SimpleConfigParser {
	return SimpleConfigParser{
		Lines: make(map[string]SimpleConfigLine),
		Options: SimpleConfigOptions{
			Separator:        *regexp.MustCompile(`(?m)^\s*(?P<OptionName>\w+)(?P<Separator>\s*)(?P<Value>.*)\s*$`),
			IgnorePattern:    *regexp.MustCompile(`^(?:#|\s*$)`),
			IdealSeparator:   " ",
			AvailableOptions: &Options,
		},
	}
}

var Parser = createOpenSSHConfigParser()
