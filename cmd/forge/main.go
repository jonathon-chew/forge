package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

type Command struct {
	Name        string
	Command     []string
	Fatal       bool
	Description string
}

func main() {

	fmt.Print("What is the name of your project?\n")
	// (#1) TODO: Read from stdin to get a project name
	reader := bufio.NewReader(os.Stdin)
	projectName, ErrGettingProjectName := reader.ReadString('\n')
	if ErrGettingProjectName != nil {
		fmt.Print("error getting project name: ", ErrGettingProjectName)
		return
	}

	projectName = strings.TrimSpace(projectName)

	currentPath, ErrGettingCurrentPath := os.Getwd()
	if ErrGettingCurrentPath != nil {
		fmt.Print("error getting current path: ", ErrGettingCurrentPath)
		return
	}

	projectPath := filepath.Join(currentPath, projectName)
	ErrMakingProjectFolder := os.Mkdir(projectPath, os.ModePerm) // fails if path already exists, this is good hear, but error is ignored when creating the below
	if ErrMakingProjectFolder != nil {
		fmt.Print("error project already exists: ", ErrMakingProjectFolder)
		return
	}

	projectPathName := filepath.Join("cmd", projectName)

	folders := []string{"Archive", "cmd", "pkg", "internal", "doc", "scripts", "dist", projectPathName}
	for _, folder := range folders {
		folderPath := filepath.Join(projectPath, folder)
		ErrMakingFolder := os.Mkdir(folderPath, os.ModePerm)
		if ErrMakingFolder != nil && ErrMakingFolder != os.ErrExist {
			fmt.Print("error making folder: ", ErrMakingFolder)
			return
		}
	}

	files := []string{"README.md", "LICENSE", "scripts/CICD.sh", "scripts/find_unused_exports.sh", "scripts/get_cmd_commands_for_help_file.zsh", ".gitignore"}
	for _, file := range files {
		filePath := filepath.Join(projectPath, file)
		filePointer, ErrMakingFile := os.Create(filePath)
		if ErrMakingFile != nil {
			fmt.Print("error making file: ", ErrMakingFile)
			return
		}
		filePointer.Close()
	}

	//	commands := make([][]string, 5)
	/* commands[0] = []string{"go", "mod", "init", projectName}
	commands[1] = []string{"git", "init"}
	commands[2] = []string{"git", "add", "."}
	commands[3] = []string{"git", "commit", "-m", "BATMAN"}
	commands[4] = []string{"git", "config", "list", "--global"} // parse user.name to be in the LICENSE */

	commands := []Command{
		{"go", []string{"go", "mod", "init", projectName}, false, "init a go project"},
		{"git", []string{"git", "init"}, true, "init a git project"},
		{"git", []string{"git", "add", "."}, false, "add everything and start tracking"},
		{"git", []string{"git", "commit", "-m", "BATMAN"}, false, "This commit has no parents"},
		{"git", []string{"git", "config", "list", "--global"}, false, "get a user name"},
	}

	for _, commmand := range commands {
		if len(commmand.Command) > 0 {
			cmd := exec.Command(commmand.Command[0], commmand.Command[1:]...)
			cmd.Dir = projectPath

			var Stdout bytes.Buffer
			var Stderr bytes.Buffer

			cmd.Stdout = &Stdout
			cmd.Stderr = &Stderr

			ErrRunningCommand := cmd.Run()
			if ErrRunningCommand != nil {
				fmt.Print("error running command: ", commmand.Name, ErrRunningCommand)
				return
			}

			if Stderr.Len() > 0 && commmand.Fatal {
				fmt.Print("error: from the command output: ", Stderr.String())
			}
		}
	}
}
