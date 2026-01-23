package cli

import aphrodite "github.com/jonathon-chew/Aphrodite"

type Flags struct {
	HelpMenu string
	Version  string
}

func Cli(commands []string) {

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
			aphrodite.PrintInfo("version: v0.1.0")
		}
	}
}
