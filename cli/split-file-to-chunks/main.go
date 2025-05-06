package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func panicUsage() {
	panic("usage: ./split-file-to-chunks file-name<string> chunk-size<int> title-index<int>")
}

func sanitizeFileName(s string) string {
	// remove prefix "title: "
	reTitle := regexp.MustCompile(`^title:\s`)
	s = reTitle.ReplaceAllString(s, "")
	// format file name
	s = strings.ToLower(s)
	s = strings.ReplaceAll(s, " ", "-")
	re := regexp.MustCompile(`[^a-z0-9\-]`)
	s = re.ReplaceAllString(s, "")
	return s + ".md"
}

func main() {
	// extract command line arguments
	args := os.Args[1:]
	if len(args) < 3 {
		panicUsage()
	}

	inputFileName := args[0]

	eachChunkSize, err := strconv.Atoi(args[1])
	if err != nil {
		panicUsage()
	}

	titleIndex, err := strconv.Atoi(args[2])
	if err != nil {
		panicUsage()
	}

	openedInputFile, err := os.Open(inputFileName)
	if err != nil {
		panic(err)
	}
	defer openedInputFile.Close()

	linesOfChunk := make([]string, eachChunkSize)
	outputName := ""
	count := 0

	scanner := bufio.NewScanner(openedInputFile)
	for scanner.Scan() {
		currentLine := scanner.Text()
		linesOfChunk[count] = currentLine

		if count == titleIndex {
			outputName = sanitizeFileName(currentLine)
		}

		count++

		if count == eachChunkSize {
			count = 0

			outputFile, err := os.Create(outputName)
			if err != nil {
				panic(err)
			}
			defer outputFile.Close()

			fmt.Println("Creating file: ", outputName)
			outputFile.WriteString(strings.Join(linesOfChunk, "\n"))
		}
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}
}
