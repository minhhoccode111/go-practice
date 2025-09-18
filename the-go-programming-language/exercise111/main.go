package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

/*
Exercise 1.11: Try fetchall with longer argument lists, such as samples from the
top million web sites available at alexa.com. How does the program behave if a
web site just doesn’t respond? (Section 8.9 describes mechanisms for coping in
such cases.)
*/

func main() {
	start := time.Now()
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		go fetch(url, ch) // start a goroutine
	}
	for range os.Args[1:] {
		fmt.Println(<-ch) // receive from channel ch
	}
	fmt.Printf("%.7fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	client := http.Client{Timeout: 1 * time.Second}
	resp, err := client.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err) //
		return
	}
	nbytes, err := io.Copy(io.Discard, resp.Body)
	resp.Body.Close() // don't leak resources
	if err != nil {
		ch <- fmt.Sprintf("While reading %s: %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.7fs %7d %s", secs, nbytes, url)
}
