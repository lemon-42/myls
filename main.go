// List the current file in the directory :
// for each files in the current directory :
// print each files
// if error, print a message

// before doing any flags, we gonna do -R first
// -R list subdirectories encountered
// for each files in the current directory :
// list all files
// then recall the -R function to list all the subdirectories

package main

import (
	"os"

	getflags "myls/getFlags"
	listFile "myls/listFile"
)

var userInput = os.Args[1]

func main() {
	flags := getflags.GetFlags()

	// "-R" flags
	if flags.Recursive {
		listFile.RecursiveFile(userInput)
	}

	// "-l" flags
	if flags.LongFiles {
		listFile.ListFile(userInput)
	}

	// "-a" flags
	if flags.HiddenDirectory {
		listFile.ListFile(userInput)
	}

	// "-r" flags
	if flags.Reverse {
		listFile.ListFile(userInput)
	}

	// "-t" flags
	if flags.ModifiedTime {
		listFile.ListFile(userInput)
	}
}
