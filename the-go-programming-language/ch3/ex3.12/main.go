package main

import "slices"

func main() {
}

func IsAnagram(s1, s2 string) bool {
	b1 := []byte(s1)
	slices.Sort(b1)
	b2 := []byte(s2)
	slices.Sort(b2)
	return string(b1) == string(b2)
}
