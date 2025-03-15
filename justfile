#!/usr/bin/env just --justfile

set dotenv-load := true

default:
  @just --list

# Lint whole project
lint:
    cd server && gofmt -s -w .
    # cd vs-code-extension && yarn run lint

# Build config-lsp and test it in nvim (config-lsp will be loaded automatically)
[working-directory: "./server"]
test-nvim file:
    go build -o ./result/bin/config-lsp && rm -rf ~/.local/state/nvim/lsp.log && DOTFILES_IGNORE_CONFIG_LSP=1 nvim {{file}} -c ':source nvim-lsp-debug.lua'

# Show Mason Logs
show-nvim-logs:
    bat ~/.local/state/nvim/lsp.log

[working-directory: "./server"]
test:
    nix develop --command bash -c 'go test ./... -count=1'

# Ready for a PR? Run this recipe before opening the PR!
ready:
    just lint
    just test

