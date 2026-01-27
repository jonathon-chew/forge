package cli

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	aphrodite "github.com/jonathon-chew/Aphrodite"
)

type Flags struct {
	HelpMenu    string
	Version     string
	ProjectName string
	ProjectType string
}

var version string

func StdInProjectName() string {

	fmt.Print("What is the name of your project?\n")
	// (#1) TODO: Read from stdin to get a project name
	reader := bufio.NewReader(os.Stdin)
	projectName, ErrGettingProjectName := reader.ReadString('\n')
	if ErrGettingProjectName != nil {
		fmt.Print("error getting project name: ", ErrGettingProjectName)
		return ""
	}
	projectName = strings.TrimSpace(projectName)

	return projectName

}

func StdInProjectType() string {

	fmt.Print("What type of project is it??\n")
	// (#1) TODO: Read from stdin to get a project name
	reader_type := bufio.NewReader(os.Stdin)
	projectType, ErrGettingProjectName := reader_type.ReadString('\n')
	if ErrGettingProjectName != nil {
		fmt.Print("error getting project type: ", ErrGettingProjectName)
		return ""
	}
	projectType = strings.TrimSpace(projectType)

	return projectType

}

func Cli(commands []string) Flags {

	var flags Flags

	for i := 0; i < len(commands); i++ {
		command := commands[i]

		switch command {
		default:
			aphrodite.PrintError("[ERROR]: did not recognise command " + command)
		case "--help", "-h":
			aphrodite.PrintInfo(`
			--help or -h
			To access the help menu

			--version or -v
			To see the version
			`)
		case "--version", "-v":
			aphrodite.PrintInfo("version: " + version)
		case "--project-name", "--projectname", "-pn", "-p":
			if len(commands)+1 > i {
				flags.ProjectName = commands[i+1]
				i++
			} else {
				flags.ProjectName = StdInProjectName()
			}
		case "--project-type", "--projecttype", "-pt", "-t":
			if len(commands)+1 > i {
				flags.ProjectType = commands[i+1]
				i++
			} else {
				flags.ProjectType = StdInProjectType()
			}
		}
	}

	if flags.ProjectName == "" {
		flags.ProjectName = StdInProjectName()
	}
	if flags.ProjectType == "" {
		flags.ProjectType = StdInProjectType()
	}

	return flags
}
