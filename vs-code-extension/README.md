# config-lsp for VS Code

`config-lsp` provides language support for various config files.
Currently it supports completions, diagnostics, hints, formatting, hover information,
and definition requests.

Install this extension and load your config files in VS Code to get started.

If `config-lsp` is unable to detect the language of your config file, you can manually
specify it by adding a line in the form of:

```plaintext
#?lsp.language=<language>

# For example

#?lsp.language=sshconfig
#?lsp.language=fstab
#?lsp.language=aliases
```

