package letter

// FreqMap records the frequency of each rune in a given text.
type FreqMap map[rune]int

// Frequency counts the frequency of each rune in a given text and returns this
// data as a FreqMap.
func Frequency(text string) FreqMap {
	freq := FreqMap{}
	for _, r := range text {
		freq[r]++
	}
	return freq
}

// ConcurrentFrequency counts the frequency of each rune in the given strings,
// by making use of concurrency.
func ConcurrentFrequency(texts []string) FreqMap {
	c := make(chan FreqMap)
	for _, text := range texts {
		go func(s string) { c <- Frequency(s) }(text)
	}
	result := FreqMap{}
	// This loop ensures that you receive a value from the channel for each
	// text. It expects to receive a FreqMap from each goroutine.
	for range texts {
		for key, value := range <-c {
			result[key] += value
		}
	}
	return result
}
