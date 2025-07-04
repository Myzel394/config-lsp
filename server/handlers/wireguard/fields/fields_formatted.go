package fields

var AllOptionsFormatted = map[NormalizedName]string{
	// Interface
	"address":    "Address",
	"listenport": "ListenPort",
	"privatekey": "PrivateKey",
	"dns":        "DNS",
	"table":      "Table",
	"mtu":        "MTU",
	"preup":      "PreUp",
	"postup":     "PostUp",
	"predown":    "PreDown",
	"postdown":   "PostDown",
	"fwmark":     "FwMark",

	// Peer Options
	"endpoint":            "Endpoint",
	"allowedips":          "AllowedIPs",
	"publickey":           "PublicKey",
	"persistentkeepalive": "PersistentKeepalive",
	"presharedkey":        "PresharedKey",
}
