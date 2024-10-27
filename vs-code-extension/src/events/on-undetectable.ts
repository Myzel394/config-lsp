import { GenericNotificationHandler } from "vscode-languageclient";
import * as vscode from "vscode";

const ACTION_SELECT_LANGUAGE = "Select Language";
const ACTION_DISABLE = "Ignore for this file";

const ignoredFiled = new Set<string>();

export const onUndetectable: GenericNotificationHandler = async (params: LSPLanguageUndetectable) => {
	if (ignoredFiled.has(params.Uri)) {
		return;
	}

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
			ignoredFiled.add(params.Uri);
			break;
		undefined:
			break;
	}
}

