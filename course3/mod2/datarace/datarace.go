package main

import (
	"fmt"
	"sync"
)

var counter = 10000

func increment(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 100; i++ {
		counter++
	}
}

// if we run this program multiple times, we will get different results
// using the -race flag will help us identify the data:
//
//	go run -race datarace.go
func main() {
	wg := sync.WaitGroup{}
	wg.Add(100)

	for i := 0; i < 100; i++ {
		go increment(&wg)
	}
	wg.Wait()
	fmt.Printf("Final counter: %.d\n", counter)

	// Since the final counter > 10000, we can conclude that a data
	// race occurred. This can be easily fixed by using either a mutex lock
	// or atomic integer type for the counter.
}
