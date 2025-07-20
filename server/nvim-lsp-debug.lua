vim.lsp.start {
    name = "config-lsp",
    cmd = { "./result/bin/config-lsp", "--env-debug", "--disable-usage-reports" },
    root_dir = vim.fn.getcwd(),
};
