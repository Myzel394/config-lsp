package fields

import (
	docvalues "config-lsp/doc-values"
	"regexp"
)

var AvailableSections map[string]string = map[string]string{
	"main": "Options for mainnet",
	"test": "Options for testnet3",
	"testnet4": "Options for testnet4",
	"signet": "Options for signet",
	"regtest": "Options for regtest",
}

var booleanValue = docvalues.EnumValue{
	Values: []docvalues.EnumString{
		docvalues.CreateEnumStringWithDoc("0", "false / off / disabled"),
		docvalues.CreateEnumStringWithDoc("1", "true / on / enabled"),
	},
	EnforceValues: true,
}

var Options = map[string]docvalues.DocumentationValue{
	"alertnotify": {
		Documentation: "Execute command when an alert is raised (%s in cmd is replaced by message)",
		Value: docvalues.StringValue{},
	},
	"allowignoredconf": {
		Documentation: "For backwards compatibility, treat an unused bitcoin.conf file in the datadir as a warning, not an error",
		Value: booleanValue,
	},
	"assumevalid": {
		Documentation: "If this block is in the chain assume that it and its ancestors are valid and potentially skip their script verification (0 to verify all, default: 00000000000000000001b658dd1120e82e66d2790811f89ede9742ada3ed6d77, testnet3: 00000000000003fc7967410ba2d0a8a8d50daedc318d43e8baf1a9782c236a57, testnet4: 0000000000003ed4f08dbdf6f7d6b271a6bcffce25675cb40aa9fa43179a89f3, signet: 000000895a110f46e59eb82bbc5bfb67fa314656009c295509c21b4999f5180a)",
		Value: docvalues.OrValue{
			Values: []docvalues.DeprecatedValue{
				docvalues.SingleEnumValue("0"),
				docvalues.RegexValue{
					Regex: *regexp.MustCompile(`^[0-9a-fA-F]{64}$`),
				},
			},
		},
	},
	"blockfilterindex": {
		Documentation: "Maintain an index of compact filters by block (default: 0, values: basic). If <type> is not supplied or if <type> = 1, indexes for all known types are enabled",
		Value: booleanValue,
	},
	"blocknotify": {
		Documentation: "Execute command when the best block changes (%s in cmd is replaced by block hash)",
		Value: docvalues.StringValue{},
	},
	"blockreconstructionextratxn": {
		Documentation: "Extra transactions to keep in memory for compact block reconstructions (default: 100)",
		Value: docvalues.NumberRangeValue(0, 999_999_999),
	},
	"blocksdir": {
		Documentation: "Specify directory to hold blocks subdirectory for *.dat files (default: <datadir>)",
		Value: docvalues.PathValue{
			RequiredType: docvalues.PathTypeDirectory,
		},
	},
	"blocksonly": {
		Documentation: "Whether to reject transactions from network peers. Disables automatic broadcast and rebroadcast of transactions, unless the source peer has the 'forcerelay' permission. RPC transactions are not affected. (default: 0)",
		Value: booleanValue,
	},
	"blocksxor": {
		Documentation: "Whether an XOR-key applies to blocksdir *.dat files. The created XOR-key will be zeros for an existing blocksdir or when `-blocksxor=0` is set, and random for a freshly initialized blocksdir. (default: 1)",
		Value: booleanValue,
	},
	"coinstatsindex": {
		Documentation: "Maintain coinstats index used by the gettxoutsetinfo RPC (default: 0)",
		Value: booleanValue,
	},
	// TODO: Add later
	"conf": {
		Documentation: "Specify path to read-only configuration file. Relative paths will be prefixed by datadir location (only useable from command line, not configuration file) (default: bitcoin.conf)",
		Value: docvalues.StringValue{},
	},
	"daemon": {
		Documentation: "Run in the background as a daemon and accept commands (default: 0)",
		Value: booleanValue,
	},
	"daemonwait": {
	Documentation: "Wait for initialization to be finished before exiting. This implies -daemon (default: 0)",
		Value: booleanValue,
	},
	"datadir": {
		Documentation: "Specify data directory",
		Value: docvalues.PathValue{
			RequiredType: docvalues.PathTypeDirectory,
		},
	},
	"dbcache": {
		Documentation: "Maximum database cache size <n> MiB (minimum 4, default: 450). Make sure you have enough RAM. In addition, unused memory allocated to the mempool is shared with this cache (see -maxmempool)",
		Value: docvalues.NumberRangeValue(4, 999_999_999),
	},
	"debuglogfile": {
		Documentation: "Specify location of debug log file (default: debug.log). Relative paths will be prefixed by a net-specific datadir location. Pass -nodebuglogfile to disable writing the log to a file.",
		Value: docvalues.StringValue{},
	},
	// TODO: Add later
	"includeconf": {
		Documentation: "Specify additional configuration file, relative to the -datadir path (only useable from configuration file, not command line)",
		Value: docvalues.PathValue{
			RequiredType: docvalues.PathTypeFile,
		},
	},
	"loadblock": {
		Documentation: "Imports blocks from external file on startup",
		Value: docvalues.PathValue{
			RequiredType: docvalues.PathTypeFile,
		},
	},
	"maxmempool": {
		Documentation: "Keep the transaction memory pool below <n> megabytes (default: 300)",
		Value: docvalues.NumberRangeValue(0, 999_999_999),
	},
	"maxorphantx": {
		Documentation: "Keep at most <n> unconnectable transactions in memory (default: 100)",
		Value: docvalues.NumberRangeValue(0, 999_999_999),
	},
	"mempoolexpiry": {
		Documentation: "Do not keep transactions in the mempool longer than <n> hours (default: 336)",
		Value: docvalues.NumberRangeValue(0, 999_999_999),
	},
	"par": {
		Documentation: "Set the number of script verification threads (0 = auto, up to 15, <0 = leave that many cores free, default: 0)",
		Value: docvalues.NumberRangeValue(0, 15),
	},
	"persistmempool": {
	Documentation: "Whether to save the mempool on shutdown and load on restart (default: 1)",
		Value: booleanValue,
	},
	"persistmempoolv1": {
	Documentation: "Whether a mempool.dat file created by -persistmempool or the savemempool RPC will be written in the legacy format (version 1) or the current format (version 2). This temporary option will be removed in the future. (default: 0)",
		Value: booleanValue,
	},
	"pid": {
		Documentation: "Specify pid file. Relative paths will be prefixed by a net-specific datadir location. (default: bitcoind.pid)",
		Value: docvalues.StringValue{},
	},
	"prune": {
		Documentation: "Reduce storage requirements by enabling pruning (deleting) of old blocks. This allows the pruneblockchain RPC to be called to delete specific blocks and enables automatic pruning of old blocks if a target size in MiB is provided. This mode is incompatible with -txindex. Warning: Reverting this setting requires re-downloading the entire blockchain. (default: 0 = disable pruning blocks, 1 = allow manual pruning via RPC, >=550 = automatically prune block files to stay under the specified target size in MiB)",
		Value: docvalues.OrValue{
			Values: []docvalues.DeprecatedValue{
				docvalues.EnumValue{
					Values: []docvalues.EnumString{
						docvalues.CreateEnumStringWithDoc("0", "disable pruning blocks"),
					},
				},
				docvalues.NumberRangeValue(1, 999_999_999),
			},
		},
	},
	"reindex": {
		Documentation: "If enabled, wipe chain state and block index, and rebuild them from blk*.dat files on disk. Also wipe and rebuild other optional indexes that are active. If an assumeutxo snapshot was loaded, its chainstate will be wiped as well. The snapshot can then be reloaded via RPC.",
		Value: booleanValue,
	},
	"reindex-chainstate": {
		Documentation: "If enabled, wipe chain state, and rebuild it from blk*.dat files on disk. If an assumeutxo snapshot was loaded, its chainstate will be wiped as well. The snapshot can then be reloaded via RPC.",
		Value: booleanValue,
	},
	"settings": {
		Documentation: "Specify path to dynamic settings data file. Can be disabled with -nosettings. File is written at runtime and not meant to be edited by users (use bitcoin.conf instead for custom settings). Relative paths will be prefixed by datadir location. (default: settings.json)",
		Value: docvalues.PathValue{
			RequiredType: docvalues.PathTypeFile,
		},
	},
	// TODO: Add later
	"shutdownnotify": {
		Documentation: "Execute command immediately before beginning shutdown. The need for shutdown may be urgent, so be careful not to delay it long (if the command doesn't require interaction with the server, consider having it fork into the background).",
		Value: docvalues.StringValue{},
	},
	"startupnotify": {
		Documentation: "Execute command on startup.",
		Value: docvalues.StringValue{},
	},
	"txindex": {
		Documentation: "Maintain a full transaction index, used by the getrawtransaction rpc call (default: 0)",
		Value: booleanValue,
	},
	"version": {
		Documentation: "Print version and exit",
		Value: booleanValue,
	},
	"addnode": {
		Documentation: "Add a node to connect to and attempt to keep the connection open (see the addnode RPC help for more info). This option can be specified multiple times to add multiple nodes; connections are limited to 8 at a time and are counted separately from the -maxconnections limit.",
		Value: docvalues.IPAddressValue{
			AllowIPv4: true,
			AllowIPv6: true,
			AllowRange: false,
		},
	},
	"asmap": {
		Documentation: "Specify asn mapping used for bucketing of the peers (default: ip_asn.map). Relative paths will be prefixed by the net-specific datadir location.",
		Value: docvalues.StringValue{},
	},
	"bantime": {
		Documentation: "Default duration (in seconds) of manually configured bans (default: 86400)",
		Value: docvalues.NumberRangeValue(0, 999_999_999),
	},
	// TODO: Add later
	"bind": {
		Documentation: `Bind to given address and always listen on it (default: 0.0.0.0). Use [host]:port notation for IPv6. Append =onion to tag any incoming connections to that address and port as incoming Tor connections (default: 127.0.0.1:8334=onion, testnet3: 127.0.0.1:18334=onion, testnet4: 127.0.0.1:48334=onion, signet: 127.0.0.1:38334=onion, regtest: 127.0.0.1=18445=onion)`,
		Value: docvalues.KeyValueAssignmentValue{
			Key: docvalues.IPAddressValue{
				AllowIPv4: true,
				AllowIPv6: true,
				AllowRange: false,
			},
			Value: docvalues.NumberRangeValue(1, 65535),
		},
	},
	"cjdnsreachable": {
		Documentation: "If set, then this host is configured for CJDNS (connecting to fc00::/8 addresses would lead us to the CJDNS network, see doc/cjdns.md) (default: 0)",
		Value: booleanValue,
	},
	"connect": {
		Documentation: "Connect only to the specified node; -noconnect disables automatic connections (the rules for this peer are the same as for -addnode). This option can be specified multiple times to connect to multiple nodes.",
		Value: docvalues.IPAddressValue{
			AllowIPv4: true,
			AllowIPv6: true,
			AllowRange: false,
		},
	},
	"discover": {
		Documentation: "Discover own IP addresses (default: 1 when listening and no -externalip or -proxy)",
		Value: booleanValue,
	},
	"dns": {
		Documentation: "Allow DNS lookups for -addnode, -seednode and -connect (default: 1)",
		Value: booleanValue,
	},
	"dnsseed": {
		Documentation: "Query for peer addresses via DNS lookup, if low on addresses (default: 1 unless -connect used or -maxconnections=0)",
		Value: booleanValue,
	},
	"externalip": {
		Documentation: "Specify your own public address",
		Value: docvalues.IPAddressValue{
			AllowIPv4: true,
			AllowIPv6: true,
			AllowRange: false,
		},
	},
	"fixedseeds": {
		Documentation: "Allow fixed seeds if DNS seeds don't provide peers (default: 1)",
		Value: booleanValue,
	},
	"forcednsseed": {
		Documentation: "Always query for peer addresses via DNS lookup (default: 0)",
		Value: booleanValue,
	},
	"i2pacceptincoming": {
		Documentation: "Whether to accept inbound I2P connections (default: 1). Ignored if -i2psam is not set. Listening for inbound I2P connections is done through the SAM proxy, not by binding to a local address and port.",
		Value: booleanValue,
	},
	"i2psam": {
		Documentation: "I2P SAM proxy to reach I2P peers and accept I2P connections (default: none)",
		Value: docvalues.IPAddressValue{
			AllowIPv4: true,
			AllowIPv6: true,
			AllowRange: false,
		},
	},
	"listen": {
		Documentation: "Accept connections from outside (default: 1 if no -proxy, -connect or -maxconnections=0)",
		Value: booleanValue,
	},
	"listenonion": {
		Documentation: "Automatically create Tor onion service (default: 1)",
		Value: booleanValue,
	},
	"maxconnections": {
		Documentation: "Maintain at most <n> automatic connections to peers (default: 125). This limit does not apply to connections manually added via -addnode or the addnode RPC, which have a separate limit of 8.",
		Value: docvalues.NumberRangeValue(0, 999_999_999),
	},
	"maxreceivebuffer": {
		Documentation: "Maximum per-connection receive buffer, <n>*1000 bytes (default: 5000)",
		Value: docvalues.NumberRangeValue(0, 999_999_999),
	},
	"maxsendbuffer": {
		Documentation: "Maximum per-connection memory usage for the send buffer, <n>*1000 bytes (default: 1000)",
		Value: docvalues.NumberRangeValue(0, 999_999_999),
	},
	// TODO: Add later
	"maxuploadtarget": {
	Documentation: "Tries to keep outbound traffic under the given target per 24h. Limit does not apply to peers with 'download' permission or blocks created within past week. 0 = no limit (default: 0M). Optional suffix units [k|K|m|M|g|G|t|T] (default: M). Lowercase is 1000 base while uppercase is 1024 base",
		Value: docvalues.StringValue{},
	},
	"natpmp": {
		Documentation: "Use PCP or NAT-PMP to map the listening port (default: 0)",
		Value: booleanValue,
	},
	"networkactive": {
		Documentation: "Enable all P2P network activity (default: 1). Can be changed by the setnetworkactive RPC command",
		Value: booleanValue,
	},
	// TODO: Add later
	"onion": {
		Documentation: "Use separate SOCKS5 proxy to reach peers via Tor onion services, set -noonion to disable (default: -proxy). May be a local file path prefixed with 'unix:'.",
		Value: docvalues.StringValue{},
	},
	"onlynet": {
		Documentation: "Make automatic outbound connections only to network <net> (ipv4, ipv6, onion, i2p, cjdns). Inbound and manual connections are not affected by this option. It can be specified multiple times to allow multiple networks.",
		Value: docvalues.EnumValue{
			Values: []docvalues.EnumString{
				docvalues.CreateEnumStringWithDoc("ipv4", "IPv4 network"),
				docvalues.CreateEnumStringWithDoc("ipv6", "IPv6 network"),
				docvalues.CreateEnumStringWithDoc("onion", "Tor onion network"),
				docvalues.CreateEnumStringWithDoc("i2p", "I2P network"),
				docvalues.CreateEnumStringWithDoc("cjdns", "CJDNS network"),
			},
			EnforceValues: true,
		},
	},
	"peerblockfilters": {
		Documentation: "Enable peer block filters (default: 1). This is a legacy feature that is not used by any known software. It is recommended to disable it.",
		Value: booleanValue,
	},
	"peerbloomfilters": {
		Documentation: "Enable peer bloom filters (default: 1). This is a legacy feature that is not used by any known software. It is recommended to disable it.",
		Value: booleanValue,
	},
	"port": {
		Documentation: "Listen for connections on <port> (default: 8333, testnet3: 18333, testnet4: 48333, signet: 38333, regtest: 18444). Use [host]:port notation for IPv6. This option is ignored if -connect is set.",
		Value: docvalues.NumberRangeValue(1, 65535),
	},
	// TODO: Add later
	"proxy": {
		Documentation: "Connect through SOCKS5 proxy (default: disabled). May be a local file path prefixed with 'unix:' if the proxy supports it.",
		Value: docvalues.OrValue{},
	},
	"proxyrandomize": {
		Documentation: "Randomize credentials for every proxy connection (default: 1). This enables Tor stream isolation (default: 1).",
		Value: booleanValue,
	},
	"seednode": {
		Documentation: "Connect to a node to retrieve peer addresses, and disconnect. This option can be specified multiple times to connect to multiple nodes. During startup, seednodes will be tried before dnsseeds.",
		Value: docvalues.IPAddressValue{
			AllowIPv4: true,
			AllowIPv6: true,
			AllowRange: false,
		},
	},
	"timeout": {
		Documentation: "Specify socket connection timeout in milliseconds. If an initial attempt to connect is unsuccessful after this amount of time, drop it (minimum: 1, default: 5000)",
		Value: docvalues.NumberRangeValue(1, 999_999_999),
	},
	// TODO: Add later
	"torcontrol": {
		Documentation: "Tor control host and port to use if onion listening enabled (default: 127.0.0.1:9051). If no port is specified, the default port of 9051 will be used.",
		Value: docvalues.StringValue{},
	},
	"torpassword": {
		Documentation: "Tor control port password (default: empty)",
		Value: docvalues.StringValue{},
	},
	"v2transport": {
		Documentation: "Support v2 transport (default: 1)",
		Value: booleanValue,
	},
	// TODO: Add later
	"whitebind": {
		Documentation: "Bind to the given address and add permission flags to the peers connecting to it. Use [host]:port notation for IPv6. Allowed permissions: bloomfilter (allow requesting BIP37 filtered blocks and transactions), noban (do not ban for misbehavior; implies download), forcerelay (relay transactions that are already in the mempool; implies relay), relay (relay even in -blocksonly mode, and unlimited transaction announcements), mempool (allow requesting BIP35 mempool contents), download (allow getheaders during IBD, no disconnect after maxuploadtarget limit), addr (responses to GETADDR avoid hitting the cache and contain random records with the most up-to-date info). Specify multiple permissions separated by commas (default: download,noban,mempool,relay). Can be specified multiple times.",
		Value: docvalues.StringValue{},
	},
	// TODO: Add later
	"whitelist": {
		Documentation: "Add permission flags to the peers using the given IP address (e.g. 1.2.3.4) or CIDR-notated network (e.g. 1.2.3.4/24). Uses the same permissions as -whitebind. Additional flags 'in' and 'out' control whether permissions apply to incoming connections and/or manual (default: incoming only). Can be specified multiple times.",
		Value: docvalues.StringValue{},
	},
	"addresstype": {
		Documentation: "What type of addresses to use ('legacy', 'p2sh-segwit', 'bech32', or 'bech32m', default: 'bech32')",
		Value: docvalues.EnumValue{
			Values: []docvalues.EnumString{
				docvalues.CreateEnumStringWithDoc("legacy", "Legacy addresses (P2PKH)"),
				docvalues.CreateEnumStringWithDoc("p2sh-segwit", "P2SH wrapped SegWit addresses (P2SH-P2WPKH)"),
				docvalues.CreateEnumStringWithDoc("bech32", "Native SegWit addresses (P2WPKH)"),
				docvalues.CreateEnumStringWithDoc("bech32m", "Native SegWit addresses (P2TR)"),
			},
			EnforceValues: true,
		},
	},
	"avoidpartialspends": {
		Documentation: "Group outputs by address, selecting many (possibly all) or none, instead of selecting on a per-output basis. Privacy is improved as addresses are mostly swept with fewer transactions and outputs are aggregated in clean change addresses. It may result in higher fees due to less optimal coin selection caused by this added limitation and possibly a larger-than-necessary number of inputs being used. Always enabled for wallets with 'avoid_reuse' enabled, otherwise default: 0.",
		Value: booleanValue,
	},
	"changetype": {
		Documentation: "What type of change to use ('legacy', 'p2sh-segwit', 'bech32', or 'bech32m'). Default is 'legacy' when -addresstype=legacy, else it is an implementation detail.",
		Value: docvalues.EnumValue{
			Values: []docvalues.EnumString{
				docvalues.CreateEnumStringWithDoc("legacy", "Legacy addresses (P2PKH)"),
				docvalues.CreateEnumStringWithDoc("p2sh-segwit", "P2SH wrapped SegWit addresses (P2SH-P2WPKH)"),
				docvalues.CreateEnumStringWithDoc("bech32", "Native SegWit addresses (P2WPKH)"),
				docvalues.CreateEnumStringWithDoc("bech32m", "Native SegWit addresses (P2TR)"),
			},
			EnforceValues: true,
		},
	},
	"consolidatefeerate": {
		Documentation: "The maximum feerate (in BTC/kvB) at which transaction building may use more inputs than strictly necessary so that the wallet's UTXO pool can be reduced (default: 0.0001).",
		Value: docvalues.NumberRangeValue(0, 999_999_999_999_999_999),
	},
	"disablewallet": {
		Documentation: "Do not load the wallet and disable wallet RPC calls",
		Value: booleanValue,
	},
	"discardfee": {
		Documentation: "The fee rate (in BTC/kvB) that indicates your tolerance for discarding change by adding it to the fee (default: 0.0001). Note: An output is discarded if it is dust at this rate, but we will always discard up to the dust relay fee and a discard fee above that is limited by the fee estimate for the longest target",
		Value: docvalues.NumberRangeValue(0, 999_999_999_999_999_999),
	},
	"fallbackfee": {
		Documentation: "A fee rate (in BTC/kvB) that will be used when fee estimation has insufficient data. 0 to entirely disable the fallbackfee feature. (default: 0.00)",
		Value: docvalues.NumberRangeValue(0, 999_999_999_999_999_999),
	},
	"keypool": {
		Documentation: "Set key pool size to <n> (default: 1000). Warning: Smaller sizes may increase the risk of losing funds when restoring from an old backup, if none of the addresses in the original keypool have been used.",
		Value: docvalues.NumberRangeValue(1, 999_999_999),
	},
	"maxapsfee": {
		Documentation: "Spend up to this amount in additional (absolute) fees (in BTC) if it allows the use of partial spend avoidance (default: 0.00)",
		Value: docvalues.NumberRangeValue(0, 999_999_999_999_999_999),
	},
	"mintxfee": {
		Documentation: "Fee rates (in BTC/kvB) smaller than this are considered zero fee for transaction creation (default: 0.00001)",
		Value: docvalues.NumberRangeValue(0, 999_999_999_999_999_999),
	},
	"paytxfee": {
		Documentation: "Fee rate (in BTC/kvB) to add to transactions you send (default: 0.00)",
		Value: docvalues.NumberRangeValue(0, 999_999_999_999_999_999),
	},
	"signer": {
		Documentation: "External signing tool, see doc/external-signer.md",
		Value: docvalues.StringValue{},
	},
	"spendzeroconfchange": {
		Documentation: "Spend unconfirmed change when sending transactions (default: 1)",
		Value: booleanValue,
	},
	"txconfirmtarget": {
		Documentation: "If paytxfee is not set, include enough fee so transactions begin confirmation on average within n blocks (default: 6)",
		Value: docvalues.NumberRangeValue(1, 999_999_999),
	},
	"wallet": {
		Documentation: "Specify wallet path to load at startup. Can be used multiple times to load multiple wallets. Path is to a directory containing wallet data and log files. If the path is not absolute, it is interpreted relative to <walletdir>. This only loads existing wallets and does not create new ones. For backwards compatibility this also loads the default wallet.",
		Value: docvalues.PathValue{
			RequiredType: docvalues.PathTypeDirectory,
		},
	},
	"walletbroadcast": {
		Documentation: "Make the wallet broadcast transactions (default: 1)",
		Value: booleanValue,
	},
	"walletdir": {
		Documentation: "Specify directory to hold wallets (default: <datadir>/wallets if it exists, otherwise <datadir>)",
		Value: docvalues.PathValue{
			RequiredType: docvalues.PathTypeDirectory,
		},
	},
	"walletnotify": {
		Documentation: "Execute command when a wallet transaction changes. %s in cmd is replaced by TxID, %w is replaced by wallet name, %b is replaced by the hash of the block including the transaction (set to 'unconfirmed' if the transaction is not included) and %h is replaced by the block height (-1 if not included). %w is not currently implemented on windows. On systems where %w is supported, it should NOT be quoted because this would break shell escaping used to invoke the command.",
		Value: docvalues.StringValue{},
	},
	"walletrbf": {
		Documentation: "Whether to use Replace-By-Fee (RBF) for wallet transactions (default: 1). If set to 0, RBF is not used and transactions are not marked as replaceable.",
		Value: booleanValue,
	},
	"zmqpubhashblock": {
		Documentation: "Enable publish hash block in <address>",
		Value: docvalues.StringValue{},
	},
	"zmqpubhashblockhwm": {
		Documentation: "Set publish hash block outbound message high water mark (default: 1000)",
		Value: docvalues.NumberRangeValue(1, 999_999_999),
	},
	"zmqpubhashtx": {
		Documentation: "Enable publish hash transaction in <address>",
		Value: docvalues.StringValue{},
	},
	"zmqpubhashtxhwm": {
		Documentation: "Set publish hash transaction outbound message high water mark (default: 1000)",
		Value: docvalues.NumberRangeValue(1, 999_999_999),
	},
	"zmqpubrawblock": {
		Documentation: "Enable publish raw block in <address>",
		Value: docvalues.StringValue{},
	},
	"zmqpubrawblockhwm": {
		Documentation: "Set publish raw block outbound message high water mark (default: 1000)",
		Value: docvalues.NumberRangeValue(1, 999_999_999),
	},
	"zmqpubrawtx": {
		Documentation: "Enable publish raw transaction in <address>",
		Value: docvalues.StringValue{},
	},
	"zmqpubrawtxhwm": {
		Documentation: "Set publish raw transaction outbound message high water mark (default: 1000)",
		Value: docvalues.NumberRangeValue(1, 999_999_999),
	},
	"zmqpubsequence": {
		Documentation: "Enable publish hash block and tx sequence in <address>",
		Value: docvalues.StringValue{},
	},
	"zmqpubsequencehwm": {
		Documentation: "Set publish hash sequence message high water mark (default: 1000)",
		Value: docvalues.NumberRangeValue(1, 999_999_999),
	},
	"debug": {
		Documentation: "Output debug and trace logging (default: -nodebug, supplying <category> is optional). If <category> is not supplied or if <category> is 1 or 'all', output all debug logging. If <category> is 0 or 'none', any other categories are ignored. Other valid values for <category> are: addrman, bench, blockstorage, cmpctblock, coindb, estimatefee, http, i2p, ipc, leveldb, libevent, mempool, mempoolrej, net, proxy, prune, qt, rand, reindex, rpc, scan, selectcoins, tor, txpackages, txreconciliation, validation, walletdb, zmq. This option can be specified multiple times to output multiple categories.",
		Value: docvalues.OrValue{
			Values: []docvalues.DeprecatedValue{
				docvalues.EnumValue{
					Values: []docvalues.EnumString{
						docvalues.CreateEnumStringWithDoc("0", "no debug logging"),
						docvalues.CreateEnumStringWithDoc("1", "all debug logging"),
						docvalues.CreateEnumStringWithDoc("all", "all debug logging"),
						docvalues.CreateEnumStringWithDoc("none", "no debug logging"),
					},
				},
				docvalues.StringValue{},
			},
		},
	},
	"debugexclude": {
		Documentation: "Exclude debug and trace logging for a category. Can be used in conjunction with -debug=1 to output debug and trace logging for all categories except the specified category. This option can be specified multiple times to exclude multiple categories. This takes priority over '-debug'",
		Value: docvalues.StringValue{},
	},
	"logips": {
		Documentation: "Include IP addresses in debug output (default: 0)",
		Value: booleanValue,
	},
	"loglevelalways": {
		Documentation: "Always prepend a category and level (default: 0)",
		Value: booleanValue,
	},
	"logsourcelocations": {
		Documentation: "Prepend debug output with name of the originating source location (source file, line number and function name) (default: 0)",
		Value: booleanValue,
	},
	"logthreadnames": {
		Documentation: "Prepend debug output with name of the originating thread (default: 0)",
		Value: booleanValue,
	},
	"logtimestamps": {
		Documentation: "Prepend debug output with timestamp (default: 1)",
		Value: booleanValue,
	},
	"maxtxfee": {
		Documentation: "Maximum total fees (in BTC) to use in a single wallet transaction; setting this too low may abort large transactions (default: 0.10)",
		Value: docvalues.NumberRangeValue(0, 999_999_999_999_999_999),
	},
	"printtoconsole": {
		Documentation: "Send trace/debug info to console (default: 1 when no -daemon. To disable logging to file, set -nodebuglogfile)",
		Value: booleanValue,
	},
	"shrinkdebugfile": {
		Documentation: "Shrink debug.log file on client startup (default: 1 when no -debug)",
		Value: booleanValue,
	},
	"uacomment": {
		Documentation: "Append comment to the user agent string",
		Value: docvalues.StringValue{},
	},
	"chain": {
		Documentation: "Use the chain <chain> (default: main). Allowed values: main, test, testnet4, signet, regtest",
		Value: docvalues.EnumValue{
			Values: []docvalues.EnumString{
				docvalues.CreateEnumStringWithDoc("main", "Main chain"),
				docvalues.CreateEnumStringWithDoc("test", "Test chain (deprecated, use testnet4)"),
				docvalues.CreateEnumStringWithDoc("testnet4", "Testnet4 chain"),
				docvalues.CreateEnumStringWithDoc("signet", "Signet chain"),
				docvalues.CreateEnumStringWithDoc("regtest", "Regtest chain"),
			},
			EnforceValues: true,
		},
	},
	"signet": {
		Documentation: "Use the signet chain. Equivalent to -chain=signet. Note that the network is defined by the -signetchallenge parameter",
		Value: booleanValue,
	},
	"signetchallenge": {
		Documentation: "Blocks must satisfy the given script to be considered valid (only for signet networks; defaults to the global default signet test network challenge)",
		Value: docvalues.StringValue{},
	},
	// TODO: Add later
	"signetseednode": {
		Documentation: "Specify a seed node for the signet network, in the hostname[:port] format, e.g. sig.net:1234 (may be used multiple times to specify multiple seed nodes; defaults to the global default signet test network seed node(s))",
		Value: docvalues.StringValue{},
	},
	"testnet": {
		Documentation: "Use the testnet3 chain. Equivalent to -chain=test. Support for testnet3 is deprecated and will be removed in an upcoming release. Consider moving to testnet4 now by using -testnet4.",
		Value: booleanValue,
	},
	"testnet4": {
		Documentation: "Use the testnet4 chain. Equivalent to -chain=testnet4.",
		Value: booleanValue,
	},
	"bytespersigop": {
	Documentation: "Equivalent bytes per sigop in transactions for relay and mining (default: 20)",
		Value: docvalues.NumberRangeValue(1, 999_999_999),
	},
	"datacarrier": {
	Documentation: "Relay and mine data carrier transactions (default: 1)",
		Value: booleanValue,
	},
	"datacarriersize": {
	Documentation: "Relay and mine transactions whose data-carrying raw scriptPubKey is of this size or less (default: 83)",
		Value: docvalues.NumberRangeValue(1, 999_999_999),
	},
	"minrelaytxfee": {
		Documentation: "Fees (in BTC/kvB) smaller than this are considered zero fee for relaying, mining and transaction creation (default: 0.00001)",
		Value: docvalues.NumberRangeValue(0, 999_999_999_999_999_999),
	},
	"permitbaremultisig": {
		Documentation: "Relay transactions creating non-P2SH multisig outputs (default: 1)",
		Value: booleanValue,
	},
	"whitelistforcerelay": {
		Documentation: "Add 'forcerelay' permission to whitelisted peers with default permissions. This will relay transactions even if the transactions were already in the mempool. (default: 0)",
		Value: booleanValue,
	},
	"whitelistrelay": {
		Documentation: "Add 'relay' permission to whitelisted peers with default permissions. This will accept relayed transactions even when not relaying transactions (default: 1)",
		Value: booleanValue,
	},
	"blockmaxweight": {
		Documentation: "Set maximum BIP141 block weight (default: 4000000)",
		Value: docvalues.NumberRangeValue(1, 999_999_999),
	},
	"blockmintxfee": {
		Documentation: "Set lowest fee rate (in BTC/kvB) for transactions to be included in block creation. (default: 0.00001)",
		Value: docvalues.NumberRangeValue(0, 999_999_999_999_999_999),
	},
	"blockreservedweight": {
		Documentation: "Reserve space for the fixed-size block header plus the largest coinbase transaction the mining software may add to the block. (default: 8000).",
		Value: docvalues.NumberRangeValue(1, 999_999_999),
	},
	"rest": {
		Documentation: "Accept public REST requests (default: 0)",
		Value: booleanValue,
	},
	// TODO: Add later
	"rpcallowip": {
		Documentation: "",
		Value: docvalues.StringValue{},
	},
	"rpcauth": {
		Documentation: "Username and HMAC-SHA-256 hashed password for JSON-RPC connections. The field <userpw> comes in the format: <USERNAME>:<SALT>$<HASH>. A canonical python script is included in share/rpcauth. The client then connects normally using the rpcuser=<USERNAME>/rpcpassword=<PASSWORD> pair of arguments. This option can be specified multiple times",
		Value: docvalues.StringValue{},
	},
	// TODO: Add later
	"rpcbind": {
		Documentation: "Bind to given address to listen for JSON-RPC connections. Do not expose the RPC server to untrusted networks such as the public internet! This option is ignored unless -rpcallowip is also passed. Port is optional and overrides -rpcport. Use [host]:port notation for IPv6. This option can be specified multiple times (default: 127.0.0.1 and ::1 i.e., localhost)",
		Value: docvalues.StringValue{},
	},
	"rpccookiefile": {
		Documentation: "Location of the auth cookie. Relative paths will be prefixed by a net-specific datadir location. (default: data dir)",
		Value: docvalues.PathValue{
			RequiredType: docvalues.PathTypeFile,
		},
	},
	"rpccookieperms": {
		Documentation: "Set permissions on the RPC auth cookie file so that it is readable by [owner|group|all] (default: owner [via umask 0077])",
		Value: docvalues.EnumValue{
			Values: []docvalues.EnumString{
				docvalues.CreateEnumStringWithDoc("owner", "Readable by owner only"),
				docvalues.CreateEnumStringWithDoc("group", "Readable by group only"),
				docvalues.CreateEnumStringWithDoc("all", "Readable by all"),
			},
			EnforceValues: true,
		},
	},
	"rpcpassword": {
		Documentation: "Password for JSON-RPC connections",
		Value: docvalues.StringValue{},
	},
	"rpcport": {
		Documentation: "Listen for JSON-RPC connections on <port> (default: 8332, testnet3: 18332, testnet4: 48332, signet: 38332, regtest: 18443)",
		Value: docvalues.NumberRangeValue(1, 65535),
	},
	"rpcthreads": {
		Documentation: "Set the number of threads to service RPC calls (default: 16)",
		Value: docvalues.NumberRangeValue(1, 999_999_999),
	},
	"rpcuser": {
		Documentation: "Username for JSON-RPC connections",
		Value: docvalues.StringValue{},
	},
	// TODO: Add later
	"rpcwhitelist": {
		Documentation: "Set a whitelist to filter incoming RPC calls for a specific user. The field <whitelist> comes in the format: <USERNAME>:<rpc 1>,<rpc 2>,...,<rpc n>. If multiple whitelists are set for a given user, they are set-intersected. See -rpcwhitelistdefault documentation for information on default whitelist behavior.",
		Value: docvalues.StringValue{},
	},
	"rpcwhitelistdefault": {
	Documentation: "Sets default behavior for rpc whitelisting. Unless rpcwhitelistdefault is set to 0, if any -rpcwhitelist is set, the rpc server acts as if all rpc users are subject to empty-unless-otherwise-specified whitelists. If rpcwhitelistdefault is set to 1 and no -rpcwhitelist is set, rpc server acts as if all rpc users are subject to empty whitelists.",
		Value: booleanValue,
	},
	"server": {
		Documentation: "Accept command line and JSON-RPC commands",
		Value: booleanValue,
	},
}

