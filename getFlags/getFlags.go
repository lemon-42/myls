package getflags

import (
	"os"
	"strings"
)

type Flags struct {
	Recursive       bool
	LongFiles       bool
	HiddenDirectory bool
	Reverse         bool
	ModifiedTime    bool
}

func ParseFlags() []string {
	var flags []string

	// loop through comman args
	for _, arg := range os.Args[1:] {
		// if it starts with a '-' then the folowing string is a flag
		if strings.HasPrefix(arg, "-") {
			// loop through every characters of the flag string (except the 1st one as it is a '-')
			for _, flag := range arg[1:] {
				// append the flag to the list of parsed flags
				flags = append(flags, string(flag))
			}
		}
	}

	return flags
}

func GetFlags() Flags {
	var flags Flags

	for _, flag := range ParseFlags() {
		if flag == "R" {
			flags.Recursive = true
		}
		if flag == "l" {
			flags.LongFiles = true
		}
		if flag == "a" {
			flags.HiddenDirectory = true
		}
		if flag == "r" {
			flags.Reverse = true
		}
		if flag == "t" {
			flags.ModifiedTime = true
		}
	}

	return flags
}
