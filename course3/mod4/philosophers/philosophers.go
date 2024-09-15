package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type ChopStick struct {
	mu     *sync.Mutex
	number int
}

type Philosopher struct {
	number          int
	leftCS, rightCS *ChopStick
	turnsLeft       int
}

func newPhilosopher(n int, l, r *ChopStick) *Philosopher {
	return &Philosopher{n, l, r, 3}
}

func (p *Philosopher) Eat(host chan bool, wg *sync.WaitGroup, times []time.Duration) {
	defer wg.Done()
	for p.turnsLeft > 0 {
		p.pickUpChopsticks()
		select {
		case <-host:
			fmt.Printf("starting to eat %d\n", p.number)
			p.turnsLeft--
			host <- true
			fmt.Printf("finishing eating %d\n", p.number)
		default:
			host <- true
		}
		p.releaseChopsticks()
	}
}

func (p *Philosopher) pickUpChopsticks() {
	rand := rand.Int() % 2
	if rand == 0 {
		p.leftCS.mu.Lock()
		p.rightCS.mu.Lock()
	} else {
		p.rightCS.mu.Lock()
		p.leftCS.mu.Lock()
	}
}

func (p *Philosopher) releaseChopsticks() {
	rand := rand.Int() % 2
	if rand == 0 {
		p.leftCS.mu.Unlock()
		p.rightCS.mu.Unlock()
	} else {
		p.rightCS.mu.Unlock()
		p.leftCS.mu.Unlock()
	}
}

func main() {

	cSticks := make([]*ChopStick, 5)
	for i := 0; i < 5; i++ {
		cSticks[i] = &ChopStick{&sync.Mutex{}, i + 1}
	}

	philos := make([]*Philosopher, 5)
	for i := 0; i < 5; i++ {
		philos[i] = newPhilosopher(i+1, cSticks[i], cSticks[(i+1)%5])
	}
	host := make(chan bool, 2)
	host <- true
	host <- true
	wg := sync.WaitGroup{}
	times := make([]time.Duration, 5)
	for _, p := range philos {
		wg.Add(1)
		go p.Eat(host, &wg, times)
	}
	wg.Wait()
}
