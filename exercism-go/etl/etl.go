package etl

import "strings"

func Transform(in map[int][]string) map[string]int {
	result := map[string]int{}
	for km, vm := range in {
		for _, va := range vm {
			va = strings.ToLower(va)
			result[va] = km
		}
	}
	return result
}
