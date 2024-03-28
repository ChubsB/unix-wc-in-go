package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	// Define command-line flags
	filePath := flag.String("file", "", "Path to the file to read")

	// Parse command-line flags
	flag.Parse()

	// Check if the file flag is provided
	if *filePath == "" {
		fmt.Println("Usage: unix-wc-in-go -file=path/to/your/file.txt")
		return
	}

	// Read the file content
	content, err := os.ReadFile(*filePath)
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return
	}

	// Print the file content
	fmt.Println(string(content))
}
