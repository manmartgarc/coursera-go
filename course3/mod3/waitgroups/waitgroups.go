package main

import (
	"fmt"
	"sync"
)

func foo(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Hello from goroutine wait group")
}

func main() {
	wg := sync.WaitGroup{}
	wg.Add(1)
	go foo(&wg)
	wg.Wait()
	fmt.Println("Hello from main waitgroup")
}
