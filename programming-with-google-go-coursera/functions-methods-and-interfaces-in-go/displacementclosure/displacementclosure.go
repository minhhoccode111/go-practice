package main

import (
	"fmt"
)

func GenDisplaceFn(a, vo, so float64) func(float64) float64 {
	return func(t float64) float64 {
		return 0.5*a*t*t + vo*t + so
	}
}

func main() {
	var a, vo, so float64
	var t float64

	fmt.Print("Enter acceleration, initial velocity, initial displacement: ")
	fmt.Scan(&a, &vo, &so)

	fn := GenDisplaceFn(a, vo, so)

	fmt.Print("Enter time: ")
	fmt.Scan(&t)

	fmt.Println(fn(t))
}
