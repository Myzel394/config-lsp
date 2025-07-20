vim.lsp.start {
    name = "config-lsp",
    cmd = { "./result/bin/config-lsp", "--env-debug" },
    root_dir = vim.fn.getcwd(),
};
