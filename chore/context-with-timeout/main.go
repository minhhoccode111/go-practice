package main

import (
	"context"
	"fmt"
	"time"
)

func process(ctx context.Context) {
	select {
	case <-time.After(time.Second * 3): // simulate a process that takes 3 seconds
		fmt.Println("process completed")
	case <-ctx.Done(): // context cancellation or timeout
		fmt.Println("process cancelled: ", ctx.Err())
	}

}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()

	process(ctx) // process cancelled:  context deadline exceeded
}
