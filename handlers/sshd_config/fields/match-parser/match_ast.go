package match_parser

import (
	"config-lsp/common"
	commonparser "config-lsp/common/parser"
)

type Match struct {
	Entries []*MatchEntry
}

type MatchCriteriaType string

const (
	MatchCriteriaTypeUser         MatchCriteriaType = "User"
	MatchCriteriaTypeGroup        MatchCriteriaType = "Group"
	MatchCriteriaTypeHost         MatchCriteriaType = "Host"
	MatchCriteriaTypeLocalAddress MatchCriteriaType = "LocalAddress"
	MatchCriteriaTypeLocalPort    MatchCriteriaType = "LocalPort"
	MatchCriteriaTypeRDomain      MatchCriteriaType = "RDomain"
	MatchCriteriaTypeAddress      MatchCriteriaType = "Address"
)

type MatchCriteria struct {
	common.LocationRange

	Type  MatchCriteriaType
	Value commonparser.ParsedString
}

type MatchSeparator struct {
	common.LocationRange
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
