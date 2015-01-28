package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"os/user"
	"path/filepath"
	"sort"
	//"strings"
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
		fmt.Println("Multiple daemons found matching '" + pattern + "'. You need to be more specific. Matches found are:")
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
	allServices := getAllServices()
	result, err := allServices.Get(pattern)
	if err != nil {
		fmt.Println("Multiple daemons found matching '" + pattern + "'. You need to be more specific. Matches found are:")
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
	var cmd_list = []exec.Cmd{*exec.Command("launchctl", "list")}
	if pattern != "" {
		cmd_list = append(cmd_list, *exec.Command("grep", "-i", pattern))
	}
	pipeCommands(cmd_list...)
}

// Starts a pattern-matched service
func start(pattern string, write bool, force bool) {
	allServices := getAllServices()
	result, err := allServices.Get(pattern)
	if err != nil {
		fmt.Println("Multiple daemons found matching '" + pattern + "'. You need to be more specific. Matches found are:")
		ls(pattern, false)
		os.Exit(1)
	}
	var f string
	if write || force {
		f = "-"
	}
	if write {
		f = f + "w"
	}
	if force {
		f = f + "F"
	}

	err = run("launchctl", "load", "-w", result.File)
	if err != nil {
		fmt.Println("Error starting", result.Name, "-", err)
		os.Exit(1)
	}
	fmt.Println("started", result.Name)
}

// Stops a pattern-matched service
func stop(pattern string, write bool) {
	allServices := getAllServices()
	result, err := allServices.Get(pattern)
	if err != nil {
		fmt.Println("Multiple daemons found matching '" + pattern + "'. You need to be more specific. Matches found are:")
		ls(pattern, false)
		os.Exit(1)
	}
	var f string
	if write {
		f = "-w"
	} else {
		f = ""
	}
	err = run("launchctl", "unload", f, result.File)
	if err != nil {
		fmt.Println("Error stopping", result.Name, "-", err)
		os.Exit(1)
	}
	fmt.Println("stopped", result.Name)
}

// Restarts a pattern-matched service
func restart(pattern string, write bool, force bool) {
	stop(pattern, write)
	start(pattern, write, force)
}

// Installs a service plist file
func install(file string, symlink bool) {
	// TODO check to see if file is installed in one of our dirs first
	filename := filepath.Base(file)
	if symlink {
		// Might need to go back to to exec.Cmd b/c os.Symlink cannot force a
		// symlink
		currentUser, _ := user.Current()
		target := currentUser.HomeDir + "/Library/LaunchAgents/" + filename
		err := os.Symlink(file, target)
		if err != nil {
			fmt.Println("Error with", file)
			fmt.Println(err)
			os.Exit(1)
		}
	} else {
		err := run("cp", file, filename)
		if err != nil {
			fmt.Println("Error copying file:", err)
			os.Exit(1)
		}
	}
	fmt.Println("Installed", file)
}

// Removes a service plist file
func uninstall(pattern string) {
	// This code here is being repeated in several places.
	allServices := getAllServices()
	result, err := allServices.Get(pattern)
	if err != nil {
		fmt.Println("Multiple daemons found matching '" + pattern + "'. You need to be more specific. Matches found are:")
		ls(pattern, false)
		os.Exit(1)
	}
	fmt.Println(result.File)
	err = os.Remove(result.File)
	if err != nil {
		fmt.Println("Error!", err)
		os.Exit(1)
	}
	fmt.Println("Uninstalled", pattern)
}
