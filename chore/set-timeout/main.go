package main

import (
	"fmt"
	"time"
)

func setTimeout(f func(), t time.Duration) {
	go func() {
		<-time.After(t * time.Millisecond)
		f()
	}()
}

func main() {
	fmt.Println(`setTimeout(func() { fmt.Println("Hello") }, 3000);`)
	setTimeout(func() { fmt.Println("Hello") }, 3000)
}
