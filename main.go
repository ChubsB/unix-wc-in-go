package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"unicode"
)

func main() {
	filePath := flag.String("file", "", "Path to the file to read")
	countBytes := flag.Bool("c", false, "Count the number of bytes in the file")
	countLines := flag.Bool("l", false, "Count the number of lines in the file")
	countWords := flag.Bool("w", false, "Count the number of words in the file")
	countCharacters := flag.Bool("m", false, "Count the number of characters in the file")

	flag.Parse()

	if *filePath == "" {
		fmt.Println("Usage: unix-wc-in-go -file=path/to/your/file.txt")
		return
	}

	file, err := os.Open(*filePath)
	if err != nil {
		log.Fatalf("Error opening file: %v", err)
	}
	defer file.Close()

	if *countBytes {
		var byteCount int64

		buf := make([]byte, 1)
		for {
			_, err := file.Read(buf)
			if err != nil {
				break
			}
			byteCount++
		}
		fmt.Printf("%d %s\n", byteCount, file.Name())
	}

	if *countLines {
		var lineCount int
		buf := make([]byte, 1)
		for {
			n, err := file.Read(buf)
			if n == 1 {
				if buf[0] == '\n' {
					lineCount++
				}
			}
			if err != nil {
				break
			}
		}
		fmt.Printf("%d %s\n", lineCount, file.Name())
	}

	if *countWords {
		var wordCount int
		inWord := false
		content, err := io.ReadAll(file)
		if err != nil {
			log.Fatalf("Error reading file: %v", err)
		}

		for _, runeValue := range string(content) {
			if unicode.IsSpace(runeValue) {
				if inWord {
					wordCount++
					inWord = false
				}
			} else {
				inWord = true
			}
		}

		if inWord {
			wordCount++
		}

		fmt.Printf("%d %s\n", wordCount, file.Name())
	}

	if *countCharacters {
		content, err := io.ReadAll(file)
		if err != nil {
			log.Fatalf("Error reading file: %v", err)
		}
		characterCount := len([]rune(string(content)))
		fmt.Printf("%d %s\n", characterCount, file.Name())
	}
}
