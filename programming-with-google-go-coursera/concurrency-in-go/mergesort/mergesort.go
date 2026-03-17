package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
	"sync"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func getUserInput() []int {
	var result []int
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	line := scanner.Text()
	parts := strings.FieldsSeq(line)
	for v := range parts {
		n, err := strconv.Atoi(v)
		check(err)
		result = append(result, n)
	}
	return result
}

func sort(nums []int, wg *sync.WaitGroup) {
	fmt.Println(nums)
	slices.Sort(nums)
	wg.Done()
}

func isNotEmpty(s []int) bool { return len(s) != 0 }

func main() {
	fmt.Print("Input array of integers: ")

	nums := getUserInput()

	l := 4
	parts := make([][]int, l)
	for i := range parts {
		parts[i] = make([]int, 0)
	}

	for i, v := range nums {
		parts[i%l] = append(parts[i%l], v)
	}

	var wg sync.WaitGroup

	for _, v := range parts {
		wg.Add(1)
		go sort(v, &wg)
	}

	wg.Wait()

	var final []int
	for slices.ContainsFunc(parts, isNotEmpty) {
		minN := math.Inf(1)
		minIndex := -1
		for i := range parts {
			if len(parts[i]) == 0 {
				continue
			}

			curr := float64(parts[i][0])

			if curr < minN {
				minN = curr
				minIndex = i
			}
		}
		parts[minIndex] = parts[minIndex][1:]
		final = append(final, int(minN))
	}

	fmt.Println(final)
}
