// Package proverb should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package proverb

import "fmt"

// Proverb should have a comment documenting it.
func Proverb(in []string) []string {
	result := []string{}
	for i, v := range in {
		if i == len(in)-1 {
			sen := fmt.Sprintf("And all for the want of a %v.", in[0])
			result = append(result, sen)
			continue
		}
		sen := fmt.Sprintf("For want of a %v the %v was lost.", v, in[i+1])
		result = append(result, sen)
	}
	return result
}
