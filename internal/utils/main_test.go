package utils

import (
	"os"
	"path/filepath"
	"testing"
)

func TestWriteFileCreatesParentDirectories(t *testing.T) {
	root := t.TempDir()
	path := filepath.Join(root, "nested", "example.txt")

	WriteFile(path, []byte("hello"))

	content, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("expected file to be created: %v", err)
	}
	if string(content) != "hello" {
		t.Fatalf("expected file content %q, got %q", "hello", string(content))
	}
}

func TestMakeFoldersCreatesNestedFolders(t *testing.T) {
	root := t.TempDir()

	MakeFolders(root, []string{filepath.Join("one", "two")})

	info, err := os.Stat(filepath.Join(root, "one", "two"))
	if err != nil {
		t.Fatalf("expected nested folder to be created: %v", err)
	}
	if !info.IsDir() {
		t.Fatal("expected nested path to be a directory")
	}
}

func TestMakeFilesCreatesFilesAndParents(t *testing.T) {
	root := t.TempDir()

	MakeFiles(root, []string{filepath.Join("nested", "one.txt"), "two.txt"})

	for _, relativePath := range []string{filepath.Join("nested", "one.txt"), "two.txt"} {
		path := filepath.Join(root, relativePath)
		info, err := os.Stat(path)
		if err != nil {
			t.Fatalf("expected %s to be created: %v", relativePath, err)
		}
		if info.IsDir() {
			t.Fatalf("expected %s to be a file", relativePath)
		}
	}
}

func TestRunCommandsUsesWorkingDirectory(t *testing.T) {
	root := t.TempDir()

	stdout, stderr, err := RunCommands(Command{
		Command: []string{"sh", "-c", "pwd"},
		Name:    "print working directory",
	}, root)
	if err != nil {
		t.Fatalf("expected command to succeed: %v", err)
	}
	if stderr != "" {
		t.Fatalf("expected no stderr, got %q", stderr)
	}

	got, err := filepath.Abs(stdout[:len(stdout)-1])
	if err != nil {
		t.Fatal(err)
	}
	expected, err := filepath.Abs(root)
	if err != nil {
		t.Fatal(err)
	}
	if got != expected {
		t.Fatalf("expected command to run in %q, got %q", expected, got)
	}
}
