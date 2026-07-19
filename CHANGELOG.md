# CHANGELOG

	## v0.2.6, tag: v0.2.5 

		### NEW
		1. entry point as needed to build
		### UPDATES
		1. repo clean up

	## v0.2.4 

		### UPDATES
		1. upgrading Aphrodite so only outputs ANSI if output is a terminal
		### MISC
		1. adjustments to setup.py template as current one doesn't work as inteded

	## v0.2.3 

		### NEW
		1. write setup.py for python projects; update: run command to generic utils pacakge
		### MISC
		1. minor update to setup.py contents and location; refactor: move 2 util funcs to util package

	## v0.2.2 

		### UPDATES
		1. venv enviornment folder labeled as such - not just the project name

	## v0.2.1 

		### UPDATES
		1. error message upon unrecognised command to check if command has a - or not

	## v0.2.0 

		### NEW
		1. clean up on failing to recognise project name

	## v0.1.6 

		### NEW
		1. python option for projects, allow for arguments/flags; refactor: project logic into it's own package and made commonents re-useable
		### UPDATES
		1. markdown rendering
		### MISC
		1. making a README.md for forge

	## v0.1.5 

		### UPDATES
		1. reimplimented filepath function calls so works on all devices again

	## v0.1.4 

		### UPDATES
		1. file path to full file path for main.go which was extracted out of file list

	## v0.1.3 

		### MISC
		1. file path for internal cli

	## v0.1.2 

		### UPDATES
		1. modified cli file making for the listed file to only be the releative path join as the full path join happens inside the loop anyways - specifically for the main file

	## v0.1.1 

		### UPDATES
		1. modified cli file making for the listed file to only be the releative path join as the full path join happens inside the loop anyways

	## v0.1.0 

		### UPDATES
		1. file path in list for internal CLI to fix issue #1; reverted refactor to CICD.sh as error message getting ./scripts/CICD.sh: line 76: y: command not found \n ./scripts/CICD.sh: line 76: y: command not found