package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func sanitizeFilename(title string) string {
	title = strings.TrimPrefix(title, "title: ")
	title = strings.ToLower(title)
	title = strings.ReplaceAll(title, " ", "-")
	title = strings.ReplaceAll(title, "'", "-")
	return title + ".md"
}

func main() {
	if len(os.Args) != 4 {
		panic("Usage: split-file-to-chunks <input-file(str)> <chunk-size(int)> <title-position(int)>")
	}
	inputFile := os.Args[1]
	chunkSize, err := strconv.Atoi(os.Args[2])
	if err != nil {
		panic("Invalid chunk size")
	}
	titlePosition, err := strconv.Atoi(os.Args[3])
	if err != nil {
		panic("Invalid title position")
	}

	file, err := os.Open(inputFile)
	if err != nil {
		panic("Can't read the file")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var chunk []string
	chunkCount := 0

	for scanner.Scan() {
		chunk = append(chunk, scanner.Text())
		if len(chunk) == chunkSize {
			if len(chunk) > titlePosition {
				filename := sanitizeFilename(chunk[titlePosition])
				err := os.WriteFile(filename, []byte(strings.Join(chunk, "\n")+"\n"), 0644)
				if err != nil {
					panic(fmt.Sprintf("Failed to write file %s", filename))
				}
				fmt.Printf("Created: %s\n", filename)
			}
			chunk = nil
			chunkCount++
		}
	}

	if err := scanner.Err(); err != nil {
		panic("Error reading file")
	}
}
