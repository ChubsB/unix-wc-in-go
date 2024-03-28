package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	filePath := flag.String("file", "", "Path to the file to read")
	countBytes := flag.Bool("c", false, "Count the number of bytes in the file")
	countLines := flag.Bool("l", false, "Count the number of lines in the file")
	countWords := flag.Bool("w", false, "Count the number of words in the file")

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
		var word []byte

		buf := make([]byte, 1)
		for {
			n, err := file.Read(buf)
			if n == 1 {
				if buf[0] == ' ' {
					if len(word) > 0 {
						fmt.Println(string(word))
						word = []byte{}
						wordCount++
					}
				} else {
					word = append(word, buf[0])
				}
			}
			if err != nil {
				if len(word) > 0 {
					fmt.Println(string(word))
					wordCount++
				}
				break
			}
		}
		fmt.Printf("Total words: %d in file %s\n", wordCount, file.Name())
	}
}
