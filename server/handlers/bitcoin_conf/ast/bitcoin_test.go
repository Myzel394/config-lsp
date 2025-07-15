package ast

import "testing"

func SimpleBTCConfigExampleTest(t *testing.T) {
	content := `
addnode=10.0.0.1
chain=main
`

	btcConfig := NewBTCConfig()
	errors := btcConfig.Parse(content)

	if len(errors) > 0 {
		t.Errorf("Expected no errors, got %d errors: %v", len(errors), errors)
	}
}

// https://github.com/Ride-The-Lightning/RTL/blob/847923533e23685d2f4e2fd2f9e07f5a44ea5198/docker/bitcoind/bitcoin.conf
func BTCConfigExample1Test(t *testing.T) {
	content := `
daemon=0
printtoconsole=1
`

	btcConfig := NewBTCConfig()
	errors := btcConfig.Parse(content)

	if len(errors) > 0 {
		t.Errorf("Expected no errors, got %d errors: %v", len(errors), errors)
	}
}

// From https://github.com/OmniLayer/OmniJ/blob/22cab42ae3acc910b07755c1b0ac790b9a83ef80/bitcoin.conf
func BTCConfigExample2Test(t *testing.T) {
	content := `
server=1
rpcauth=bitcoinrpc:b0d928012ce362e4d47345a346595dc6$8f0b2dde3617aa4aeb7aaf15eaaa8f3f085b7963e30ea973743f32e585552b7a
rpcallowip=127.0.0.1
txindex=1
debug=1
logtimestamps=1
omniseedblockfilter=0

[regtest]
rpcport=18443
`

	btcConfig := NewBTCConfig()
	errors := btcConfig.Parse(content)

	if len(errors) > 0 {
		t.Errorf("Expected no errors, got %d errors: %v", len(errors), errors)
	}

	if !(len(btcConfig.Sections) == 2) {
		t.Errorf("Expected 2 sections, got %d", len(btcConfig.Sections))
	}

	if !(btcConfig.Sections[0].Header == nil) {
		t.Errorf("Expected first section to be a root section, but it has a header: %v", btcConfig.Sections[0].Header)
	}

	if !(btcConfig.Sections[0].Properties.Size() == 7) {
		t.Errorf("Expected first section to have 7 properties, but it has %d", btcConfig.Sections[0].Properties.Size())
	}

	if !(btcConfig.Sections[1].Header.Name == "regtest") {
		t.Errorf("Expected second section header to be 'regtest', but got '%s'", btcConfig.Sections[1].Header.Name)
	}

	if !(btcConfig.Sections[1].Properties.Size() == 1) {
		t.Errorf("Expected second section to have 1 property, but it has %d", btcConfig.Sections[1].Properties.Size())
	}
}

// From https://github.com/AdamISZ/pathcoin-poc/blob/78142488d60298f8541973891acd7f1d0013fdca/bitcoin.conf.sample
func BTCConfigComplexExampleTest(t *testing.T) {
	content := `
[regtest]
rpcuser=bitcoinrpc
rpcpassword=123456abcdef
fallbackfee=0.00005
[signet]
addnode=inquisition.bitcoin-signet.net
rpcport = 38332
rpcuser=bitcoinrpc
rpcpassword=123456abcdef
fallbackfee=0.00005
`

	btcConfig := NewBTCConfig()
	errors := btcConfig.Parse(content)

	if len(errors) > 0 {
		t.Errorf("Expected no errors, got %d errors: %v", len(errors), errors)
	}
}
