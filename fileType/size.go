package filetype

import (
	"os"

	"myls/error"
)

const (
	path = "."
)

func Size() int64 {
	fileInfo, err := os.Stat(path)
	error.Error(err)

	return fileInfo.Size()
}
