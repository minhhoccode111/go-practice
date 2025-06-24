package main

import (
	"fmt"
	"sync"
	"time"
)

func process(i int, wg *sync.WaitGroup) {
	fmt.Println("start process: ", i)
	time.Sleep(2 * time.Second)
	fmt.Println("finish process: ", i)
	wg.Done()
	wg.Done()
}

func main() {
	n := 3
	var wg sync.WaitGroup
	for i := range n {
		go process(i, &wg)
		wg.Add(2)
	}
	wg.Wait()
	fmt.Println("all gorountine has finished")
}
