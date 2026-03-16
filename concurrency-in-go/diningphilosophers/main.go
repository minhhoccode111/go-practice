package main

import (
	"fmt"
	"sync"
	"time"
)

type ChopS struct {
	sync.Mutex
}

type Philo struct {
	left  *ChopS
	right *ChopS
}

func (p Philo) eat(i int) {
	// each philo each 3 times
	for range 3 {
		p.left.Lock()
		p.right.Lock()

		fmt.Println("starting to eat", i)
		fmt.Println("finishing eating", i)

		p.left.Unlock()
		p.right.Unlock()
	}
}

func main() {
	// initialization
	chops := make([]*ChopS, 5)
	for i := range 5 {
		chops[i] = new(ChopS)
	}

	philos := make([]*Philo, 5)
	for i := range 5 {
		philos[i] = &Philo{
			left:  chops[i],
			right: chops[(i+1)%5],
		}
	}

	for i := range 5 {
		go philos[i].eat(i)
	}

	time.Sleep(100 * time.Millisecond)
}
