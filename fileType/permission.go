package filetype

import (
	"os"

	"myls/error"
)

func Permission() string {
	fileInfo, err := os.Stat(path)
	error.Error(err)

	// Get the file's permission bits
	perm := fileInfo.Mode().Perm()

	// Convert the permission bits to a string representation
	return perm.String()
}
