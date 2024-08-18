package wireguard

import "testing"

func TestWireguardAvailable(
	t *testing.T,
) {
	if !areWireguardToolsAvailable() {
		t.Skip("Wireguard tools not available")
	}
}

func TestWireguardPrivateKey(
	t *testing.T,
) {
	privateKey, err := createNewPrivateKey()

	if err != nil {
		t.Fatal(err)
	}

	t.Log(privateKey)
}

func TestWireguardPublicKey(
	t *testing.T,
) {
	privateKey := "UPBKR0kLF2C/+Ei5fwN5KHsAcon9xfBX+RWhebYFGWg="
	publicKey, err := createPublicKey(privateKey)

	if err != nil {
		t.Fatal(err)
	}

	if publicKey != "3IPUqUKXUkkU7tNp/G/KgcBqUh3N0WWJpfQf79lGdl0=" {
		t.Fatalf("Public key does not match, it's: %v", publicKey)
	}

	t.Log(publicKey)
}
