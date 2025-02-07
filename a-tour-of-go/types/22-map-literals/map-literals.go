package main

import "fmt"

/*
Map literals
Map literals are like struct literals, but the keys are required.
*/

type V struct {
	Lat, Long float64
}

var m = map[string]V{
	"Google": V{
		37.42202, -122.08408,
	},
	"Microsoft": V{
		137.42202, -122.08408,
	},
}

func main() {
	fmt.Println(m) // map[Google:{37.42202 -122.08408} Microsoft:{137.42202 -122.08408}]
}
