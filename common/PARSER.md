# Parser

The parser used in `config-lsp` work on the following principles:

1. Read the configuration file and divide each line into sections
2. On changes, run the "analyzer", which will then actually parse the sections and check for errors
3. On completion requests, check which section is being requested and return completions for it
4. On hover requests, check which section is being requested and return the hover information for it

