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
// This was reqired before making the switch to a map of *Option rather than
// Option, as it's not possible to assign to the value of a struct inside of
// map. If the values are pointers to structs, then everything is peachy.
// Instead of switching to pointers first, I just replaced the value with a
// copy of the struct that was first updated with the appropriate boolean
// Value.
func (o *Option) copy() Option {
	return Option{o.Short, o.Long, o.Value}
}

// Compares the short and long flags in a slice of Options to see if any of the
// flags were used. If so then the Value for each used flag is toggled to true
func ParseOptions(args []string, options map[string]*Option) {

	// Slices for aggregating one letter (short) options and double dash full
	// word options
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

	for k, opt := range options {
		for _, sOption := range shortOptions {
			if opt.Short == sOption {
				options[k].Value = true
			}
		}
		for _, lOption := range longOptions {
			if opt.Long == lOption {
				options[k].Value = true
			}
		}
	}
}
