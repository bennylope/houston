package main

// Would have used a single array/slice of flags instead of Short and Long
// *but* it seemed easier to create the Flags using this method - a lot easier
// - and it should make comparing different flag types simpler, too.
type Flag struct {
	Name  string // Identifying name
	Short string // Long command line flag
	Long  string // Short option-style flag
	Value bool   // Value even for bool, so a list of possible flags can be set
	// with values based on whether they're present or not
}


// Compares the short and long flags in a slice of Flags to see if any of the
// flags were used. If so then the Value for each used flag is toggled to true
// TODO make an inverted map of flags to options (?)
func parseOptions (args []string, options []Flag) {
    var shortFlags []string
    var longFlags []string
    for _, arg := range args {
        if arg[0] != "-" {
            continue
        }
        if arg[1] == "-" {
            longFlags = append(longFlags, arg)
        }
    }
}
