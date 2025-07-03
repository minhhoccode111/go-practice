package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

type Response struct {
	Message string `json:"message"`
}

func main() {
	port := ":8000"
	fmt.Println("start listen on http://localhost", port)
	s := &http.Server{
		Addr:         ":8000",
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
		Handler:      nil,
	}
	http.HandleFunc("/", handleIndex)
	http.HandleFunc("/hello", handleHello)

	log.Fatal(s.ListenAndServe())
}

func handleIndex(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Hello, World")
}

func handleHello(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	name := r.URL.Query().Get("name")
	var message Response
	if name == "" {
		message = Response{Message: "Hello, World!"}
	} else {
		message = Response{Message: fmt.Sprintf("Hello, %s!", name)}
	}

	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(message)

	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}
}
