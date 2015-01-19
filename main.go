package main

import (
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
		}
	}

	if command == "ls" {
		ls(pattern, options["long"].Value)
	} else if command == "show" {
		show(pattern)
	} else if command == "edit" {
		edit(pattern)
	} else if command == "status" {
		status(pattern, options["verbose"].Value)
	}
}
