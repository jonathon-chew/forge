package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	runcommand "github.com/jonathon-chew/forge/internal/runCommand"
)

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

	rootFolder := filepath.Join(currentPath, projectName)
	ErrMakingProjectFolder := os.Mkdir(rootFolder, os.ModePerm) // fails if path already exists, this is good hear, but error is ignored when creating the below
	if ErrMakingProjectFolder != nil {
		fmt.Print("error project already exists: ", ErrMakingProjectFolder)
		return
	}

	mainFileFolder := filepath.Join("cmd", projectName)

	folders := []string{"Archive", "cmd", "pkg", "internal", "doc", "scripts", "dist", mainFileFolder, filepath.Join("internal", "cli")}
	for _, folder := range folders {
		folderPath := filepath.Join(rootFolder, folder)
		ErrMakingFolder := os.Mkdir(folderPath, os.ModePerm)
		if ErrMakingFolder != nil && ErrMakingFolder != os.ErrExist {
			fmt.Print("error making folder: ", ErrMakingFolder)
			return
		}
	}

	files := []string{"README.md", "LICENSE", filepath.Join("scripts", "CICD.sh"), filepath.Join("scripts", "find_unused_exports.sh"), filepath.Join("scripts", "get_cmd_commands_for_help_file.zsh"), ".gitignore", filepath.Join("internal", "cli", "cli.go"), filepath.Join(mainFileFolder, "main.go")}
	for _, file := range files {
		filePath := filepath.Join(rootFolder, file)
		filePointer, ErrMakingFile := os.Create(filePath)
		if ErrMakingFile != nil {
			fmt.Print("error making file: ", ErrMakingFile)
			return
		}
		filePointer.Close()
	}

	// commands[4] = []string{"git", "config", "list", "--global"} // parse user.name to be in the LICENSE */

	commands := []runcommand.Command{
		{Name: "go", Command: []string{"go", "mod", "init", projectName}, Fatal: false, Description: "init a go project"},
		{Name: "git", Command: []string{"git", "init"}, Fatal: true, Description: "init a git project"},
		{Name: "git", Command: []string{"git", "add", "."}, Fatal: false, Description: "add everything and start tracking"},
		{Name: "git", Command: []string{"git", "commit", "-m", "BATMAN"}, Fatal: false, Description: "This commit has no parents"},
	}

	for _, commmand := range commands {
		if len(commmand.Command) > 0 {

			switch commmand.Description {
			default:
				runcommand.RunCommands(commmand, rootFolder)
			}
		}
	}
}
