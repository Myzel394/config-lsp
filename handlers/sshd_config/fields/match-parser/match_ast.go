package match_parser

import (
	"config-lsp/common"
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

	Type MatchCriteriaType
}

type MatchEntry struct {
	common.LocationRange
	Value string

	Criteria MatchCriteriaType
	Values   []*MatchValue
}

type MatchValue struct {
	common.LocationRange
	Value string
}
