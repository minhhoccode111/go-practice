package grep

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func Search(pattern string, flags, files []string) []string {
	result := []string{}
	var (
		lineNumber      bool
		onlyFileName    bool
		caseInsensitive bool
		invertProgram   bool
		matchLine       bool
	)
	for _, v := range flags {
		switch v {
		case "-n":
			lineNumber = true
		case "-l":
			onlyFileName = true
		case "-i":
			caseInsensitive = true
			pattern = strings.ToLower(pattern)
		case "-v":
			invertProgram = true
		case "-x":
			matchLine = true
		default:
			panic("invalid flag")
		}
	}
	for _, fileName := range files {
		file, err := os.Open(fileName)
		if err != nil {
			return nil
		}
		defer file.Close()
		scanner := bufio.NewScanner(file)
		count := 1
		for scanner.Scan() {
			if onlyFileName && len(result) > 0 && result[len(result)-1] == fileName {
				continue
			}
			line := scanner.Text()
			currLineFormatted := line
			if caseInsensitive {
				line = strings.ToLower(line)
			}
			if lineNumber {
				currLineFormatted = strconv.Itoa(count) + ":" + currLineFormatted
			}
			count++
			if len(files) > 1 {
				currLineFormatted = fileName + ":" + currLineFormatted
			}
			if onlyFileName {
				currLineFormatted = fileName
			}
			matched := false
			if line == pattern && matchLine {
				matched = true
			}
			if strings.Contains(line, pattern) && !matchLine {
				matched = true
			}
			if invertProgram && !matched {
				result = append(result, currLineFormatted)
			}
			if !invertProgram && matched {
				result = append(result, currLineFormatted)
			}
		}
		if err := scanner.Err(); err != nil {
			return nil
		}
	}
	return result
}
