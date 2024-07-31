package openssh

import "regexp"

var isJustDigitsPattern = regexp.MustCompile(`^\d+$`)
