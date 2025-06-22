package ast

import (
	"config-lsp/common"
	"config-lsp/parsers/ini"
)


type WGConfig struct {
	*ini.Config
}

// NewWGConfig creates a new WireGuard configuration
func NewWGConfig() *WGConfig {
	return &WGConfig{
		Config: ini.NewConfig(),
	}
}

// Reset the configuration
func (c *WGConfig) Clear() {
	c.Config.Clear()
}

// Parse a WireGuard configuration string
func (c *WGConfig) Parse(input string) []common.LSPError {
	return c.Config.Parse(input)
}

