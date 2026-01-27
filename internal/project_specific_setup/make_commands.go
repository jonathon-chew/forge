package projectspecificsetup

import (
	"fmt"
	"os"
	"path/filepath"
)

func makeFolders(rootFolder string, folders []string) {
	for _, folder := range folders {
		folderPath := filepath.Join(rootFolder, folder)
		ErrMakingFolder := os.Mkdir(folderPath, os.ModePerm)
		if ErrMakingFolder != nil && ErrMakingFolder != os.ErrExist {
			fmt.Print("error making folder: ", ErrMakingFolder)
			return
		}
	}
}

func makeFiles(rootFolder string, files []string) {
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
