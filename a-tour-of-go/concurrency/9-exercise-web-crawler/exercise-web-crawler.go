package main

/*
Exercise: Web Crawler
In this exercise you'll use Go's concurrency features to parallelize a web
crawler.

Modify the Crawl function to fetch URLs in parallel without fetching the same
URL twice.

Hint: you can keep a cache of the URLs that have been fetched on a map, but
maps alone are not safe for concurrent use!
*/

import (
	"fmt"
	"sync"
)

type Fetcher interface {
	Fetch(url string) (body string, urls []string, err error)
}

type SafeFetch struct {
	mu          sync.Mutex
	wg          sync.WaitGroup
	fetchedUrls map[string]bool
}

func (sf *SafeFetch) Crawl(url string, depth int, fetcher Fetcher) {
	// after this function finishes, tell the waitgroup that 1 goroutine has
	// finished
	defer sf.wg.Done()

	if depth <= 0 {
		return
	}

	// lock the mutex
	sf.mu.Lock()
	// to check is url is already fetched
	if _, ok := sf.fetchedUrls[url]; ok {
		// if yes, unlock and return
		sf.mu.Unlock()
		return
	}
	// else mark it as fetched
	sf.fetchedUrls[url] = true
	// and unlock the mutex, for other goroutines to use
	sf.mu.Unlock()

	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("found: %s %q\n", url, body)

	// for every url in current page
	for _, u := range urls {
		// tell the waitgroup that we spawn 1 goroutine
		sf.wg.Add(1)
		// launch the child goroutines
		// to crawl that url, with depth-1
		go sf.Crawl(u, depth-1, fetcher)
	}
}

func main() {
	sf := SafeFetch{fetchedUrls: make(map[string]bool)}
	// tell the waitgroup that we have 1 goroutine
	sf.wg.Add(1)
	// 4 levels of depth is reasonable, because we will pretty much go cycle at that point
	go sf.Crawl("https://golang.org/", 4, fetcher)
	// tell the waitgroup to wait for all goroutines to finish
	sf.wg.Wait()
}

type fakeFetcher map[string]*fakeResult

type fakeResult struct {
	body string
	urls []string
}

func (f fakeFetcher) Fetch(url string) (string, []string, error) {
	if res, ok := f[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found: %s", url)
}

var fetcher = fakeFetcher{
	"https://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"https://golang.org/pkg/",
			"https://golang.org/cmd/",
		},
	},
	"https://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"https://golang.org/",
			"https://golang.org/cmd/",
			"https://golang.org/pkg/fmt/",
			"https://golang.org/pkg/os/",
		},
	},
	"https://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
	"https://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
}
