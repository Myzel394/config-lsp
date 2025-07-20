import * as path from "path";
import { ExtensionContext, workspace, env } from "vscode";

import {
	Executable,
	LanguageClient,
	type LanguageClientOptions,
	type ServerOptions,
} from "vscode-languageclient/node";
import { onLanguageUndetectable } from "./events/on-language-undetectable";
import { onLanguageDetectable } from "./events/on-language-detected";

const IS_DEBUG =
	process.env.VSCODE_DEBUG_MODE === "true" ||
	process.env.NODE_ENV === "development";
let client: LanguageClient;

export async function activate({subscriptions}: ExtensionContext) {
	console.info("config-lsp activated");

	await startClient();

	subscriptions.push(client.onNotification("$/config-lsp/languageUndetectable", onLanguageUndetectable))
	subscriptions.push(client.onNotification("$/config-lsp/detectedLanguage", onLanguageDetectable))
	subscriptions.push(env.onDidChangeTelemetryEnabled(async (enabled) => {
		console.info(`Telemetry enabled: ${enabled}, restarting client...`);

		try {
			client.stop()
		} catch (error) {
			console.error("Error stopping client:", error);
		}

		await startClient();

		console.info("Client restarted after telemetry change");
	}))
}

async function startClient() {
	const initOptions = workspace.getConfiguration("config-lsp");
	const clientOptions: LanguageClientOptions = {
		documentSelector: [
			{language: "sshconfig"},
			{language: "sshdconfig"},
			{language: "aliases"},
			{language: "fstab"},
			{language: "hosts"},
			{language: "wireguard"},
			{language: "bitcoin_conf"}
		],
		initializationOptions: initOptions,
	};

	const path = getBundledPath();
	console.info(`Found config-lsp path at ${path}`);

	const args = [
		"--no-undetectable-errors",
	];
	if (env.isTelemetryEnabled) {
		const telemetryLevel = workspace.getConfiguration("telemetry").get<string>("telemetryLevel", "off");

		switch (telemetryLevel) {
			case "off":
				console.info("Telemetry is disabled, passing --usage-reports-disable to config-lsp");
				args.push("--usage-reports-disable");
				break;
			case "error":
				args.push("--usage-reports-errors-only");
				break;
		}
	} else {
		console.info("Telemetry is disabled, passing --usage-reports-disable to config-lsp");
		args.push("--usage-reports-disable")
	}

	const run: Executable = {
		command: path,
		args,
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
}

function getBundledPath(): string {
	if (process.platform === "win32") {
		return path.resolve(__dirname, "config-lsp.exe");
	}

	return path.resolve(__dirname, "config-lsp");
}

export function deactivate(): Thenable<void> | undefined {
	if (!client) {
		return undefined;
	}

	return client.stop();
}
