package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

/*
Exercise 1.9: Modify fetch to also print the HTTP status code, found in
resp.Status.
*/

const prefix = "http://"

func main() {
	for _, url := range os.Args[1:] {
		if !strings.HasPrefix(url, prefix) {
			url = prefix + url
		}
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		b, err := io.ReadAll(resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}
		fmt.Printf("%s\n", b)
		fmt.Println("http status code:", resp.Status)
	}
}
