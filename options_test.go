package main

import "testing"

func TestToggleOptions(t *testing.T) {

	options := map[string]Option{
		"force":   Option{"F", "force", false},
		"verbose": Option{"v", "verbose", false},
		"write":   Option{"w", "write", false},
		"long":    Option{"l", "long", false},
		"symlink": Option{"s", "symlink", false},
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
