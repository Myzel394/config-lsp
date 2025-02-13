interface LSPNotification {
	Uri: string;
}

interface LSPLanguageUndetectable extends LSPNotification {}

interface LSPLanguageDetected extends LSPNotification {
	Language: string;
}

