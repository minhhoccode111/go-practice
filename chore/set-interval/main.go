package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println(
		`id := setInterval(func() { fmt.Println("tick") }, 1000)
<-time.After(5 * time.Second)
clearInterval(id)
// wait but no output this time
<-time.After(2 * time.Second)`,
	)
	id := setInterval(func() { fmt.Println("tick") }, 1000*time.Millisecond)
	<-time.After(5 * time.Second)
	clearInterval(id)
	// wait but no output this time
	<-time.After(3 * time.Second)
}

func clearInterval(id chan struct{}) { close(id) }

func setInterval(f func(), d time.Duration) chan struct{} {
	stop := make(chan struct{})
	go func() {
		ticker := time.NewTicker(d)
		defer ticker.Stop()
		for {
			select {
			case <-ticker.C:
				f()
			case <-stop:
				return
			}
		}
	}()
	return stop
}
