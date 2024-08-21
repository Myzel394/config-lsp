package wgcommands

import "testing"

func TestWireguardAvailable(
	t *testing.T,
) {
	if !AreWireguardToolsAvailable() {
		t.Skip("Wireguard tools not available")
	}
}

func TestWireguardPrivateKey(
	t *testing.T,
) {
	privateKey, err := CreateNewPrivateKey()

	if err != nil {
		t.Fatal(err)
	}

	t.Log(privateKey)
}

func TestWireguardPublicKey(
	t *testing.T,
) {
	privateKey := "UPBKR0kLF2C/+Ei5fwN5KHsAcon9xfBX+RWhebYFGWg="
	publicKey, err := CreatePublicKey(privateKey)

	if err != nil {
		t.Fatal(err)
	}

	if publicKey != "3IPUqUKXUkkU7tNp/G/KgcBqUh3N0WWJpfQf79lGdl0=" {
		t.Fatalf("Public key does not match, it's: %v", publicKey)
	}

	t.Log(publicKey)
}
