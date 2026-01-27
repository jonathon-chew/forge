package projectspecificsetup

import (
	"path/filepath"

	runcommand "github.com/jonathon-chew/forge/internal/runCommand"
)

func PythonProject(projectName, rootFolder string) {

	mainFileFolder := filepath.Join("src", projectName)

	folders := []string{"Archive", "tests", "src", "docs", mainFileFolder}

	makeFolders(rootFolder, folders)

	// files := []string{"README.md", "LICENSE", "requirements.txt", filepath.Join(mainFileFolder, "__init__.py"), filepath.Join(mainFileFolder, "__main__.py"), filepath.Join(mainFileFolder, "moduel1.py"), filepath.Join(mainFileFolder, "moduel2.py"), filepath.Join("tests", "__init__.py"), filepath.Join("tests", "test1.py"), filepath.Join("tests", "test2.py")}
	files := []string{"README.md", "LICENSE", "requirements.txt"}
	mainFolderFiles := []string{"__init__.py", "__main__.py", "moduel_1.py", "moduel_2.py"}
	testFolderFiles := []string{"__init__.py", "test1.py", "test2.py"}

	for _, file := range mainFolderFiles {
		files = append(files, filepath.Join(mainFileFolder, file))
	}

	for _, file := range testFolderFiles {
		files = append(files, filepath.Join("tests", file))
	}

	makeFiles(rootFolder, files)

	commands := []runcommand.Command{
		{Name: "git", Command: []string{"git", "init"}, Fatal: true, Description: "init a git project"},
		{Name: "git", Command: []string{"git", "add", "."}, Fatal: false, Description: "add everything and start tracking"},
		{Name: "git", Command: []string{"git", "commit", "-m", "BATMAN"}, Fatal: false, Description: "This commit has no parents"},
		{Name: "Virtual Enviornment", Command: []string{"python3", "-m", "venv", projectName}, Fatal: false, Description: "make a local venv"},
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
