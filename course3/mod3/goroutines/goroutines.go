package main

import (
	"fmt"
	"sync"
)

func main() {
	// when main finishes, it will kill all goroutines
	go func() {
		fmt.Println("Hello from goroutine")
	}()
	fmt.Println("Hello from main")
	// in this case, the goroutine will not have time to execute
	// because the main function will finish first

	// there are many ways to fix this, one of them is using a wait group
	// to wait for the goroutine to finish
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("Hello from goroutine wait group")
	}()
	wg.Wait()
	fmt.Println("Hello from main waitgroup")

	// another way is to use a channel to block the main function
	// until the goroutine finishes
	done := make(chan bool)
	go func() {
		fmt.Println("Hello from goroutine channel")
		done <- true
		// this will close the channel. it is not necessary in this case
		// because the channel is unbuffered and the goroutine will block
		// until the value is consumed. however it's a good practice to close
		// the channel when it's no longer needed to notify consumers that they
		// should stop waiting for values
		close(done)
	}()
	// this is consuming the value from the channel
	// it will block until the goroutine sends a value
	<-done
	fmt.Println("Hello from main channel")
}
