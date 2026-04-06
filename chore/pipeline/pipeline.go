package main

import "fmt"

func generate(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		for _, v := range nums {
			out <- v
		}
		close(out)
	}()
	return out
}

func square(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for v := range in {
			out <- v * v
		}
		close(out)
	}()
	return out
}

func main() {
	belt1 := generate(1, 2, 3, 4, 5, 6)
	belt2 := square(belt1)

	for v := range belt2 {
		fmt.Println(v)
	}
}
