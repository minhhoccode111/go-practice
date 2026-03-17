package main

import (
	"fmt"
	"sync"
	"time"
)

type chopstick struct {
	sync.Mutex
}

type philo struct {
	left  *chopstick
	right *chopstick
}

func (p philo) eat(
	request chan struct{},
	finish chan struct{},
	wg *sync.WaitGroup,
	n int,
) {
	defer wg.Done()

	for i := range 3 {
		request <- struct{}{}
		p.left.Lock()
		p.right.Lock()

		fmt.Printf("%d starting to eat, the %d time\n", n+1, i+1)

		time.Sleep(1 * time.Second)

		fmt.Printf("%d finishing eating, the %d time\n", n+1, i+1)

		p.left.Unlock()
		p.right.Unlock()
		finish <- struct{}{}
	}
}

const n = 5

func main() {
	// initialization
	chops := make([]*chopstick, n)
	for i := range n {
		chops[i] = new(chopstick)
	}
	philos := make([]*philo, n)
	for i := range n {
		philos[i] = &philo{
			left:  chops[i],
			right: chops[(i+1)%n],
		}
	}

	var wg sync.WaitGroup
	request := make(chan struct{})
	finish := make(chan struct{})

	go func() {
		eating := 2
		for {
			select {
			case <-request:
				if eating >= 2 {
					<-finish
					eating--
				}
				eating++
			case <-finish:
				eating--
			}
		}
	}()

	for i, v := range philos {
		wg.Add(1)
		go v.eat(request, finish, &wg, i)
	}

	wg.Wait()
}
