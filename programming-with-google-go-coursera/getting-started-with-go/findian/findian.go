package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

func main() {
	re := regexp.MustCompile(`(?i)^i.*a.*n$`)
	fmt.Print("Your string: ")
	scanner := bufio.NewScanner(os.Stdin)
	var input string
	if scanner.Scan() {
		input = scanner.Text()
	}

	if re.MatchString(input) {
		fmt.Println("Found!")
	} else {
		fmt.Println("Not Found!")
	}
}
