package projectspecificsetup

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestGoProjectCreatesBasicStructure(t *testing.T) {
	root := t.TempDir()
	GoProject("demo", root)

	paths := []string{
		"go.mod",
		filepath.Join("cmd", "demo", "main.go"),
		filepath.Join("src", "demo", "main.go"),
		filepath.Join("internal", "cli", "cli.go"),
		filepath.Join("scripts", "CICD.sh"),
	}
	for _, relativePath := range paths {
		if _, err := os.Stat(filepath.Join(root, relativePath)); err != nil {
			t.Fatalf("expected Go project file %s: %v", relativePath, err)
		}
	}

	content, err := os.ReadFile(filepath.Join(root, "src", "demo", "main.go"))
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(string(content), "package main") {
		t.Fatal("generated main.go should declare package main")
	}
}

func TestPythonProjectCreatesBasicStructure(t *testing.T) {
	root := t.TempDir()
	PythonProject("demo", root)

	paths := []string{
		"setup.py",
		filepath.Join("src", "demo", "__init__.py"),
		filepath.Join("src", "demo", "__main__.py"),
		filepath.Join("tests", "test1.py"),
		filepath.Join("scripts", "CICD.sh"),
	}
	for _, relativePath := range paths {
		if _, err := os.Stat(filepath.Join(root, relativePath)); err != nil {
			t.Fatalf("expected Python project file %s: %v", relativePath, err)
		}
	}

	content, err := os.ReadFile(filepath.Join(root, "setup.py"))
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(string(content), `name="demo"`) {
		t.Fatal("generated setup.py should contain the project name")
	}
}
