package projectspecificsetup

import (
	"bytes"
	"os"
	"path/filepath"
	"testing"
)

func Test_cli_builder(t *testing.T) {
	temp_dir, err := os.MkdirTemp(".", "*")
	if err != nil {
		t.Error("[ERROR]: Unable to make temp directory\n")
		t.Error(err.Error(), "\n")
		t.FailNow()
	}
	defer os.RemoveAll(temp_dir)

	cliPath := filepath.Join(temp_dir, "cli.go")
	cli_builder(cliPath)

	files, err := os.ReadDir(temp_dir)
	if err != nil {
		t.Error("[ERROR]: Trying to read directory\n")
		t.Error(err.Error(), "\n")
	}

	found := false
	for _, file := range files {
		if file.Name() == "cli.go" {
			found = true
			break
		}

		t.Log(file.Name(), "\n")
	}

	if !found {
		t.Error("[ERROR]: cli.go could not be found", "\n")
		t.FailNow()
	}

	content, err := os.ReadFile(filepath.Join(temp_dir, "cli.go"))
	if err != nil {
		t.Error("[ERROR]: Unable to open cli file", "\n")
		t.Error(err, "\n")

		t.FailNow()
	}

	if !bytes.Contains(content, []byte("package cli")) {
		t.Error("[ERROR]: Package was not declared", "\n")
	}
}

func TestMain(t *testing.T) {

	tmpdir, err := os.MkdirTemp(".", "*")
	if err != nil {
		t.Error(err, "\n")
		t.FailNow()
	}

	defer os.RemoveAll(tmpdir)

	GoProject("eg", tmpdir)
}
