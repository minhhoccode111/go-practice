package letter

import "sync"

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
	freq := FreqMap{}
	var wg sync.WaitGroup
	c := make(chan FreqMap)
	go func() {
		wg.Wait()
		close(c)
	}()
	for _, text := range texts {
		wg.Add(1)
		go func(s string) {
			defer wg.Done()
			result := Frequency(s)
			c <- result
		}(text)
	}
	for v := range c {
		for key, val := range v {
			freq[key] += val
		}
	}
	return freq
}
