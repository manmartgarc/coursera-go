package main

import (
	"fmt"
)

func main() {
	// slices are like dynamic arrays, they can grow and shrink
	// they are a "window" to an underlying array

	// There are three parts to a slice:
	// 1. A pointer to the first element of the array
	// 2. The length of the slice
	// 3. The capacity of the slice

	// The capacity of a slice is the number of elements in the underlying array,
	// starting from the index of the first element in the slice.
	// Mathematically, the capacity of a slice s is:
	// capacity(s) = len(array) - index(s)

	// we start with the underlying array
	arr := [...]string{"a", "b", "c", "d", "e", "f", "g"}
	s1 := arr[1:3] // s1 is a slice of arr from index 1 to 3 (not inclusive)
	s2 := arr[2:5] // s2 is a slice of arr from index 2 to 5 (not inclusive)

	// let's check the underling memory addresses are the same
	fmt.Printf("c in arr: %p | c in s1: %p | c in s2: %p\n", &arr[1], &s1[1], &s2[0])

	// let's check the lengths and capacities
	fmt.Printf("s1: len=%d cap=%d\n", len(s1), cap(s1))
	fmt.Printf("s2: len=%d cap=%d\n", len(s2), cap(s2))

	// let's change the element "c" in arr and see if the slices are affected
	arr[2] = "z"
	fmt.Printf("s1: %v | s2: %v\n", s1, s2)

	// we can also initialize slices with literals
	// this will create a new underlying array and a new slice
	// the capacity of the slice will be the same as the length of the array
	sli := []int{1, 2, 3, 4, 5}
	fmt.Printf("sli: %v\n", sli)
	fmt.Printf("sli: len=%d cap=%d\n", len(sli), cap(sli))

	// a third way to create a slice is using the make function
	// we first create a slice with a length of 10 and a capacity of 10
	sliOne := make([]int, 10)
	fmt.Printf("sliOne: %v\n", sliOne)
	// we can also specify the capacity of the slice. This will create a slice
	// with a length of 10 and a capacity of 15
	sliTwo := make([]int, 10, 15)
	fmt.Printf("sliTwo: %v\n", sliTwo)

	// we can append to slices, even beyond their capacity
	// there is a performance penalty when the slice has to be resized
	// the new capacity will be double the old capacity
	sliThree := make([]int, 0)
	fmt.Printf("sliThree: len=%d cap=%d\n", len(sliThree), cap(sliThree))
	fmt.Println("appending 1, 2, 3, 4, 5, 6, 7, 8, 9, 10")
	sliThree = append(sliThree, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	fmt.Printf("sliThree: len=%d cap=%d\n", len(sliThree), cap(sliThree))
}
