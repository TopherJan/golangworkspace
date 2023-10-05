package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	// Check if the filename is provided
	//
	if len(os.Args) != 2 {
		fmt.Println("Error: Filename not provided.")
		fmt.Println("Usage: go run main.go <filename>")
		os.Exit(1)
	}

	// Retrieve the filename from the argument
	filename := os.Args[1]

	// Open the file
	//
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Failed to open file with error:", err)
		os.Exit(1)
	}

	// Print out the text inside the file
	//
	io.Copy(os.Stdout, file)
}
