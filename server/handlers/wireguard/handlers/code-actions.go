package handlers

import (
	"config-lsp/handlers/wireguard/ast"

	protocol "github.com/tliron/glsp/protocol_3_16"
)

type CodeActionName string

const (
	CodeActionGeneratePrivateKey   CodeActionName = "generatePrivateKey"
	CodeActionGeneratePresharedKey CodeActionName = "generatePresharedKey"
	CodeActionCreatePeer           CodeActionName = "createPeer"
	CodeActionGenerateDownRule     CodeActionName = "generatePostDown"
)

type CodeAction interface {
	RunCommand(*ast.WGConfig) (*protocol.ApplyWorkspaceEditParams, error)
}

type CodeActionArgs interface{}
