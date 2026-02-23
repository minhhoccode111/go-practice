package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	nums := make([]int, 3)
	scanner := bufio.NewScanner(os.Stdin)
	count := 0
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if strings.EqualFold(line, "x") {
			break
		}
		n, err := strconv.Atoi(line)
		if err != nil {
			panic(err)
		}
		if count < 3 {
			nums[count] = n
		} else {
			nums = append(nums, n)
		}
		count++
		activeSlice := nums[:count]
		slices.Sort(activeSlice)
		fmt.Println(activeSlice)
	}
}
