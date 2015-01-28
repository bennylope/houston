package main

import (
	"fmt"
	"os"
	"os/exec"
	"os/user"
	"path/filepath"
	"strings"
)

// Returns a list of all relevant directories to search for plist files
// TODO handle root user, make dir list appendable?
func dirs() []string {
	// Just using the tilde ~/ does not work!
	current, _ := user.Current()
	return []string{current.HomeDir + "/Library/LaunchAgents", "/Library/LaunchAgents"}
}

// Returns a list of all plist files for a given directory
func plists(dir string) []string {
	globpath := dir + "/*.plist"
	files, _ := filepath.Glob(globpath)
	return files
}

func getAllServices() Services {
	var ServicesList Services
	for _, dir := range dirs() {
		for _, file := range plists(dir) {
			_, filename := filepath.Split(file)
			shortName := strings.TrimSuffix(filename, filepath.Ext(filename))
			ServicesList.AddService(Service{shortName, file})
		}
	}
	return ServicesList
}

// Wraps the exec.Command function using stdin and stdout so that the command
// is executed on the user's shell
func run(name string, arg ...string) error {
	c := exec.Command(name, arg...)
	c.Stdin = os.Stdin
	c.Stdout = os.Stdout
	err := c.Run()
	return err
}

// Takes 1 or more commands that are expected to be executed by piping together
// in a stdout/stdin train.
func pipeCommands(commands ...exec.Cmd) error {
	last := len(commands) - 1

	if last == -1 {
		fmt.Println("Error, no commands sent to pipeCommands function")
		os.Exit(1)
	}

	// The caboose always uses os.Stdout
	commands[last].Stdout = os.Stdout

	// If only one command was provided just run it.
	if len(commands) == 1 {
		c := commands[0]
		err := c.Run()
		return err
	}

	// Chain the the output of each command to the input of the subsequent
	// command.
	for i, _ := range commands[:last] {
		out, err := commands[i].StdoutPipe()
		if err != nil {
			return err
		}
		commands[i+1].Stdin = out
	}

	// Start all but the first command
	for i, _ := range commands {
		if i == 0 {
			continue
		}
		if err := commands[i].Start(); err != nil {
			return err
		}
	}

	// The first command is run (start + wait)
	commands[0].Run()

	// Wait all but the first command
	for i, _ := range commands {
		if i == 0 {
			continue
		}
		if err := commands[i].Wait(); err != nil {
			return err
		}
	}

	return nil
}
