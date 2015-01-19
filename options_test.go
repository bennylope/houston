package main

import "testing"

func TestToggleOptions(t *testing.T) {

	// It's not possible to assign a struct value for a struct within a
	// map, not directly at least. You can have the map be a map of
	// *pointers* to the struct type. Hence the switch to a map of struct
	// pointers here.
	options := map[string]*Option{
		"force":   {"F", "force", false},
		"verbose": {"v", "verbose", false},
		"write":   {"w", "write", false},
		"long":    {"l", "long", false},
		"symlink": {"s", "symlink", false},
	}

	args := []string{"-l", "-wF", "--verbose"}
	ParseOptions(args, options)

	if options["force"].Value == false {
		t.Error("'force' option should be false")
	}
	if options["verbose"].Value == false {
		t.Error("'verbose' option should be false")
	}
	if options["write"].Value == false {
		t.Error("'write' option should be false")
	}
	if options["long"].Value == false {
		t.Error("'long' option should be false")
	}
	if options["symlink"].Value == true {
		t.Error("'symlink' option should be true")
	}
}
