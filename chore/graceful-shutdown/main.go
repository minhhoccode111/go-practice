package main

import (
	"context"
	"log"
	"net/http"
	"time"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("request to /")
		sleep := r.URL.Query().Get("sleep")
		d, err := time.ParseDuration(sleep)
		if err != nil {
			d = 0 * time.Second
		}
		time.Sleep(d)
		w.Write([]byte("done\n"))
	})
	svr := http.Server{
		Addr:         ":8080",
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  10 * time.Second,
	}
	log.Println("starting server at :8080")
	go func() {
		if err := svr.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal(err)
		}
	}()
	ctx := context.Background()
	svr.Shutdown(ctx)
}
