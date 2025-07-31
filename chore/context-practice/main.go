package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	contextValue()

	contextWithCancel()

	contextWithTimeout()

	contextWithDeadline()
}

func contextWithDeadline() {
	base := context.Background()
	ctx, cancel := context.WithDeadline(base, time.Now().Add(time.Millisecond*1))
	defer cancel()
	select {
	case <-time.After(time.Second * 1):
		fmt.Println("time.After() is called, which mean 1 second passed")
	case <-ctx.Done():
		fmt.Println("ctx.Done() is called, which mean 1 millisecond passed")
	}
}

func contextWithTimeout() {
	base := context.Background()
	// timeout in 1 second
	ctx, cancel := context.WithTimeout(base, 1*time.Second)
	defer cancel()
	select {
	case <-time.After(time.Second * 2):
		fmt.Println("time.After() is called, which mean 2 seconds passed")
	case <-ctx.Done():
		fmt.Println("ctx.Done() is called, which mean 1 second passed")
	}
}

func contextWithCancel() {
	base := context.Background()
	_, cancel := context.WithCancel(base)
	defer func() {
		fmt.Println("Go is so great bro")
		// call cancel() function to cancel the context when this function returns
		defer func() { defer cancel() }()
	}()
}

func contextValue() {
	type key string
	type val string
	base := context.Background() // base context at the edge of your program
	ctx := context.WithValue(base, key("username"), val("minhhoccode111"))
	// reuse variable 'ctx'
	ctx = context.WithValue(ctx, key("password"), val("super_secret_password"))
	// new variable derived from 'ctx'
	newVarCtx := context.WithValue(ctx, key("age"), 18)
	fmt.Println(newVarCtx.Value(key("username")))
	fmt.Println(newVarCtx.Value(key("password")))
	fmt.Println(newVarCtx.Value(key("age")))
}
