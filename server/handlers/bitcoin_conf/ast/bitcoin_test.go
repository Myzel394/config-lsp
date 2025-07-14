package ast

import "testing"

func SimpleBTCConfigExample(t *testing.T) {
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
