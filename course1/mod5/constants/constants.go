package main

import (
	"fmt"
	"slices"
)

func main() {
	const x = 1.3
	fmt.Println(x)

	// assigning multiple as one
	const (
		y = 4
		z = "Hi"
	)
	fmt.Println(y, z)

	// iota; this is like an enum
	type Grades int
	const (
		A Grades = iota
		B
		C
		D
		F
	)
	fmt.Println(A, B, C)
	grades := []Grades{C, A, B, F, A, D}
	slices.Sort(grades)
	fmt.Println(grades)
}
