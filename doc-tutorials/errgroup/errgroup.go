package main

import (
	"fmt"
	"net/http"
	"time"

	"golang.org/x/sync/errgroup"
)

func main() {
	start := time.Now()
	g := new(errgroup.Group)
	var users = []string{
		"http://www.golang.org/",
		"http://www.google.com/",
		"http://www.youtube.com/",
		"http://www.somestupidname.com/",
	}

	for _, url := range users {
		// launch a goroutine to fetch the url
		g.Go(func() error {
			resp, err := http.Get(url)
			if err != nil {
				return err
			}
			defer resp.Body.Close()
			fmt.Println(url, time.Since(start))
			return nil
		})
	}
	// wait for all fetches to complete
	if err := g.Wait(); err != nil {
		return
	}

	fmt.Println("Successfully fetched all urls")
	fmt.Println(time.Since(start))

	/*
	   $ go run errgroup.go
	   http://www.google.com/ 110.904497ms
	   http://www.youtube.com/ 295.687404ms
	   http://www.golang.org/ 1.25146058s
	   http://www.somestupidname.com/ 14.157130046s
	   Successfully fetched all urls
	   14.157168909s
	*/
}
