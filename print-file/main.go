package main

import (
	"fmt"
	"io"
	"os"
)

// Create a program that reads the contents of
// a text file then prints its contents to the
// terminal
func main() {
	// os.Args is a slice of type string with
	// command lines used in the program
	fileName := os.Args[1]
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	io.Copy(os.Stdout, file)
}
