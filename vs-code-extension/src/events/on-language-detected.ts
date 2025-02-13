import { GenericNotificationHandler } from "vscode-languageclient";
import { tempUndetectableFiles } from "./shared";

export const onLanguageDetectable: GenericNotificationHandler = async (params: LSPLanguageDetected) => {
	// Hide warning box
	tempUndetectableFiles.add(params.Uri);
}

