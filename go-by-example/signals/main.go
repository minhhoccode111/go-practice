package main

import (
	"context"
	"fmt"
	"os/signal"
	"syscall"
)

func main() {
	ctx, stop := signal.NotifyContext(
		context.Background(),
		syscall.SIGINT,
		syscall.SIGTERM,
	)
	defer stop()
	fmt.Println("awaiting signal")
	<-ctx.Done()

	fmt.Println()
	fmt.Println(context.Cause(ctx))
	fmt.Println("exiting")
}
