package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	fmt.Println("loop start====================================================")
	loop()
	fmt.Println("loop end======================================================")

	fmt.Println("join start====================================================")
	join()
	fmt.Println("join end======================================================")
	/*
	   $ ./main Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.
	   loop start====================================================
	   ./mainLorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.
	   runtime: 51.46µs
	   loop end======================================================
	   join start====================================================
	   ./main Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.
	   runtime: 3.744µs
	   join end======================================================
	*/
}

func loop() {
	var s, sep string
	t := time.Now()
	for _, v := range os.Args {
		s += v + sep
		sep = " "
	}
	fmt.Println(s)
	d := time.Since(t)
	fmt.Println("runtime:", d)
}

func join() {
	t := time.Now()
	fmt.Println(strings.Join(os.Args, " "))
	d := time.Since(t)
	fmt.Println("runtime:", d)
}
