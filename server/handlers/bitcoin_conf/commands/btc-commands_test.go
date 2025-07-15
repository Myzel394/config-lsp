package commands

import "testing"

func TestGenerateRPCAuth(t *testing.T) {
	if !IsPythonAvailable() {
		t.Skip("Python is not available")
	}

	username := "testuser"
	password := "testpassword"

	auth, err := GenerateRPCAuth(username, password)

	if err != nil {
		t.Fatal(err)
	}

	if auth == "" {
		t.Fatal("Generated auth is empty")
	}

	t.Logf("Generated RPC Auth: %s", auth)
}
