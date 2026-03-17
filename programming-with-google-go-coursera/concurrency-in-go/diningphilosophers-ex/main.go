package main

import (
	"fmt"
	"sync"
)

type ChopS struct {
	sync.Mutex
}

type Philo struct {
	leftCS  *ChopS
	rightCS *ChopS
}

func (p Philo) eat() {
	for {
		p.leftCS.Lock()
		p.rightCS.Lock()

		fmt.Println("eating")

		p.leftCS.Unlock()
		p.rightCS.Unlock()
	}
}

func main() {
	CSticks := make([]*ChopS, 5)
	for i := range 5 {
		CSticks[i] = new(ChopS)
	}

	philos := make([]*Philo, 5)
	for i := range 5 {
		philos[i] = &Philo{
			CSticks[i],
			CSticks[(i+1)%5],
		}
	}

	for i := range 5 {
		go philos[i].eat()
	}

	// but careful with this implementation, because all of them can grab the left chopstick at the same time, which can cause a deadlock

	// the other solution is the the philo always pick the lower index chopstick first, which makes the philos[4] pick the chopsticks[0] first, and can wait for the philos[0] to finish first if it's in use, and prevent deadlock. But this solution can make philos[4] starve.
}
