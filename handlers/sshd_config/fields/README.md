# /sshd_config/fields

This folder contains the glue between the config documentation and
our language server.

`fields.go` usually contains a list of all the available options / fields
the config file offers.

It's usually a map of:

```
OptionName: docValue
```

A docvalue is a simple parser that can validate the value and fetch completions for it.
You should use a docvalue for most types. See `doc-values` for all available docvalues.

However, some types are so complex or require more context to validate or fetch completions,
that we need to write a custom function for it. In that case we do it dirty and simply check
in the completions if the key is the one we are looking for, and then execute the custom function.
```
