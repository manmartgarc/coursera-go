package main

import "fmt"

// this is a global variable and allocated in the heap
var x = 0

func f() {
	// this is a local variable and allocated in the stack
	var y = 10
	fmt.Printf("(%d, %d)\n", x, y)
	y += 1
	x += 1
	// y is deallocated here since the scope of y is the function f.
	// it's not deallocated by the garbage collector. the garbage collector
	// deallocates stuff from the heap only
}

func main() {
	for range 10 {
		f()
	}
}
