import { GenericNotificationHandler } from "vscode-languageclient";
import * as vscode from "vscode";
import { tempUndetectableFiles, undetectableIgnoredFiles } from "./shared";

const ACTION_SELECT_LANGUAGE = "Select Language";
const ACTION_DISABLE = "Ignore for this file";

export const onLanguageUndetectable: GenericNotificationHandler = async (params: LSPLanguageUndetectable) => {
	if (undetectableIgnoredFiles.has(params.Uri) || tempUndetectableFiles.has(params.Uri)) {
		return;
	}

	tempUndetectableFiles.add(params.Uri);

	const result = await vscode.window.showWarningMessage(
		"config-lsp was unable to detect the appropriate language for this file",
		{
			detail: "Either select a language or add '#?lsp.language=<language>' to the top of the file",
		},
		ACTION_SELECT_LANGUAGE,
		ACTION_DISABLE,
	)

	switch (result) {
		case ACTION_SELECT_LANGUAGE:
			vscode.commands.executeCommand("workbench.action.editor.changeLanguageMode");
			break;
		case ACTION_DISABLE:
			undetectableIgnoredFiles.add(params.Uri);
			break;
	}
}

