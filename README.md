# config-lsp

A language server for configuration files. The goal is to make editing config files modern and easy.

## Supported Features

|             | diagnostics | `completion` | `hover` | `code-action` | `definition` | `rename` | `signature-help` |
|-------------|-------------|--------------|---------|---------------|--------------|----------|------------------|
| aliases     | ✅           | ✅            | ✅       | ✅             | ✅            | ✅        | ✅               |
| fstab       | ✅           | ✅            | ✅       | ❓             | ❓            | ❓        | 🟡               |
| hosts       | ✅           | ✅            | ✅       | ✅             | ❓            | ❓        | 🟡               |
| ssh_config  | ✅           | ✅            | ✅       | ✅             | ✅            | ✅        | ✅               |
| sshd_config | ✅           | ✅            | ✅       | ❓             | ✅            | ❓        | ✅               |
| wireguard   | ✅           | ✅            | ✅       | ✅             | ❓            | ❓        | 🟡               |

✅ = Supported

🟡 = Will be supported, but not yet implemented

❓ = No idea what to implement here, please let me know if you have any ideas

## What further configs will be supported?

As config-lsp is a hobby project and I'm working completely alone on it, 
I will first focus on widely used and well known config files.

You are welcome to request any config file, as far as it's fairly well known.

## Installation

Download the latest binary from the [releases page](https://github.com/Myzel394/config-lsp/releases) and put it in your PATH.

Follow the instructions for your editor below.

### Neovim installation

Using [nvim-lspconfig](https://github.com/neovim/nvim-lspconfig) you can add `config-lsp` by adding the following to your `lsp.lua` (filename might differ):

```lua
if not configs.config_lsp then
    configs.config_lsp = {
        default_config = {
            cmd = { 'config-lsp' },
            filetypes = {
                "sshconfig",
                "sshdconfig",
                "fstab",
                "aliases",
                -- Matches wireguard configs and /etc/hosts
                "conf",
            },
            root_dir = vim.loop.cwd,
        },
    }
end

lspconfig.config_lsp.setup {}
`````

### VS Code installation

The VS Code extension is currently in development. An official extension will be released soon.

However, at the moment you can also compile the extension yourself and run it in development mode.

**Do not create an extension and publish it yourself. Contribute to the official extension instead.**

## Supporting config-lsp

You can either contribute to the project, [see CONTRIBUTING.md](CONTRIBUTING.md), or you can sponsor me via [GitHub Sponsors](https://github.com/sponsors/Myzel394) or via [crypto currencies](https://github.com/Myzel394/contact-me?tab=readme-ov-file#donations).

Oh and spreading the word about config-lsp is also a great way to support the project :)


