package strain

// Implement the "Keep" and "Discard" function in this file.

// You will need typed parameters (aka "Generics") to solve this exercise.
// They are not part of the Exercism syllabus yet but you can learn about
// them here: https://go.dev/tour/generics/1

func Keep[T any](list []T, filterFunc func(T) bool) []T {
	returnList := []T{}
	for _, v := range list {
		if filterFunc(v) {
			returnList = append(returnList, v)
		}
	}
	return returnList
}

func Discard[T any](list []T, filterFunc func(T) bool) []T {
	returnList := []T{}
	for _, v := range list {
		if !filterFunc(v) {
			returnList = append(returnList, v)
		}
	}
	return returnList
}
