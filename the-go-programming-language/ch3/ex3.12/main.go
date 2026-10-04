package main

func main() {
}

func IsAnagram(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	var freq [256]int
	for i := range len(s1) {
		freq[s1[i]]++
		freq[s2[i]]--
	}
	for _, v := range freq {
		if v != 0 {
			return false
		}
	}
	return true
}
