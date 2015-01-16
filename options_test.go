package main

import "testing"

func TestToggleOptions(t *testing.T) {

	options := make([]Option, 5)
	options[0] = Option{"force", "F", "force", false}
	options[1] = Option{"verbose", "v", "verbose", false}
	options[2] = Option{"write", "w", "write", false}
	options[3] = Option{"long", "l", "long", false}
	options[4] = Option{"symlink", "s", "symlink", false}

	args := []string{"-l", "-wF", "--verbose"}
	ParseOptions(args, options)

	if options[0].Value == false {
		t.Error("Option should be true: ", options[0].Name)
	}
	if options[1].Value == false {
		t.Error("Option should be true: ", options[1].Name)
	}
	if options[2].Value == false {
		t.Error("Option should be true: ", options[2].Name)
	}
	if options[3].Value == false {
		t.Error("Option should be true: ", options[3].Name)
	}
	if options[4].Value == true {
		t.Error("Option should be false: ", options[4].Name)
	}
}
