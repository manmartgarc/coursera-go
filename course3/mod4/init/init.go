package main

import (
	"fmt"
	"sync"
)

var on sync.Once

func setup() {
	fmt.Println("Set up")
}

func doStuff(wg *sync.WaitGroup) {
	// this will only run once and the other goroutines will wait for it to finish
	on.Do(setup)
	fmt.Println("Doing stuff")
	wg.Done()
}

func main() {
	wg := sync.WaitGroup{}
	wg.Add(2)
	go doStuff(&wg)
	go doStuff(&wg)
	wg.Wait()
}
