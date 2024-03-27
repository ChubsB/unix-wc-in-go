package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	// Specify the path to your test.txt file
	filePath := "test.txt"

	// Read the file content
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}

	// Print the file content as a string
	fmt.Println(string(content))
}
