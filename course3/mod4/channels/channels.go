package main

import (
	"fmt"
	"math/rand"
)

func producer(rInts chan int) {
	for i := 0; i < 5; i++ {
		rInts <- rand.Intn(100)
	}
	close(rInts)
}

func consumer(rInts chan int, done chan bool) {
	for v := range rInts {
		fmt.Printf("%d mod 7: %d\n", v, v%7)
	}
	done <- true
}

// can use select to consume from multiple channels
func consumeTwo(done chan bool, rInts1 chan int, rInts2 chan int) {
	for {
		select {
		case v1, ok := <-rInts1:
			if ok {
				fmt.Printf("from rInts1: %d mod 7: %d\n", v1, v1%7)
			} else {
				rInts1 = nil
			}
		case v2, ok := <-rInts2:
			if ok {
				fmt.Printf("from rInts2: %d mod 7: %d\n", v2, v2%7)
			} else {
				rInts2 = nil
			}
		}

		if rInts1 == nil && rInts2 == nil {
			break
		}
	}
	done <- true
}

// can also use select with an abort channel
func consumeAbort(inChan chan int, abort chan bool, done chan bool) {
	for {
		select {
		case v, ok := <-inChan:
			if ok {
				fmt.Printf("from inChan: %d mod 7: %d\n", v, v%7)
				if v%7 < 3 {
					abort <- true
				}
			} else {
				inChan = nil
			}
		case <-abort:
			return
		}

		if inChan == nil {
			break
		}
	}
	done <- true
}

func main() {
	rInts := make(chan int)
	done := make(chan bool)
	go producer(rInts)
	go consumer(rInts, done)
	<-done

	// can use select to consume from multiple channels
	rInts1 := make(chan int)
	rInts2 := make(chan int)
	rInts3 := make(chan int)
	done = make(chan bool)
	go producer(rInts1)
	go producer(rInts2)
	go producer(rInts3)
	go consumeTwo(done, rInts1, rInts2)
	<-done

	// can also use select to synchronize input and output
	inChan := make(chan int)
	abort := make(chan bool)
	go producer(inChan)
	done = make(chan bool)
	go consumeAbort(inChan, abort, done)
	select {
	case <-done:
		fmt.Println("Done")
	case <-abort:
		fmt.Println("Aborted")
	}
}
