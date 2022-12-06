package listfile

import (
	"fmt"
	"log"
	"os"
)

func ListFile(path string) {
	files, err := os.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		fmt.Println(file.Name())
	}
}

func RecursiveFile(path string, directoryExist bool) {
	files, err := os.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		if directoryExist { // if a subdirectory exist, it will list the file inside it
			fmt.Println(file.Name())
		}
	}
}
