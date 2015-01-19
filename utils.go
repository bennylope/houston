package main

import (
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
	return []string{"/Library/LaunchAgents", current.HomeDir + "/Library/LaunchAgents"}
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
