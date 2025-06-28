package utils

import (
	"errors"
	"regexp"
	"strings"
)

type IpTableActionEnum uint8

const (
	IpTableActionAppend IpTableActionEnum = iota
	IpTableActionInsert
	IpTableActionCheck
	IpTableActionDelete
)

type IpTableRule struct {
	Action IpTableActionEnum
	// Position in the rule list
	// This is used to make the deletion command as similar to the original rule as possible
	ActionIndex uint32
	Command     string
}

func (i IpTableRule) InvertAction() IpTableRule {
	actionMap := map[IpTableActionEnum]IpTableActionEnum{
		IpTableActionAppend: IpTableActionDelete,
		IpTableActionInsert: IpTableActionDelete,
		IpTableActionCheck:  IpTableActionDelete,
		IpTableActionDelete: IpTableActionAppend,
	}

	return IpTableRule{
		Action:      actionMap[i.Action],
		ActionIndex: i.ActionIndex,
		Command:     i.Command,
	}
}

func (i IpTableRule) String() string {
	actionStr := ""

	switch i.Action {
	case IpTableActionAppend:
		actionStr = "-A"
	case IpTableActionInsert:
		actionStr = "-I"
	case IpTableActionCheck:
		actionStr = "-C"
	case IpTableActionDelete:
		actionStr = "-D"
	}

	ruleRaw := i.Command[:i.ActionIndex] + " " + actionStr + " " + i.Command[i.ActionIndex:]

	return strings.ReplaceAll(ruleRaw, "  ", " ")
}

var rulePattern = regexp.MustCompile(`\B(-I|-D|-C|-A|--insert|--append|--check|--delete)\b`)
var actionMap = map[string]IpTableActionEnum{
	"-I":       IpTableActionInsert,
	"--insert": IpTableActionInsert,
	"-D":       IpTableActionDelete,
	"--delete": IpTableActionDelete,
	"-C":       IpTableActionCheck,
	"--check":  IpTableActionCheck,
	"-A":       IpTableActionAppend,
	"--append": IpTableActionAppend,
}

// A very simple parser for iptable rules.
// Better approach: Use something like antlr or ast to parse the rules.
func ParseIpTableRule(rule string) (*IpTableRule, error) {
	matches := rulePattern.FindIndex([]byte(rule))

	if len(matches) != 2 {
		return nil, errors.New("Invalid iptable rule. Couldn't find action")
	}

	actionStart := matches[0]
	actionEnd := matches[1]

	action := rule[actionStart:actionEnd]

	actionEnum, found := actionMap[action]

	if !found {
		return nil, errors.New("Invalid iptable rule. Unknown action: " + action)
	}

	// Remove the action from the rule
	newRule := rule[:actionStart] + rule[actionEnd:]

	return &IpTableRule{
		Action:      actionEnum,
		ActionIndex: uint32(actionStart),
		Command:     newRule,
	}, nil
}
