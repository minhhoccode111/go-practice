package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
	"time"
)

// dummy handler
type fooHandler struct{}

// have to have a ServeHTTP() method
func (h fooHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("dummy foo"))
}

// ListenAndServe starts an HTTP server with a given address and handler. The handler is usually nil, which means to use DefaultServeMux. Handle and HandleFunc add handlers to DefaultServeMux:
func main() {
	http.Handle("/foo", fooHandler{})
	http.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	})

	log.Fatal(http.ListenAndServe(":8080", nil))

	// More control over the server's behavior is available by creating a custom Server:
	s := &http.Server{
		Addr:           ":8080",
		Handler:        nil,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	log.Fatal(s.ListenAndServe())
}
