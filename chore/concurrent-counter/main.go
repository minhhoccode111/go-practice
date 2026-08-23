package main

import (
	"fmt"
	"sync"
)

// // single-threaded approach, fastest so far :)
// func main() {
// 	count := 0
// 	for range 1_000_000 {
// 		for range 1_000 {
// 			count++
// 		}
// 	}
// 	fmt.Print(count)
// 	/*
// 	   $ time go run .
// 	   1000000000
// 	   real    0m0.365s
// 	   user    0m0.378s
// 	   sys     0m0.053s
// 	*/
// }

// // nothing approach
// func main() {
// 	count := 0
// 	for range 1_000_000 {
// 		go func() {
// 			for range 1_000 {
// 				count++
// 			}
// 		}()
// 	}
// 	time.Sleep(5 * time.Second)
// 	fmt.Print(count)
// 	/*
// 	   $ time go run .
// 	   113631468
// 	   real    0m5.938s
// 	   user    0m10.561s
// 	   sys     0m0.089s
// 	*/
// }

// // with mutex approach to prevent race condition, but still no syncing mechanism
// func main() {
// 	count := 0
// 	var mu sync.Mutex
// 	for range 1_000_000 {
// 		go func() {
// 			mu.Lock()
// 			defer mu.Unlock()
// 			for range 1_000 {
// 				count++
// 			}
// 		}()
// 	}
// 	time.Sleep(5 * time.Second)
// 	fmt.Print(count)
// 	/*
// 	   $ time go run .
// 	   1000000000
// 	   real    0m6.105s
// 	   user    0m5.804s
// 	   sys     0m1.199s
// 	*/
// }

// // with mutex approach, and sync mechanism
// func main() {
// 	count := 0
// 	var mu sync.Mutex
// 	var wg sync.WaitGroup
// 	for range 1_000_000 {
// 		wg.Add(1)
// 		go func() {
// 			mu.Lock()
// 			defer mu.Unlock()
// 			defer wg.Done()
// 			for range 1_000 {
// 				count++
// 			}
// 		}()
// 	}
// 	wg.Wait()
// 	fmt.Print(count)
// 	/*
// 	   $ time go run .
// 	   1000000000
// 	   real    0m2.180s
// 	   user    0m6.106s
// 	   sys     0m1.123s
// 	*/
// }

// // with atomic counter, which is recommended, but perform the worst by far
// func main() {
// 	var count atomic.Int64
// 	var wg sync.WaitGroup
// 	for range 1_000_000 {
// 		wg.Go(func() {
// 			for range 1_000 {
// 				count.Add(1)
// 			}
// 		})
// 	}
// 	wg.Wait()
// 	fmt.Print(count.Load())
// 	/*
// 	   $ time go run .
// 	   1000000000
// 	   real    0m15.464s
// 	   user    3m0.635s
// 	   sys     0m0.258s
// 	*/
// }

func main() {
	count := 0
	ch := make(chan struct{}, 1_000_000_000)
	done := make(chan struct{})
	var wg sync.WaitGroup
	go func() {
		defer close(done)
		for range ch {
			count++
		}
	}()
	for range 1_000_000 {
		wg.Go(func() {
			for range 1_000 {
				ch <- struct{}{}
			}
		})
	}
	go func() {
		defer close(ch)
		wg.Wait()
	}()
	<-done
	fmt.Print(count)
	// this approach only works because the channel is buffered, it would be worse with an unbuffered channel
	/*
	   $ time go run .
	   1000000000
	   real    0m51.737s
	   user    1m21.256s
	   sys     0m9.526s
	*/
}
