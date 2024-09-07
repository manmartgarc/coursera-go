package main

import "fmt"

// call by value does not change the original value
// the value passed to the function is a copy of the original value and
// assigned to a new memory location. Any changes made to the value inside
// the function does not affect the original value.

// this is good for immutability and data integrity but not good for
// performance since it requires more memory to store the copy of the value.
func foo(y int) {
	y += 1
}

// call by reference changes the original value since the memory address
// of the original value is passed to the function. Any changes made to the
// value inside the function affects the original value.
// call by reference or call by pointer is achieved by passing the memory
// address of the value to the function.

// this is good for performance since it does not require more memory to store
// a copy of the value but bad for immutability and data integrity.
func bar(y *int) {
	*y += 1
}

func check(x int) {
	if x == 2 {
		fmt.Println("x is 2")
	} else {
		fmt.Println("x is not 2")
	}
}

func main() {
	x := 2
	foo(x)
	check(x)
	bar(&x)
	check(x)
}
