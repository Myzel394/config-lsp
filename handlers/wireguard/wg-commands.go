package wireguard

import (
	"os/exec"
	"strings"
)

func areWireguardToolsAvailable() bool {
	_, err := exec.LookPath("wg")

	return err == nil
}

func createNewPrivateKey() (string, error) {
	cmd := exec.Command("wg", "genkey")

	bytes, err := cmd.Output()

	if err != nil {
		return "", err
	}

	return string(bytes), nil
}

func createPublicKey(privateKey string) (string, error) {
	cmd := exec.Command("wg", "pubkey")
	cmd.Stdin = strings.NewReader(privateKey)

	bytes, err := cmd.Output()

	if err != nil {
		return "", err
	}

	return strings.ReplaceAll(string(bytes), "\n", ""), nil
}
