package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
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

// Prints a list of all matching services
func ls(pattern string, long bool) {
	var results []string
	allServices := getAllServices()
	filteredServices := allServices.Filter(pattern)

	if long == true {
		results = filteredServices.GetFiles()
	} else {
		results = filteredServices.GetNames()
	}

	sort.Strings(results)
	for _, sortedName := range results {
		fmt.Println(sortedName)
	}
}

// Shows the text of a daemon plist file
// If no files are matched it outputs nothing. If multiple matches are found it
// outputs a warning message and lists the daemon file names
func show(pattern string) {
	allServices := getAllServices()
	result, err := allServices.Get(pattern)
	if err != nil {
		fmt.Println("Multiple daemons found matching 'com'. You need to be more specific. Matches found are:")
		ls(pattern, false)
		os.Exit(1)
	}
	if result.File == "" {
		os.Exit(0)
	}

	file, err := ioutil.ReadFile(result.File)
	if err != nil {
		fmt.Println("Error opening file:", err)
		os.Exit(1)
	}
	fmt.Println(string(file))
}

// Allow the user to edit a service's plist file using their editor.
// Searches for the one match and if found uses the EDITOR defined in the
// environment to edit the file.
func edit(pattern string) {
	fmt.Println(pattern)
	allServices := getAllServices()
	result, err := allServices.Get(pattern)
	if err != nil {
		fmt.Println("Multiple daemons found matching 'com'. You need to be more specific. Matches found are:")
		ls(pattern, false)
		os.Exit(1)
	}
	if result.File == "" {
		os.Exit(0)
	}
	editor := os.Getenv("EDITOR")
	if editor == "" {
		fmt.Println("EDITOR environment variable is not set")
		os.Exit(1)
	}
	cmd := exec.Command(editor, result.File)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	err = cmd.Run()
	if err != nil {
		fmt.Println(err)
	}
}

func main() {

	options := map[string]Option{
		"force":   Option{"F", "force", false},
		"verbose": Option{"v", "verbose", false},
		"write":   Option{"w", "write", false},
		"long":    Option{"l", "long", false},
		"symlink": Option{"s", "symlink", false},
	}

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
		ls(pattern, options["long"].Value)
	} else if command == "show" {
		show(pattern)
	} else if command == "edit" {
		edit(pattern)
	}
}
