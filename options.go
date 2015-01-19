package main

import "strings"

// Would have used a single array/slice of flags instead of Short and Long
// *but* it seemed easier to create the Options using this method - a lot easier
// - and it should make comparing different flag types simpler, too.
type Option struct {
	Short string // Long command line flag
	Long  string // Short option-style flag
	Value bool   // Value even for bool, so a list of possible flags can be set
	// with values based on whether they're present or not
}

// Returns a copy of an option
func (o *Option) copy() Option {
	return Option{o.Short, o.Long, o.Value}
}

// Compares the short and long flags in a slice of Options to see if any of the
// flags were used. If so then the Value for each used flag is toggled to true
func ParseOptions(args []string, options map[string]Option) {
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
	var c Option
	for k, opt := range options {
		for _, sOption := range shortOptions {
			if opt.Short == sOption {
				// It's not possible to assign a struct value for a struct within a
				// map, not directly at least. You can have the map be a map of
				// *pointers* to the struct type, and while that'd be more efficient
				// and probably more idiomatic, I wanted to try this first.
				c = opt.copy()
				c.Value = true
				options[k] = c
			}
		}
		for _, lOption := range longOptions {
			if opt.Long == lOption {
				c = opt.copy()
				c.Value = true
				options[k] = c
			}
		}
	}
}
