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
func ls(pattern string, options []Option) {
	var plistNames []string
	for _, dir := range dirs() {
		for _, file := range plists(dir) {
			_, filename := filepath.Split(file)
			shortName := strings.TrimSuffix(filename, filepath.Ext(filename))

			if strings.Contains(shortName, pattern) != true {
				continue
			}

			// Long option is used
			if options[3].Value == true {
				plistNames = append(plistNames, file)
			} else {
				plistNames = append(plistNames, shortName)
			}
		}
	}
	sort.Strings(plistNames)
	for _, sortedName := range plistNames {
		fmt.Println(sortedName)
	}
}

func main() {

	options := make([]Option, 5)
	options[0] = Option{"force", "F", "force", false}
	options[1] = Option{"verbose", "v", "verbose", false}
	options[2] = Option{"write", "w", "write", false}
	options[3] = Option{"long", "l", "long", false}
	options[4] = Option{"symlink", "s", "symlink", false}

	ParseOptions(os.Args, options)

	if len(os.Args) < 2 {
		os.Exit(1)
	}
	command := os.Args[1]

	var pattern string

	if len(os.Args) > 2 {
		if string(os.Args[2][0]) != "-" {
			pattern = os.Args[2]
		}
	}

	if command == "ls" {
		ls(pattern, options)
	}
}
