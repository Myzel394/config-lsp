import * as path from "path";
import { ExtensionContext } from 'vscode';

import {
    Executable,
	LanguageClient,
	type LanguageClientOptions,
	type ServerOptions,
} from 'vscode-languageclient/node';

const IS_DEBUG = process.env.VSCODE_DEBUG_MODE === 'true' || process.env.NODE_ENV === 'development';
let client: LanguageClient;

export function activate(context: ExtensionContext) {
		console.info("config-lsp activated");
	const clientOptions: LanguageClientOptions = {
		documentSelector: [
			{
				scheme: 'file',
				language: 'plaintext',
				pattern: "**/{config,sshconfig,sshd_config,sshdconfig,fstab,hosts,aliases}",
			},
			// Some configs seem to be incorrectly detected as yaml
			{
				scheme: 'file',
				language: 'yaml',
				pattern: "**/{config,sshconfig,sshd_config,sshdconfig,fstab,hosts,aliases}",
			},
		]
	};
	
	const path = getBundledPath();
		console.info(`Found config-lsp path at ${path}`);
	const run: Executable = {
		command: getBundledPath(),
	}
	const serverOptions: ServerOptions = {
		run,
		debug: run,
	}

	client = new LanguageClient(
		'config-lsp',
		serverOptions,
		clientOptions,
		IS_DEBUG,
	);


	client.start();
	console.info("config-lsp started");

	// const serverOptions: ServerOptions = {
	// }
	//
	// // Create the language client and start the client.
	// client = new LanguageClient(
	// 	'languageServerExample',
	// 	clientOptions
	// );
	//
	// // Start the client. This will also launch the server
	// client.start();
}

function getBundledPath(): string {
	const filePath = path.resolve(__dirname, "config-lsp")

	return filePath;
}

export function deactivate(): Thenable<void> | undefined {
	if (!client) {
		return undefined;
	}

	return client.stop();
}
