package analyzer

import (
	testutils_test "config-lsp/handlers/fstab/test_utils"
	"testing"
)

func TestSambaExample(t *testing.T) {
	// Example from https://askubuntu.com/a/313389/1090198
	document := testutils_test.DocumentFromInput(t, `
//192.168.0.5/storage /media/myname/TK-Public/ cifs guest,uid=myuser,iocharset=utf8,file_mode=0777,dir_mode=0777,noperm 0 0
`)

	diagnostics := Analyze(document)

	if len(diagnostics) != 0 {
		t.Fatalf("Expected 0 diagnostics, got %d", len(diagnostics))
	}
}

func TestSambaExampleWithSubdirectories(t *testing.T) {
	// Example from https://github.com/Myzel394/config-lsp/issues/51#issuecomment-3076535354
	document := testutils_test.DocumentFromInput(t, `
//192.168.0.5/sharename/sub/directory /mnt/mountpoint cifs nofail,username=myuser,password=mypass 0 0
`)

	diagnostics := Analyze(document)

	if len(diagnostics) != 0 {
		t.Fatalf("Expected 0 diagnostics, got %d", len(diagnostics))
	}
}
