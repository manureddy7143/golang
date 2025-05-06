package main

import (
	"fmt"
	"os"
)

func handleErr(err error) {
	if err != nil {
		panic(err)
	}
}

func checkFileExists(filePath string) bool {
	if _, err := os.Stat(filePath); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

func writeFileData() {
	// TODO: create a new file
	f, err := os.Create("textfile.txt")
	if err != nil {
		handleErr(err)
	}

	// TODO: it is idiomatic in Go to defer the close after you open the file
	defer f.Close()

	// TODO: get the Name of the file
	fmt.Println(f.Name())

	// TODO: write some string data to the file
	f.WriteString("this is test data\n")

	// TODO: write some individual bytes to the file
	//data2 := []byte("this is byte")
	f.Write([]byte{'a', 'b', 'c'})

	// TODO: the Sync function forces the data to be written out
	f.Sync()
}

func appendFileData(fname string, data string) {
	// TODO: Use the lower-level OpenFile function to open the file in append mode
}

func main() {
	// TODO: Simple example: dump some data to a file
	// data := []byte("this is some text\n")
	// ioutil.WriteFile("datafile.txt", data, 0644)

	writeFileData()
	// TODO: More complex example: Granular writing of data

	// TODO: append data to a file
}
