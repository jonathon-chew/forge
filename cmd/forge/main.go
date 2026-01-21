package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

// Folders
// Archive / cmd / pkg / internal / doc / scripts / dist

// Files
// LICSENCE
// README.md
// scripts/CICD.sh
// scripts/find_unused_exports.sh
// script/get_cmd_commands_for_help_file.zsh
// cmd / <project-name> / main.go >> hello world?

// Commands
// go mod init <project-name>
// git init
// git add .
// git commit -m "BATMAN"
// github repo create?

func main() {

	// (#1) TODO: Read from stdin to get a project name
	projectName := "forge"
	currentPath, ErrGettingCurrentPath := os.Getwd()
	if ErrGettingCurrentPath != nil {
		fmt.Print("error: ", ErrGettingCurrentPath)
		return
	}

	projectPath := filepath.Join(currentPath, projectName)
	ErrMakingProjectFolder := os.Mkdir(projectPath, os.ModePerm)
	if ErrMakingProjectFolder != nil {
		fmt.Print("error: ", ErrMakingProjectFolder)
		return
	}

	projectPathName := filepath.Join("cmd", projectName)

	folders := []string{"Archive", "cmd", "pkg", "internal", "doc", "scripts", "dist", projectPathName}
	for _, folder := range folders {
		folderPath := filepath.Join(projectPath, folder)
		ErrMakingFolder := os.Mkdir(folderPath, os.ModePerm)
		if ErrMakingFolder != nil {
			fmt.Print("error: ", ErrMakingFolder)
			return
		}
	}

	files := []string{"README.md", "LICENSE", "scripts/CICD.sh", "scripts/find_unused_exports.sh", "scripts/get_cmd_commands_for_help_file.zsh"}
	for _, file := range files {
		filePath := filepath.Join(projectPath, file)
		_, ErrMakingFile := os.Create(filePath)
		if ErrMakingFile != nil {
			fmt.Print("error: ", ErrMakingFile)
			return
		}
	}

	commands := make([][]string, 5)
	commands[0] = []string{"go", "mod", "init", projectName}
	commands[1] = []string{"git", "init"}
	commands[2] = []string{"git", "add", "."}
	commands[3] = []string{"git", "commit", "-m", "BATMAN"}
	commands[4] = []string{"go", "mod", "init", projectName}

	for _, commmand := range commands {
		cmd := exec.Command(commmand[0], commmand...)
		cmd.Dir = projectPath

		// Stdout := &cmd.Stdout
		Stderr := &cmd.Stderr

		if Stderr != nil {
			fmt.Print("error: ", Stderr)
		}
	}
}
