package dna

import "errors"

// Histogram is a mapping from nucleotide to its count in given DNA.
// Choose a suitable data type.
type Histogram map[rune]int

// DNA is a list of nucleotides. Choose a suitable data type.
type DNA []rune

// Counts generates a histogram of valid nucleotides in the given DNA.
// Returns an error if d contains an invalid nucleotide.
// /
// Counts is a method on the DNA type. A method is a function with a special receiver argument.
// The receiver appears in its own argument list between the func keyword and the method name.
// Here, the Counts method has a receiver of type DNA named d.
func (d DNA) Counts() (Histogram, error) {
	h := map[rune]int{
		'A': 0,
		'G': 0,
		'T': 0,
		'C': 0,
	}
	for _, v := range d {
		if !isValid(v) {
			return nil, errors.New("invalid nucleotide")
		}
		nuc, exists := h[v]
		if !exists {
			h[v] = 1
		} else {
			h[v] = nuc + 1
		}
	}
	return h, nil
}

func isValid(c rune) bool {
	return c == 'G' || c == 'A' || c == 'T' || c == 'C'
}
