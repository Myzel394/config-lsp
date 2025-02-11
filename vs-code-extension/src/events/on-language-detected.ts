import { GenericNotificationHandler } from "vscode-languageclient";
import { tempUndetectableFiles } from "./shared";

export const onDetectable: GenericNotificationHandler = async (params: LSPLanguageDetected) => {
	tempUndetectableFiles.add(params.Uri);
}

