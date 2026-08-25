package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
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
		Handler:      mux,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  10 * time.Second,
	}

	go func() {
		log.Println("starting server at :8080")
		if err := svr.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal(err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	defer signal.Stop(quit)

	// this block
	<-quit
	log.Println("signal received, wait 5 seconds before shutdown, Ctrl-C again to force")

	// this isn't block
	go func() {
		<-quit
		log.Println("force exit!")
		os.Exit(1)
	}()

	// graceful shutdown for 5 seconds
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := svr.Shutdown(ctx); err != nil {
		log.Println("error shutting down")
		log.Fatal(err)
	}
}
