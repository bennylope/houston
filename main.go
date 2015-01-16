package main

import (
	"fmt"
	"os"
	"os/user"
	"path/filepath"
	"sort"
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

// Prints a list of all
func ls() {
	var plistNames []string
	for _, dir := range dirs() {
		for _, file := range plists(dir) {
			if true == false {
				plistNames = append(plistNames, file)
			} else {
				_, filename := filepath.Split(file)
				plistNames = append(plistNames, strings.TrimSuffix(filename, filepath.Ext(filename)))
			}
		}
	}
	sort.Strings(plistNames)
	for _, sortedName := range plistNames {
		fmt.Println(sortedName)
	}
}

func main() {

	if len(os.Args) < 2 {
		os.Exit(1)
	}
	command := os.Args[1]
	if command == "ls" {
		ls()
	}
}
