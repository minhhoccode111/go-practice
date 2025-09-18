package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

/*
Exercise 1.10: Find a web site that produces a large amount of data. Investigate
caching by running fetchall twice in succession to see whether the reported time
changes much. Do you get the same content each time? Modify fetchall to print
its output to a file so it can be examined.
*/

func main() {
	start := time.Now()
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		go fetch(url, ch)
	}
	for range os.Args[1:] {
		fmt.Println(<-ch)
	}
	fmt.Printf("%.7fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}
	nbytes, err := io.Copy(os.Stdout, resp.Body)
	resp.Body.Close()
	if err != nil {
		ch <- fmt.Sprintf("While reading %s: %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.7fs %7d %s", secs, nbytes, url)
}
