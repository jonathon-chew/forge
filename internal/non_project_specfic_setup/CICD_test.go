package nonprojectspecficsetup

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestCICDScriptCreatesScript(t *testing.T) {
	root := t.TempDir()

	CICD_Script(root)

	path := filepath.Join(root, "scripts", "CICD.sh")
	content, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("expected CI script to be created: %v", err)
	}
	if !strings.Contains(string(content), "go test -v ./...") {
		t.Fatal("expected CI script to run Go tests")
	}
}
