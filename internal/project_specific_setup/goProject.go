package projectspecificsetup

import (
	"path/filepath"

	runcommand "github.com/jonathon-chew/forge/internal/runCommand"
)

func GoProject(projectName, rootFolder string) {
	mainFileFolder := filepath.Join("cmd", projectName)

	folders := []string{"Archive", "cmd", "pkg", "internal", "doc", "scripts", "dist", mainFileFolder, filepath.Join("internal", "cli")}

	makeFolders(rootFolder, folders)

	files := []string{"README.md", "LICENSE", filepath.Join("scripts", "CICD.sh"), filepath.Join("scripts", "find_unused_exports.sh"), filepath.Join("scripts", "get_cmd_commands_for_help_file.zsh"), ".gitignore", filepath.Join("internal", "cli", "cli.go"), filepath.Join(mainFileFolder, "main.go")}

	makeFiles(rootFolder, files)

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
