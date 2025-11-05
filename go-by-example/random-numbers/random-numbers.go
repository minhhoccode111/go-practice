package main

import (
	"fmt"
	"math/rand/v2"
)

func main() {
	fmt.Print(rand.IntN(100), ",")
	fmt.Print(rand.IntN(100))
	fmt.Println()

	fmt.Println(rand.Float64())

	fmt.Print((rand.Float64()*5)+5, ",")
	fmt.Print((rand.Float64() * 5) + 5)
	fmt.Println()

	s2 := rand.NewPCG(42, 1024)
	r2 := rand.New(s2)
	fmt.Print(r2.IntN(100), ",")
	fmt.Print(r2.IntN(100))
	fmt.Println()

	s3 := rand.NewPCG(42, 1024)
	r3 := rand.New(s3)
	fmt.Print(r3.IntN(100), ",")
	fmt.Print(r3.IntN(100))
	fmt.Println()
	/*
	   $ go run random-numbers.go
	   96,59
	   0.7275094999762297
	   9.208115731350528,5.095073493424972
	   94,49
	   94,49
	*/
}
