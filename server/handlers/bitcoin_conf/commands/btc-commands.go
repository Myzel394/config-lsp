package commands

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"os"
	"os/exec"
)

const BITCOIN_RPC_AUTH_SCRIPT_URL = "https://raw.githubusercontent.com/bitcoin/bitcoin/184159e4f30c6846cd0324ee07b59e7f8121a5ea/share/rpcauth/rpcauth.py"

// Not sure yet how to properly delete the file while caching it.
// At the moment, we simply delegate this responsibility to the operating system.
//
//	defer func() {
//		_ = file.Close()
//		_ = os.Remove(file.Name())
//	}()
var scriptFile *os.File

// Download the Bitcoin RPC auth script
// This caches the script in a temporary file
func downloadScript() (*os.File, error) {
	if scriptFile != nil {
		return scriptFile, nil
	}

	file, err := os.CreateTemp("", "bitcoin-rpc-auth-*.py")

	if err != nil {
		return nil, err
	}

	url := BITCOIN_RPC_AUTH_SCRIPT_URL

	resp, err := http.Get(url)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("Failed to download script: " + resp.Status)
	}

	scriptContent, err := io.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	_, err = file.Write(scriptContent)

	if err != nil {
		return nil, err
	}

	scriptFile = file

	return scriptFile, nil
}

func executeScript(file *os.File, username, password string) (string, error) {
	cmd := exec.Command("python3", file.Name(), username, password, "--json")

	output, err := cmd.CombinedOutput()

	if err != nil {
		return "", err
	}

	return string(output), nil
}

type scriptResponse struct {
	RPCauth string `json:"rpcauth"`
}

func GenerateRPCAuth(
	username string,
	password string,
) (string, error) {
	if !IsPythonAvailable() {
		return "", errors.New("Python 3 must be installed to generate the RPC auth")
	}

	if username == "" || password == "" {
		return "", errors.New("Username and password must be provided")
	}

	file, err := downloadScript()

	if err != nil {
		return "", err
	}

	output, err := executeScript(file, username, password)

	if err != nil {
		return "", err
	}

	response := scriptResponse{}
	err = json.Unmarshal([]byte(output), &response)

	if err != nil {
		return "", errors.New("Failed to parse script output: " + err.Error())
	}

	return response.RPCauth, nil
}
