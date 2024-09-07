package main

import "fmt"

func foo(arr [3]int) int {
	return arr[0]
}

func bar(arr *[3]int) int {
	return arr[0]
}

func sliceFoo(s []int) int {
	return s[0]
}

func main() {
	a := [3]int{1, 2, 3}
	// this will copy the array to the function
	fmt.Println(foo(a))
	// this will pass the memory address of the array to the function
	fmt.Println(bar(&a))
	// this will pass the memory address of the array to the function
	sli := a[:]
	fmt.Println(sliceFoo(sli))
}
