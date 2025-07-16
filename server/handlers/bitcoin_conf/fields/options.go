package fields

// Stores which options configure the chain. Those options are exclusive to each other
var ChainOptions = map[string]struct{}{
	"chain":    {},
	"test":     {},
	"testnet4": {},
	"signet":   {},
}

// Available sections to configure chain-specific options
var AvailableSections map[string]string = map[string]string{
	"main":     "Options for mainnet",
	"test":     "Options for testnet3",
	"testnet4": "Options for testnet4",
	"signet":   "Options for signet",
	"regtest":  "Options for regtest",
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
	"rpcbind":       {},
}
