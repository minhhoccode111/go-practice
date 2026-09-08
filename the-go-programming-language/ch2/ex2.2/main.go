package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"ch2/ex2.2/weightconv"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		line := scanner.Text()
		args = strings.Fields(line)
	}
	var weights []float64
	for _, arg := range args {
		w, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "parse: %v\n", err)
			os.Exit(1)
		}
		weights = append(weights, w)
	}
	conv(weights)
	/*
	   mhc in ~/projects/go-practice/the-go-programming-language/ch2/ex2.2 on develop [!?]
	   $ echo "12 11" | ./main
	   12kg = 26.4552lb, 12lb = 5.443164292842239kg
	   11kg = 24.250600000000002lb, 11lb = 4.9895672684387185kg
	   mhc in ~/projects/go-practice/the-go-programming-language/ch2/ex2.2 on develop [!?]
	   $ ./main 12 11
	   12kg = 26.4552lb, 12lb = 5.443164292842239kg
	   11kg = 24.250600000000002lb, 11lb = 4.9895672684387185kg
	*/
}

func conv(n []float64) {
	for _, v := range n {
		kg := weightconv.Kilogram(v)
		lb := weightconv.Pound(v)
		fmt.Printf("%s = %s, %s = %s\n", kg, weightconv.KgToLb(kg), lb, weightconv.LbToKg(lb))
	}
}
