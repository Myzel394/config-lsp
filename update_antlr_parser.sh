#!/bin/sh

GIT_ROOT=$(git rev-parse --show-toplevel)

# aliases
cd $GIT_ROOT/handlers/aliases && antlr4 -Dlanguage=Go -o ast/parser Aliases.g4

# sshd_config
cd $GIT_ROOT/handlers/sshd_config && antlr4 -Dlanguage=Go -o ast/parser Config.g4
cd $GIT_ROOT/handlers/sshd_config/match-parser && antlr4 -Dlanguage=Go -o parser Match.g4

# ssh_config
cd $GIT_ROOT/handlers/ssh_config && antlr4 -Dlanguage=Go -o ast/parser Config.g4
cd $GIT_ROOT/handlers/ssh_config/match-parser && antlr4 -Dlanguage=Go -o parser Match.g4

# hosts
cd $GIT_ROOT/handlers/hosts && antlr4 -Dlanguage=Go -o ast/parser Hosts.g4

