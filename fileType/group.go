package filetype

import (
	"os"
)

func Group() int {
	// Get the group ID
	gid := os.Getgid()

	return gid
}
