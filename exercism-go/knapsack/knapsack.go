package knapsack

type Item struct {
	Weight, Value int
}

// Knapsack takes in a maximum carrying capacity and a collection of items
// and returns the maximum value that can be carried by the knapsack
// given that the knapsack can only carry a maximum weight given by maximumWeight
func Knapsack(maxWeight int, items []Item) int {
	// create a slice of int that over weight by one item
	tableOfWeights := make([]int, maxWeight+1)
	// loop through slice of items
	for _, item := range items {
		// for each item, we loop from max possible weight to current weight of item (backward)
		for weight := maxWeight; weight >= item.Weight; weight-- {
			// then update value of current weight on the table
			// by comparing between its current value and the maximum value of
			// the available weight before this (weight - item.weight) plus
			// current value and choose the larger
			tableOfWeights[weight] = max(tableOfWeights[weight], tableOfWeights[weight-item.Weight]+item.Value)
		}
	}
	// then return the max weight possition of the table
	return tableOfWeights[maxWeight]
}

// compare 2 numbers and return greater one
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
