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

## Supporting config-lsp

You can either contribute to the project, [see CONTRIBUTING.md](CONTRIBUTING.md), or you can sponsor me via [GitHub Sponsors](https://github.com/sponsors/Myzel394) or via [crypto currencies](https://github.com/Myzel394/contact-me?tab=readme-ov-file#donations).

Oh and spreading the word about config-lsp is also a great way to support the project :)


