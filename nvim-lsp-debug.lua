vim.lsp.start {
    name = "config-lsp",
    cmd = { "./bin/config-lsp" },
    root_dir = vim.fn.getcwd(),
};
