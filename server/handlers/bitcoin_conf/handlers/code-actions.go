package handlers

import (
	"config-lsp/handlers/wireguard/ast"

	protocol "github.com/tliron/glsp/protocol_3_16"
)

type CodeActionName string

const (
	CodeActionGenerateRPCAuth CodeActionName = "generateRPCAuth"
)

type CodeAction interface {
	RunCommand(*ast.WGConfig) (*protocol.ApplyWorkspaceEditParams, error)
}

type CodeActionArgs interface{}
