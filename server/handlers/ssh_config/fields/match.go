package fields

import (
	docvalues "config-lsp/doc-values"
	matchparser "config-lsp/handlers/ssh_config/match-parser"
)

var MatchExecField = docvalues.StringValue{}
var MatchLocalNetworkField = docvalues.IPAddressValue{
	AllowIPv4: true,
	AllowIPv6: true,
}
var MatchHostField = docvalues.StringValue{}
var MatchOriginalHostField = docvalues.StringValue{}
var MatchTypeTaggedField = docvalues.StringValue{}
var MatchUserField = docvalues.UserValue("", false)
var MatchTypeLocalUserField = docvalues.UserValue("", false)

var MatchValueFieldMap = map[matchparser.MatchCriteriaType]docvalues.DeprecatedValue{
	matchparser.MatchCriteriaTypeExec:         MatchExecField,
	matchparser.MatchCriteriaTypeLocalNetwork: MatchLocalNetworkField,
	matchparser.MatchCriteriaTypeHost:         MatchHostField,
	matchparser.MatchCriteriaTypeOriginalHost: MatchOriginalHostField,
	matchparser.MatchCriteriaTypeTagged:       MatchTypeTaggedField,
	matchparser.MatchCriteriaTypeUser:         MatchUserField,
	matchparser.MatchCriteriaTypeLocalUser:    MatchTypeLocalUserField,
}

var MatchAllArgumentAllowedPreviousOptions = map[matchparser.MatchCriteriaType]struct{}{
	matchparser.MatchCriteriaTypeCanonical: {},
	matchparser.MatchCriteriaTypeFinal:     {},
}

var MatchSingleOptionCriterias = map[matchparser.MatchCriteriaType]struct{}{
	matchparser.MatchCriteriaTypeAll:       {},
	matchparser.MatchCriteriaTypeCanonical: {},
	matchparser.MatchCriteriaTypeFinal:     {},
}
