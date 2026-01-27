package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/jonathon-chew/forge/internal/cli"
	projectspecificsetup "github.com/jonathon-chew/forge/internal/project_specific_setup"
)

var flags cli.Flags

func main() {
	if len(os.Args) >= 2 {
		flags = cli.Cli(os.Args[1:])
	} else {
		flags.ProjectName = cli.StdInProjectName()
		flags.ProjectType = cli.StdInProjectType()
	}

	// default?

	currentPath, ErrGettingCurrentPath := os.Getwd()
	if ErrGettingCurrentPath != nil {
		fmt.Print("error getting current path: ", ErrGettingCurrentPath)
		return
	}

	rootFolder := filepath.Join(currentPath, flags.ProjectName)
	ErrMakingProjectFolder := os.Mkdir(rootFolder, os.ModePerm) // fails if path already exists, this is good hear, but error is ignored when creating the below
	if ErrMakingProjectFolder == os.ErrExist {
		fmt.Print("error project already exists: ", ErrMakingProjectFolder)
		return
	} else if ErrMakingProjectFolder != nil {
		fmt.Print("error making project folder: ", ErrMakingProjectFolder)
		return
	}

	switch flags.ProjectType {
	case "go", "golang":
		projectspecificsetup.GoProject(flags.ProjectName, rootFolder)
	case "py", "python":
		projectspecificsetup.PythonProject(flags.ProjectName, rootFolder)
	// case "sql", "data":

	default:
		fmt.Println("error: the project type ", flags.ProjectType, " has not been recognised")
		ErrCleaningUp := os.Remove(rootFolder)
		if ErrCleaningUp != nil {
			fmt.Println("error: unable to remove the root folder")
		}
		return
	}
}
