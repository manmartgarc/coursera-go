package main

import "fmt"

func main() {
	// arrays are of fixed size!
	// elements are initialized to their zero value
	var a [3]int
	for i, old := range a {
		a[i] = i
		fmt.Printf("Element %d: %d -> %d\n", i, old, a[i])
	}

	// we can also define array literals
	b := [3]int{1, 2, 3}
	for i := range b {
		fmt.Printf("Element %d: %d\n", i, b[i])
	}

	c := [...]int{1, 2, 3} // this is the same as above
	for i := range c {
		fmt.Printf("Element %d: %d\n", i, c[i])
	}
}
