// config-lsp constantly sends a undetectable message for each LSP request.
// This gets annoying to users quickly as they will be spammed with error messages.
// To avoid this, we temporarily ignore files that have sent an undetectable message.
// 
// This will be cleared once the language has been detected.
// This is different from `undetectableIgnoredFiles`.
// When a user selects "Ignore for this file", we will add the file to `undetectableIgnoredFiles`.
// Then, we **never** show a warning for that file again.
export const tempUndetectableFiles = new Set<string>();

export const undetectableIgnoredFiles = new Set<string>();


