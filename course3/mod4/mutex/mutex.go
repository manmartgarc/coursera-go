package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var i int = 0

func inc(wg *sync.WaitGroup) {
	i = i + 1
	wg.Done()
}

func incMutex(mu *sync.Mutex, wg *sync.WaitGroup) {
	mu.Lock()
	i = i + 1
	mu.Unlock()
	wg.Done()
}

func incAtomic(i *atomic.Int32, wg *sync.WaitGroup) {
	i.Add(1)
	wg.Done()

}

func main() {
	// even though this looks benign, it is not thread safe. This is because
	// the interleaving of operations happens not at the Go statement level
	// but at the machine instruction level. This means that the i = i + 1
	// operation can be broken down into multiple machine instructions, which
	// means that there can be many interleavings between reads, increments and
	// writes to the i variable.
	wg := sync.WaitGroup{}
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go inc(&wg)
	}
	wg.Wait()
	fmt.Println(i)

	// we can use a mutex to protect the i variable
	i = 0
	wg = sync.WaitGroup{}
	mu := sync.Mutex{}
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go incMutex(&mu, &wg)
	}
	wg.Wait()
	fmt.Println(i)

	// we can also use the atomic package to protect the i variable
	wg = sync.WaitGroup{}
	iAtomic := atomic.Int32{}
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go incAtomic(&iAtomic, &wg)
	}
	wg.Wait()
	fmt.Println(iAtomic.Load())
}
