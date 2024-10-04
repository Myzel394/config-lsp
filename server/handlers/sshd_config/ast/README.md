# /sshd_config/ast

This folder contains the AST (Abstract Syntax Tree) for the handlers.
The AST is defined in a filename that's the same as the handler's name.

Each AST node must extend the following fields:

```go
type ASTNode struct {
    common.LocationRange
    Value commonparser.ParsedString
}
```

Each node should use a shared prefix for the node name,
e.g. `SSHDConfig`, `SSDKey` for the `sshd_config` handler.