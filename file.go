package main

import (
	"fmt"
	"io"
	"os"
)

var path = "/Users/mymac/Documents/Programming/go/Bootcamp/novgolang/test.txt"

// create error helper
func isError(err error) bool {
	if err != nil {
		fmt.Println(err.Error())
	}
	return (err != nil)
}

// to create new file
func createFile() {
	// to detect file is exist or not
	var _, err = os.Stat(path)

	// create new file if not exist
	if os.IsNotExist(err) {
		var file, err = os.Create(path)
		if isError(err) {
			return
		}
		defer file.Close()
	}
	fmt.Println("==> success create file")
}

// to edit file that was created before
func writeFile() {
	// open file with access level READ & WRITE
	var file, err = os.OpenFile(path, os.O_RDWR, 0644)
	if isError(err) {
		return
	}
	defer file.Close()

	// write data into existing file
	_, err = file.WriteString("Hello\n")
	if isError(err) {
		return
	}
	_, err = file.WriteString("Let's to learn Golang!\n")
	if isError(err) {
		return
	}

	// save changes
	err = file.Sync()
	if isError(err) {
		return
	}

	fmt.Println("==> success input into existing file")
}

// to read existing file
func readFile() {
	// open file
	var file, err = os.OpenFile(path, os.O_RDONLY, 0644)
	if isError(err) {
		return
	}
	defer file.Close()

	// read file
	var text = make([]byte, 1024)

	for {
		n, err := file.Read(text)
		if err != io.EOF {
			if isError(err) {
				break
			}
		}
		if n == 0 {
			break
		}

		if isError(err) {
			return
		}

		fmt.Println("==> success read existing file")
		fmt.Println(string(text))
	}
}

// to delete file
func deleteFile() {
	var err = os.Remove(path)
	if isError(err) {
		return
	}
	fmt.Println("==> success deleting file")
}

func main() {
	createFile()
	writeFile()
	readFile()
	deleteFile()
}
