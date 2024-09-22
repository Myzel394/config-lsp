package matchparser

import (
	"config-lsp/common"
	commonparser "config-lsp/common/parser"
)

type Match struct {
	Entries []*MatchEntry
}

type MatchCriteriaType string

const (
	MatchCriteriaTypeCanonical    MatchCriteriaType = "canonical"
	MatchCriteriaTypeFinal        MatchCriteriaType = "final"
	MatchCriteriaTypeExec         MatchCriteriaType = "exec"
	MatchCriteriaTypeLocalNetwork MatchCriteriaType = "localnetwork"
	MatchCriteriaTypeHost         MatchCriteriaType = "host"
	MatchCriteriaTypeOriginalHost MatchCriteriaType = "originalhost"
	MatchCriteriaTypeTagged       MatchCriteriaType = "tagged"
	MatchCriteriaTypeUser         MatchCriteriaType = "user"
	MatchCriteriaTypeLocalUser    MatchCriteriaType = "localuser"
)

type MatchCriteria struct {
	common.LocationRange

	Type  MatchCriteriaType
	Value commonparser.ParsedString
}

type MatchSeparator struct {
	common.LocationRange
	Value commonparser.ParsedString
}

type MatchValues struct {
	common.LocationRange

	Values []*MatchValue
}

type MatchEntry struct {
	common.LocationRange
	Value commonparser.ParsedString

	Criteria  MatchCriteria
	Separator *MatchSeparator
	Values    *MatchValues
}

type MatchValue struct {
	common.LocationRange
	Value commonparser.ParsedString
}
