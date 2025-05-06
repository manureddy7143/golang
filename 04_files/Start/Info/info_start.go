package main

import (
	"fmt"
	"os"
)

func checkFileExists(filePath string) bool {
	if _, err := os.Stat(filePath); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}

	return true
}

func main() {
	// Use the Stat function to get file stats
	stasts, err := os.Stat("sampletext.txt")
	if err != nil {
		panic(err)
	}

	// Check if a file exists
	exits := checkFileExists("sampletext.txt")
	fmt.Println("file exists checks", exits)

	// Get the file's modification time
	fmt.Println(stasts.ModTime())
	fmt.Println(stasts.Mode())
	fmode := stasts.Mode()

	if fmode.IsRegular() {
		fmt.Println("regular file")
	}

	// Get the file size
	fmt.Println(stasts.Size())

	// Is this a directory?
	fmt.Println(stasts.IsDir())

}
