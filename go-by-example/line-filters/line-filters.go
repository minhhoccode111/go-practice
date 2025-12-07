package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

/*
Problem: read from stdin and print upper case to stdout, errors to stderr
*/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		str := strings.ToUpper(scanner.Text())
		fmt.Println(str)
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, err.Error())
		os.Exit(1)
	}
	/*
	   $ echo 'Hello, World!' | ./line-filters
	   HELLO, WORLD!
	*/
}
