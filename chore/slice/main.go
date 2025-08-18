package main

import "fmt"

func main() {
	arr := [5]int{1, 2, 3, 4, 5}
	fmt.Println(arr)

	s0 := arr[:]
	fmt.Println(s0, len(s0), cap(s0))

	s1 := arr[1:]
	fmt.Println(s1, len(s1), cap(s1))

	s2 := arr[1:3]
	fmt.Println(s2, len(s2), cap(s2))

	s3 := s2[:cap(s2)]
	fmt.Println(s3, len(s3), cap(s3))

	// create slice with make
	s4 := make([]int, 50, 100)
	fmt.Println(s4, len(s4), cap(s4))
	// is the same as this, create a slice with 100 int, and the first 50
	s5 := new([100]int)[0:50]
	fmt.Println(s5, len(s5), cap(s5))
}
