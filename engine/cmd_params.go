package engine

import "flag"

var flagCRT = flag.Bool("crt", false, "enable the CRT effect")

// CMDArguments represents all command line arguments
type CMDArguments struct {
	UseCRT bool
}

// ParseCmdArguments parses cmd arguments and returns them in a structure
func ParseCmdArguments() CMDArguments {
	flag.Parse()
	return CMDArguments{
		UseCRT: *flagCRT,
	}
}
