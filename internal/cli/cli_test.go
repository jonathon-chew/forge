package cli

import "testing"

func TestCliParsesProjectFlags(t *testing.T) {
	flags := Cli([]string{"--project-name", "demo", "--project-type", "go"})

	if flags.ProjectName != "demo" {
		t.Fatalf("expected project name %q, got %q", "demo", flags.ProjectName)
	}
	if flags.ProjectType != "go" {
		t.Fatalf("expected project type %q, got %q", "go", flags.ProjectType)
	}
}

func TestCliAcceptsShortProjectFlags(t *testing.T) {
	flags := Cli([]string{"-p", "demo", "-t", "python"})

	if flags.ProjectName != "demo" || flags.ProjectType != "python" {
		t.Fatalf("unexpected flags: %+v", flags)
	}
}
