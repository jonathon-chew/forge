package projectspecificsetup

import (
	"path/filepath"

	nonprojectspecficsetup "github.com/jonathon-chew/forge/internal/non_project_specfic_setup"
	utils "github.com/jonathon-chew/forge/internal/utils"
)

func cli_builder(location string) {
	content := []byte(`package cli

import (
    "fmt"
)

type Flags struct {
    HelpMenu string
    Version  string
}

var version string

func Cli(commands []string) Flags {
    var flags Flags

    for _, command := range commands {
        switch command {
        default:
            fmt.Print("[ERROR]: did not recognise command " + command + "\n")

            if len(command) == 0 || command[0] != '-' {
                fmt.Print("please pass in a flag with - " + command + "\n")
            }

        case "--help", "-h":
            fmt.Println("--help or -h\nTo access the help menu\n\n--version or -v\nTo see the version\n")

        case "--version", "-v":
            fmt.Println("version: " + version)
        }
    }

    return flags
}
`)

	utils.WriteFile(location, content)
}

func mainbuilder(location string) {
	content := []byte(`package main

import (
    "fmt"
)

func main(){
	fmt.Println("Hello World")
}
`)

	utils.WriteFile(filepath.Join(location, "main.go"), content)
}

func GoProject(projectName, rootFolder string) {
	mainFileFolder := filepath.Join("cmd", projectName)

	folders := []string{"Archive", "cmd", "pkg", "internal", "doc", "scripts", "dist", mainFileFolder, filepath.Join("internal", "cli")}

	utils.MakeFolders(rootFolder, folders)

	files := []string{"README.md", "LICENSE", filepath.Join("scripts", "CICD.sh"), filepath.Join("scripts", "find_unused_exports.sh"), filepath.Join("scripts", "get_cmd_commands_for_help_file.zsh"), ".gitignore", filepath.Join("internal", "cli", "cli.go"), filepath.Join(mainFileFolder, "main.go")}

	utils.MakeFiles(rootFolder, files)

	// commands[4] = []string{"git", "config", "list", "--global"} // parse user.name to be in the LICENSE */

	commands := []utils.Command{
		{Name: "go", Command: []string{"go", "mod", "init", projectName}, Fatal: false, Description: "init a go project"},
		{Name: "git", Command: []string{"git", "init"}, Fatal: true, Description: "init a git project"},
		{Name: "git", Command: []string{"git", "add", "."}, Fatal: false, Description: "add everything and start tracking"},
		{Name: "git", Command: []string{"git", "commit", "-m", "BATMAN"}, Fatal: false, Description: "This commit has no parents"},
	}

	for _, commmand := range commands {
		if len(commmand.Command) > 0 {

			switch commmand.Description {
			default:
				utils.RunCommands(commmand, rootFolder)
			}
		}
	}

	mainbuilder(filepath.Join(rootFolder, "src", projectName))
	cli_builder(filepath.Join(rootFolder, "internal", "cli", "cli.go"))
	nonprojectspecficsetup.CICD_Script(rootFolder)
}
