#!/usr/bin/env just --justfile

set dotenv-load := true

alias build-vsc-windows := build-vs-code-extension-for-windows

default:
  @just --list

# Lint whole project
[working-directory: "./server"]
lint:
    gofmt -s -w .
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

[working-directory: "./server"]
update-antlr-parsers:
    # aliases
    cd handlers/aliases && antlr4 -Dlanguage=Go -o ast/parser Aliases.g4

    # fstab
    cd handlers/fstab && antlr4 -Dlanguage=Go -o ast/parser Fstab.g4

    # sshd_config
    cd handlers/sshd_config && antlr4 -Dlanguage=Go -o ast/parser Config.g4
    cd handlers/sshd_config/match-parser && antlr4 -Dlanguage=Go -o parser Match.g4

    # ssh_config
    cd handlers/ssh_config && antlr4 -Dlanguage=Go -o ast/parser Config.g4
    cd handlers/ssh_config/match-parser && antlr4 -Dlanguage=Go -o parser Match.g4

    # hosts
    cd handlers/hosts && antlr4 -Dlanguage=Go -o ast/parser Hosts.g4

[working-directory: "./vs-code-extension"]
update-yarn:
    yarn install --no-frozen-lockfile && yarn2nix > yarn.nix

build-vs-code-extension-for-windows:
    nix build .#vs-code-extension-bare
    cd server && GOOS="windows" GOARCH="amd64" go build -a -gcflags=all="-l -B" -ldflags="-s -w" -o config-lsp.exe
    rm -rf dist
    mkdir dist
    cp -r result/* dist/
    chmod -R 755 dist
    mv server/config-lsp.exe ./dist/out/
    cd dist && vsce package --target "win32-x64"
    cp dist/config-lsp-*.vsix .
    rm -rf result
    rm -rf dist

# Ready for a PR? Run this recipe before opening the PR!
ready:
    just lint
    just test

[working-directory: "./website"]
dev-website:
    yarn run dev

