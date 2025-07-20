# config-lsp

A language server for configuration files. The goal is to make editing config files modern and easy.

## Supported Features

|             | diagnostics | `completion` | `hover` | `code-action` | `definition` | `rename` | `signature-help` |
|-------------|-------------|--------------|---------|---------------|--------------|----------|------------------|
| aliases     | ✅           | ✅            | ✅       | ✅             | ✅            | ✅        | ✅               |
| fstab       | ✅           | ✅            | ✅       | ❓             | ❓            | ❓        | 🟡               |
| hosts       | ✅           | ✅            | ✅       | ✅             | ❓            | ❓        | ✅               |
| ssh_config  | ✅           | ✅            | ✅       | ✅             | ✅            | ✅        | ✅               |
| sshd_config | ✅           | ✅            | ✅       | ❓             | ✅            | ❓        | ✅               |
| wireguard   | ✅           | ✅            | ✅       | ✅             | ❓            | ❓        | 🟡               |
| bitcoin_conf| ✅           | ✅            | ✅       | ✅             | ❓            | ❓        | 🟡               |

✅ = Supported

🟡 = Will be supported, but not yet implemented

❓ = No idea what to implement here, please let me know if you have any ideas

### Supported Code Actions

#### Bitcoin
- `Generate rpcauth` - config-lsp reads the `rpcuser` and `rpcpassword` from the config file and generates a new `rpcauth` line for you, which includes a properly salted and hashed password (this requires Python to be installed)
- `Fix typo` - fix a typo

#### Aliases
- `Send test mail` - send a test mail to the current email address in the `alias` file (this requires `sendmail` / `postfix` to be installed)

#### Hosts
- `Inline aliases` - inline duplicated aliases into one line

#### SSH Config
- `Generate SSH Key` - generate a new SSH key and add it to the config file (this requires `ssh-keygen` to be installed)
- `Add option to unknown` - add an option to `IgnoreUnknown`. This is mostly used for `UseKeychain`
- `Fix typo` - fix a typo

#### Wireguard
- `Generate new peer based on this one` - generate a new peer config section that's similar the the current one. Useful for quickly adding new peers to a Wireguard config file
- `Generate PostDown` - generate an adjacent `PostDown` command based on the current `PreUp` command. Only supports iptables currently
- `Add PersistentKeepalive` - add a `PersistentKeepalive` option to the current peer config section
- `Generate PrivateKey` / `Generate PresharedKey` - generate a new private key or preshared key and insert it automatically
- `Fix typo` - fix a typo

## Installation

### VS Code Extension

[Install the extension from the marketplace](https://marketplace.visualstudio.com/items?itemName=myzel394.config-lsp)

Alternatively, you can also manually install the extension:

1. Download the latest extension version from the [release page](https://github.com/Myzel394/config-lsp/releases) - You can find the extension under the "assets" section. The filename ends with `.vsix`
2. Open VS Code
3. Open the extensions sidebar
4. In the top bar, click on the three dots and select "Install from VSIX..."
5. Select the just downloaded `.vsix` file
6. You may need to restart VS Code
7. Enjoy!

### Manual installation

To use `config-lsp` in any other editor, you'll need to install it manually.
Don't worry, it's easy!

#### Installing the latest Binary

##### Brew

```sh
brew install myzel394/formulae/config-lsp
```

##### Manual Binary

Download the latest binary from the [releases page](https://github.com/Myzel394/config-lsp/releases) and put it in your PATH.

##### Compiling

You can either compile the binary using go:

```sh
go build -o config-lsp
```

or build it using Nix:

```sh
nix flake build
```

#### Neovim installation

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


## What further configs will be supported?

As config-lsp is a hobby project and I'm working completely alone on it, 
I will first focus on widely used and well known config files.

You are welcome to request any config file, as far as it's fairly well known.

## Configuration

config-lsp supports the following CLI options:

| Option                        | Description                                                               | Default value                                         |
|-------------------------------|---------------------------------------------------------------------------|-------------------------------------------------------|
| `--no-undetectable-errors`    | If config-lsp is unable to detect the language, do not return any errors. | Disabled by default; Enabled in the VS Code extension |
| `--no-typo-suggestions`       | Do not suggest typo fixes                                                 | Disabled by default                                   |
| `--usage-reports-errors-only` | Only report errors.                                                       | By default all usage telemetry is enabled             |
| `--usage-reports-disable` | Disable all telemetry                                                       | By default all usage telemetry is enabled             |

### Telemetry

Starting with version `0.3.0`, config-lsp collects usage telemetry to improve the project by default. The VS Code extension also collects telemetry by default, unless the built-in telemetry setting is changed to either errors-only or disabled.

#### What telemetry is collected?

Data is sent to my self-hosted Sentry instance. 30 days retention, no personal data, not shared with third parties.

We do not collect any personal data. We also do not collect the contents of your config files. We may collect the following data (but not limited to):
* Version of config-lsp
* The operating system
* CLI arguments used to start config-lsp
* Detected language of the config file
* The path to the file
* Logs generated by config-lsp
* Errorrs produced by config-lsp
* Used editor (if available)

We would be really happy if you would _not_ disable telemetry, as this helps me improve the project, add new features, and fix bugs faster. No data is shared with third parties, and I do not use the data to track you personally. The data is only used to improve config-lsp.


## Supporting config-lsp

You can either contribute to the project, [see CONTRIBUTING.md](CONTRIBUTING.md), or you can sponsor me via [GitHub Sponsors](https://github.com/sponsors/Myzel394) or via [crypto currencies](https://github.com/Myzel394/contact-me?tab=readme-ov-file#donations).

Oh and spreading the word about config-lsp is also a great way to support the project :)


