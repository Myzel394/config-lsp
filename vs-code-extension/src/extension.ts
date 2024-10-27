import * as path from "path";
import { ExtensionContext, workspace } from "vscode";

import {
	Executable,
	LanguageClient,
	type LanguageClientOptions,
	type ServerOptions,
} from "vscode-languageclient/node";
import { onUndetectable } from "./events/on-undetectable";

const IS_DEBUG =
	process.env.VSCODE_DEBUG_MODE === "true" ||
	process.env.NODE_ENV === "development";
let client: LanguageClient;

export async function activate({subscriptions}: ExtensionContext) {
	console.info("config-lsp activated");
	const initOptions = workspace.getConfiguration("config-lsp");
	const clientOptions: LanguageClientOptions = {
		documentSelector: [
			{language: "sshconfig"},
			{language: "sshdconfig"},
			{language: "aliases"},
			{language: "fstab"},
			{language: "hosts"},
			{language: "wireguard"},
		],
		initializationOptions: initOptions,
	};

	const path = getBundledPath();
	console.info(`Found config-lsp path at ${path}`);
	const run: Executable = {
		command: getBundledPath() ,
	};
	const serverOptions: ServerOptions = {
		run,
		debug: run,
	};

	client = new LanguageClient(
		"config-lsp",
		serverOptions,
		clientOptions,
		IS_DEBUG
	);

	console.info("Starting config-lsp...");
	await client.start();
	console.info("Started config-lsp");

	subscriptions.push(client.onNotification("$/config-lsp/languageUndetectable", onUndetectable))
}

function getBundledPath(): string {
	const filePath = path.resolve(__dirname, "config-lsp");

	return filePath;
}

export function deactivate(): Thenable<void> | undefined {
	if (!client) {
		return undefined;
	}

	return client.stop();
}
