package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"sort"
)

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
	err = run(editor, result.File)
	if err != nil {
		fmt.Println(err)
	}
}

// Provides the status of a pattern-matched service
func status(pattern string, verbose bool) {
	cmd := "launchctl"
	args := []string{"list"}
	run(cmd, args...)
}

// Starts a pattern-matched service
func start(pattern string, write bool, force bool) {
	fmt.Println("Start", pattern, write, force)
	fmt.Println("NOT IMPLEMENTED")
}

// Stops a pattern-matched service
func stop(pattern string, write bool) {
	fmt.Println("Stop", pattern, write)
	fmt.Println("NOT IMPLEMENTED")
}

// Restarts a pattern-matched service
func restart(pattern string, write bool, force bool) {
	stop(pattern, write)
	start(pattern, write, force)
}
