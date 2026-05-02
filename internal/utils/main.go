package utils

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	aphrodite "github.com/jonathon-chew/Aphrodite"
)

type Command struct {
	Command     []string
	Name        string
	Description string
	Fatal       bool
}

/*
Returns stdout, stderr and error with the function in that order
*/
func RunCommands(command Command, folderPath string) (string, string, error) {
	cmd := exec.Command(command.Command[0], command.Command[1:]...)
	cmd.Dir = folderPath

	var Stdout bytes.Buffer
	var Stderr bytes.Buffer

	cmd.Stdout = &Stdout
	cmd.Stderr = &Stderr

	ErrRunningCommand := cmd.Run()
	if ErrRunningCommand != nil {
		fmt.Print("error running command: ", command.Name, ErrRunningCommand)
		return "", "", ErrRunningCommand
	}

	if Stderr.Len() > 0 && command.Fatal {
		fmt.Print("error: from the command output: ", Stderr.String())
	}
	return Stdout.String(), Stderr.String(), nil
}

func WriteFile(filename string, filecontents []byte) {
	err := os.WriteFile(filename, filecontents, os.ModePerm)
	if err != nil {
		aphrodite.PrintError(err.Error())
	}
}

func MakeFolders(rootFolder string, folders []string) {
	for _, folder := range folders {
		folderPath := filepath.Join(rootFolder, folder)
		ErrMakingFolder := os.Mkdir(folderPath, os.ModePerm)
		if ErrMakingFolder != nil && ErrMakingFolder != os.ErrExist {
			fmt.Print("error making folder: ", ErrMakingFolder)
			return
		}
	}
}

func MakeFiles(rootFolder string, files []string) {
	for _, file := range files {
		filePath := filepath.Join(rootFolder, file)
		filePointer, ErrMakingFile := os.Create(filePath)
		if ErrMakingFile != nil {
			fmt.Print("error making file: ", ErrMakingFile)
			return
		}
		filePointer.Close()
	}
}
