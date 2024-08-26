package main

import "fmt"

var x = 4

func f() {
	fmt.Printf("x = %d\n", x)
}

func g() {
	fmt.Printf("x = %d\n", x)
}

func main() {
	y := 5
	f()
	g()
	// this is an anonymous function, i.e. lambda
	h := func() { fmt.Printf("y = %d\n", y) }
	h()
}
