package stopgracefully

import (
	"fmt"
	"time"
)

func stop(quit <-chan bool, ch <-chan int) {
	for {
		select {
		case <-quit:
			fmt.Println("stopping after received quit signal")
			return
		case <-time.After(5 * time.Second):
			fmt.Println("stopping after 5 seconds")
			return
		case v := <-ch:
			fmt.Println("received:", v)
		}
	}
}
