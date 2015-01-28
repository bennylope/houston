package main

import (
	"fmt"
	"os"
)

func main() {

	options := map[string]*Option{
		"force":   {"F", "force", false},
		"verbose": {"v", "verbose", false},
		"write":   {"w", "write", false},
		"long":    {"l", "long", false},
		"symlink": {"s", "symlink", false},
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
		} else {
			// TODO catch index error here - not very nice right now
			pattern = os.Args[3]
		}
	}

	// Better way?
	if command == "list" || command == "ls" {
		ls(pattern, options["long"].Value)
	} else if command == "show" {
		show(pattern)
	} else if command == "edit" {
		edit(pattern)
	} else if command == "status" {
		status(pattern, options["verbose"].Value)
	} else if command == "start" {
		start(pattern, options["write"].Value, options["force"].Value)
	} else if command == "stop" {
		stop(pattern, options["write"].Value)
	} else if command == "restart" {
		restart(pattern, options["write"].Value, options["force"].Value)
	} else if command == "install" {
		install(pattern, options["symlink"].Value)
	} else if command == "uninstall" || command == "rm" {
		uninstall(pattern)
	} else {
		fmt.Println("No such command '", pattern, "'")
		// TODO show help
		os.Exit(1)
	}
}
