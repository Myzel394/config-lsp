package ast

import (
	"config-lsp/common"
	"config-lsp/parsers/ini"
)

type BTCConfig struct {
	*ini.Config
}

// NewBTCConfig creates a new Bitcoin configuration
func NewBTCConfig() *BTCConfig {
	config := ini.NewConfig()
	config.XParseConfig = ini.INIParseConfig{
		AllowRootProperties: true,
	}

	return &BTCConfig{
		Config: ini.NewConfig(),
	}
}

// Reset the configuration
func (c *BTCConfig) Clear() {
	c.Config.Clear()
}

// Parse a Bitcoin configuration string
func (c *BTCConfig) Parse(input string) []common.LSPError {
	return c.Config.Parse(input)
}
