package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

func main() {
	start := time.Now()
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		go fetch(url, ch)
	}
	f, err := os.Create("output")
	if err != nil {
		fmt.Printf("open file: %v\n", err)
		os.Exit(1)
	}
	defer f.Close()
	for range os.Args[1:] {
		_, err := f.Write([]byte(<-ch + "\n"))
		if err != nil {
			fmt.Printf("write file: %v\n", err)
		}
	}
	_, err = f.Write([]byte(fmt.Sprintf("%.2fs elapsed\n", time.Since(start).Seconds())))
	if err != nil {
		fmt.Printf("write file: %v\n", err)
	}
}

func fetch(url string, ch chan string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}
	n, err := io.Copy(io.Discard, resp.Body)
	resp.Body.Close()
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	ch <- fmt.Sprintf("%.2fs %7d %s", time.Since(start).Seconds(), n, url)
}
