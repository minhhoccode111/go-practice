package main

import (
	"fmt"
	"net/http"
	"time"
)

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Println("hello handler started")
	defer fmt.Println("hello handler ended")

	ctx := r.Context()

	select {
	case <-time.After(5 * time.Second):
		fmt.Println("work done")
	case <-ctx.Done():
		err := ctx.Err()
		fmt.Println("server:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func main() {
	http.HandleFunc("/hello", hello)
	http.ListenAndServe("localhost:8000", nil)
}
