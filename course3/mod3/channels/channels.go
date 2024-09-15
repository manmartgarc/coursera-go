package main

import "fmt"

func sum(a, b int, c chan int) {
	c <- a + b
}

func main() {
	// by default a channel is unbuffered; meaning it has a capacity of 0.
	// this means that it will block the receiver until the sender has put
	// a value in the channel and vice versa. The sender will be blocked until
	// a receiver has consumed a value from the channel.
	c := make(chan int)
	go sum(1, 2, c)
	go sum(3, 4, c)
	x, y := <-c, <-c
	fmt.Printf("x=%d, y=%d, x+y=%d\n", x, y, x+y)

	// to create a buffered channel, you can specify the capacity of the channel
	cb := make(chan int, 2)
	// since the channel has a capacity of 2, the sender can send up to 2
	// values without blocking. The receiver will block only when the channel
	// is full; preventing the sender from sending more values.
	go sum(5, 6, cb)
	go sum(7, 8, cb)
	z, w := <-cb, <-cb
	fmt.Printf("z=%d, w=%d, z+w=%d\n", z, w, z+w)
	// they key thing about buffered channels is that they allow producers
	// and consumers to operate at different speeds without blocking each other.

	// either buffered or unbuffered, channels always block the sender when
	// the channel is full and the receiver when the channel is empty.
	// the key idea behind channels is not to communicate by sharing memory,
	// i.e. a wait group, but to share memory by communicating. This is a
	// powerful concept that allows goroutines to communicate with each other
	// without the need for explicit locks, condition variables, atomic operations or
	// even non-blocking synchronization primitives.
}
