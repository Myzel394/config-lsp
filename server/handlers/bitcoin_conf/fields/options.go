package fields

// Stores which options configure the chain. Those options are exclusive to each other
var ChainOptions = map[string]struct{}{
	"chain":    {},
	"test":     {},
	"testnet4": {},
	"signet":   {},
}

var AllowedDuplicateOptions = map[string]struct{}{
	"addnode":        {},
	"whitebind":      {},
	"whitelist":      {},
	"wallet":         {},
	"debug":          {},
	"debugexclude":   {},
	"signetseednode": {},
	"rpcallowip":     {},
	"rpcauth":        {},
	"ripcbind":       {},
}
