package main

import "fmt"

func foo() *int {
	x := 1
	// even though x is a local variable in the function scope, it's not
	// allocated on the stack, but on the heap. This is because some other
	// process might need to access the value of x after the function has
	// returned. This is an example of how the compiler determines whether
	// to allocate memory on the stack or the heap. If it's allocated on the
	// heap, the garbage collector will free the memory when it's no longer
	// needed.
	return &x
}

func main() {
	y := foo()
	fmt.Printf("Value of y is %d\n", *y)
}
