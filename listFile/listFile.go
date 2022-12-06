package listfile

import (
	"fmt"
	"os"

	"myls/colors"
	error "myls/error"
	filetype "myls/fileType"
)

func ListFile(path string) {
	files, err := os.ReadDir(path)
	error.Error(err)

	for _, file := range files {
		fmt.Println(filetype.Permission(), filetype.Size(), file.Name())
	}
}

func RecursiveFile(path string) {
	files, err := os.ReadDir(path)
	error.Error(err)

	for _, file := range files {
		fmt.Println(file.Name())
		if file.IsDir() {
			fmt.Println(colors.ColorReset)
			ListFile(file.Name())
		}
		if !file.Type().IsDir() {
			fmt.Println(colors.ColorGreen)
		}
	}
}
