vim.lsp.start {
    name = "config-lsp",
    cmd = { "./result/bin/config-lsp" },
    root_dir = vim.fn.getcwd(),
};
