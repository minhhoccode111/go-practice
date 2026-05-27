package main

import (
	"flag"
	"fmt"
)

func main() {
	wordPtr := flag.String("word", "foo", "a string")
	numberPtr := flag.Int("number", 1, "a number")
	forkPtr := flag.Bool("fork", false, "a boolean")
	var svar string
	flag.StringVar(&svar, "svar", "bar", "a string")
	flag.Parse()

	fmt.Println(*wordPtr)
	fmt.Println(*numberPtr)
	fmt.Println(*forkPtr)
	fmt.Println(svar)
	fmt.Println("tail: ", flag.Args())
}
