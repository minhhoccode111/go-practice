package main

import (
	"fmt"
	"math/bits"
	"strconv"
)

func countBits(n uint32) int32 {
	return int32(bits.OnesCount32(n))
}

func main() {
	var num uint32
	for {
		var numStr string
		fmt.Printf("num: ")
		fmt.Scanln(&numStr)
		n, err := strconv.ParseUint(numStr, 10, 32)
		if err == nil {
			num = uint32(n)
			break
		}
	}

	fmt.Println(countBits(num))
}
