package main

import "strings"

// Would have used a single array/slice of flags instead of Short and Long
// *but* it seemed easier to create the Options using this method - a lot easier
// - and it should make comparing different flag types simpler, too.
type Option struct {
	Name  string // Identifying name
	Short string // Long command line flag
	Long  string // Short option-style flag
	Value bool   // Value even for bool, so a list of possible flags can be set
	// with values based on whether they're present or not
}

// Compares the short and long flags in a slice of Options to see if any of the
// flags were used. If so then the Value for each used flag is toggled to true
func ParseOptions(args []string, options []Option) {
	var shortOptions []string
	var longOptions []string
	for _, arg := range args {
		if string(arg[0]) != "-" {
			continue
		}
		if string(arg[1]) == "-" {
			longOptions = append(longOptions, arg[2:])
		} else {
			for _, flag := range strings.Split(arg[1:], "") {
				shortOptions = append(shortOptions, flag)
			}
		}
	}
	for i, option := range options {
		// If you use `_, option` from the range statement and just use the
		// list item, it won't be set in the original slice.
		for _, sOption := range shortOptions {
			if option.Short == sOption {
				options[i].Value = true
			}
		}
		for _, lOption := range longOptions {
			if option.Long == lOption {
				options[i].Value = true
			}
		}
	}
}
