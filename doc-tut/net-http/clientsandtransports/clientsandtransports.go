package main

import (
	"errors"
	"fmt"
	"net/http"
	"time"
)

/*
Clients and Transports are safe for concurrent use by multiple goroutines and
for efficiency should only be created once and re-used.
*/

// type CheckRedirect func(req *http.Request, via []*http.Request) error

func redirectPolicyFunc(req *http.Request, via []*http.Request) error {
	if len(via) >= 10 {
		return errors.New("stopped after 10 redirects")
	}
	// copy headers from old request to new request
	req.Header.Add("User-Agent", "my-client")
	return nil
}

func main() {
	clients()
	transports()
}

// for controler over HTTP client headers, redirect policy, and other settings,
// create a Client
func clients() {
	client := &http.Client{
		CheckRedirect: redirectPolicyFunc,
	}

	resp, err := client.Get("http://example.com")
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}
	_ = resp
	req, err := http.NewRequest("GET", "http://example.com", nil)
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}
	req.Header.Add("If-None-Match", `W/"wyzzy`)
	resp, err = client.Do(req)
}

// for control over procies, TLD configuration, keep-alive, compression, and
// other settings, create a Transport
func transports() {
	tr := &http.Transport{
		MaxIdleConns:       10,
		IdleConnTimeout:    30 * time.Second,
		DisableCompression: true,
	}
	client := &http.Client{
		Transport: tr,
	}
	resp, err := client.Get("https://example.com")
	_ = err
	_ = resp
}
