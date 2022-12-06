package subdirectory

import (
	"os"
)

// func IsSubDirectory() bool {
// 	for _, arg := range os.Args[1:] {
// 		if strings.HasPrefix(arg, "..") {
// 			return true
// 		}
// 	}
// 	return false
// }

func DirectoryExist() bool {
	currentDirectory := "."

	if _, err := os.Stat(currentDirectory); os.IsNotExist(err) {
		return true
	}

	return false
}
