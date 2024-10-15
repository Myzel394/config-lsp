package fields

var AllowedDuplicateOptions = map[NormalizedOptionName]struct{}{
	"allowgroups":   {},
	"allowusers":    {},
	"denygroups":    {},
	"denyusers":     {},
	"listenaddress": {},
	"match":         {},
	"port":          {},
}
