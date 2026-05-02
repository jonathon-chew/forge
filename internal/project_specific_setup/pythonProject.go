package projectspecificsetup

import (
	"fmt"
	"path/filepath"
	"strings"

	utils "github.com/jonathon-chew/forge/internal/utils"
)

func PythonProject(projectName, rootFolder string) {

	mainFileFolder := filepath.Join("src", projectName)

	folders := []string{"Archive", "tests", "src", "docs", mainFileFolder}

	utils.MakeFolders(rootFolder, folders)

	// files := []string{"README.md", "LICENSE", "requirements.txt", filepath.Join(mainFileFolder, "__init__.py"), filepath.Join(mainFileFolder, "__main__.py"), filepath.Join(mainFileFolder, "moduel1.py"), filepath.Join(mainFileFolder, "moduel2.py"), filepath.Join("tests", "__init__.py"), filepath.Join("tests", "test1.py"), filepath.Join("tests", "test2.py")}
	files := []string{"README.md", "LICENSE", "requirements.txt", "setup.py"}
	mainFolderFiles := []string{"__init__.py", "__main__.py", "moduel_1.py", "moduel_2.py"}
	testFolderFiles := []string{"__init__.py", "test1.py", "test2.py"}

	for _, file := range mainFolderFiles {
		files = append(files, filepath.Join(mainFileFolder, file))
	}

	for _, file := range testFolderFiles {
		files = append(files, filepath.Join("tests", file))
	}

	utils.MakeFiles(rootFolder, files)

	var author, author_email string

	commands := []utils.Command{
		{Name: "git", Command: []string{"git", "init"}, Fatal: true, Description: "init a git project"},
		{Name: "git", Command: []string{"git", "add", "."}, Fatal: false, Description: "add everything and start tracking"},
		{Name: "git", Command: []string{"git", "commit", "-m", "BATMAN"}, Fatal: false, Description: "This commit has no parents"},
		{Name: "Virtual Enviornment", Command: []string{"python3", "-m", "venv", projectName + "_venv"}, Fatal: false, Description: "make a local venv"},
		{Name: "author", Command: []string{"git", "config", "list"}, Fatal: false, Description: "get the author"},
	}

	for _, commmand := range commands {
		if len(commmand.Command) > 0 {

			switch commmand.Name {
			case "author":
				output, stderr, err := utils.RunCommands(commmand, rootFolder)
				if err != nil && stderr != "" {
					continue
				}

				lines := strings.Split(output, "\n")
				for _, line := range lines {
					if strings.Contains(line, "user.name=") {
						author = strings.ReplaceAll(line, "user.name=", "")
					}
					if strings.Contains(line, "user.email=") {
						author_email = strings.ReplaceAll(line, "user.email=", "")
					}

				}
			default:
				utils.RunCommands(commmand, rootFolder)
			}
		}
	}

	var setup = fmt.Appendf(nil, `
from setuptools import find_packages
from setuptools import setup

setup (
    name="%s",
    version="",
    description="",
    author="%s",
    author_email="%s",
    url="",
    package_dir={"": "src"},
    packages=find_packages(where="src", exclude=("tests*",)),
)
	`, projectName, author, author_email)

	utils.WriteFile(filepath.Join(rootFolder, "setup.py"), setup)
}
