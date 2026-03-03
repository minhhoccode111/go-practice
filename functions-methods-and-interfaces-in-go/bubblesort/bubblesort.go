package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	fmt.Print("Enter at max 10 numbers: ")
	scanner := bufio.NewScanner(os.Stdin)
	var nums []int
	if scanner.Scan() {
		line := scanner.Text()
		nums = StringToNums(line)
	}
	BubbleSort(nums)
	fmt.Println(nums)
}

func StringToNums(line string) []int {
	var result []int
	for v := range strings.FieldsSeq(line) {
		n, err := strconv.Atoi(v)
		check(err)
		result = append(result, n)
	}
	return result
}

func BubbleSort(nums []int) {
	l := len(nums)
	// the "-1" is because the last element is not compared
	for i := 0; i < l-1; i++ {
		swapped := false
		for j := 0; j < l-1-i; j++ {
			if nums[j] > nums[j+1] {
				Swap(nums, j)
				swapped = true
			}
		}
		if !swapped {
			break
		}
	}
}

func Swap(nums []int, i int) {
	nums[i], nums[i+1] = nums[i+1], nums[i]
}
