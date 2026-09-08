package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

// pc[i] is the population count of i
var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

// PopCountLoop returns the population count (number of set bits) of x
func PopCountLoop(x uint64) int {
	var count int
	for i := range 8 {
		count += int(pc[byte(x>>(i*8))])
	}
	return count
}

// PopCount returns the population count (number of set bits) of x
func PopCount(x uint64) int {
	return int(
		pc[byte(x>>(0*8))] +
			pc[byte(x>>(1*8))] +
			pc[byte(x>>(2*8))] +
			pc[byte(x>>(3*8))] +
			pc[byte(x>>(4*8))] +
			pc[byte(x>>(5*8))] +
			pc[byte(x>>(6*8))] +
			pc[byte(x>>(7*8))])
}

func main() {
	start := time.Now()
	for _, arg := range os.Args[1:] {
		v, err := strconv.ParseUint(arg, 10, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
			os.Exit(1)
		}
		fmt.Println(PopCountLoop(v))
		fmt.Println(PopCount(v))
	}
	fmt.Println("elapsed", time.Since(start))
	// go run main.go 10000000000000000000
}
