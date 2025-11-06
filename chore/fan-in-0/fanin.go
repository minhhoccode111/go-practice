package main

import (
	"fmt"
	"math/rand/v2"
	"sync"
	"time"
)

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	ch3 := make(chan int)
	ch4 := make(chan int)

	go func() {
		randomSend(ch1)
		randomSend(ch2)
		randomSend(ch3)
		randomSend(ch4)
	}()

	out := fanin(ch1, ch2, ch3, ch4)
	c := 0
	for v := range out {
		fmt.Println("value:", v)
		c++
		fmt.Println("count:", c)
	}
}

func fanin(chs ...<-chan int) chan int {
	out := make(chan int)
	var wg sync.WaitGroup
	wg.Add(len(chs))
	for _, ch := range chs {
		go func(c *<-chan int) {
			for v := range ch {
				out <- v
			}
			wg.Done()
		}(&ch)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

func randomSend(c chan int) {
	for range 10 {
		time.Sleep(time.Duration(rand.IntN(5)) * time.Second)
		c <- rand.IntN(10)
	}
	close(c)
}
