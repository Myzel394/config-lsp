# /sshd_config/lsp

This folder is the glue between our language server and the LSP
clients.
This folder only contains LSP commands.
It only handles very little actual logic, and instead calls
the handlers from `../handlers`.