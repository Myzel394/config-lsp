#!/bin/sh

ROOT=$(git rev-parse --show-toplevel)/server

# aliases
cd $ROOT/handlers/aliases && antlr4 -Dlanguage=Go -o ast/parser Aliases.g4

# fstab
cd $ROOT/handlers/fstab && antlr4 -Dlanguage=Go -o ast/parser Fstab.g4

# gitconfig
cd $ROOT/handlers/gitconfig && antlr4 -Dlanguage=Go -o ast/parser Config.g4

# sshd_config
cd $ROOT/handlers/sshd_config && antlr4 -Dlanguage=Go -o ast/parser Config.g4
cd $ROOT/handlers/sshd_config/match-parser && antlr4 -Dlanguage=Go -o parser Match.g4

# ssh_config
cd $ROOT/handlers/ssh_config && antlr4 -Dlanguage=Go -o ast/parser Config.g4
cd $ROOT/handlers/ssh_config/match-parser && antlr4 -Dlanguage=Go -o parser Match.g4

# hosts
cd $ROOT/handlers/hosts && antlr4 -Dlanguage=Go -o ast/parser Hosts.g4

