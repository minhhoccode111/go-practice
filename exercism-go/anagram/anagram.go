package anagram

import "strings"

func Detect(subject string, candidates []string) []string {
	result := []string{}
	subjectLower := strings.ToLower(subject)
	subjectMap := toMap(subjectLower)
	for _, candidate := range candidates {
		candidateLower := strings.ToLower(candidate)
		candidateMap := toMap(candidateLower)
		if candidateLower == subjectLower {
			continue
		}
		if compareMap(subjectMap, candidateMap) {
			result = append(result, candidate)
		}
	}
	return result
}

func toMap(s string) map[rune]int {
	runeMap := map[rune]int{}
	for _, v := range s {
		runeMap[v]++
	}
	return runeMap
}

func compareMap(m1, m2 map[rune]int) bool {
	if len(m1) != len(m2) {
		return false
	}
	for k, v := range m1 {
		if v != m2[k] {
			return false
		}
	}
	return true
}
