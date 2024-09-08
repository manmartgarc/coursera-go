package main

import "fmt"

var funcVar func(int) int // Declaring a variable of function type
func incFn(x int) int {
	return x + 1
}

// home-grown map function
func mapFn(f func(int) int, list []int) []int {
	result := make([]int, len(list))
	for i, v := range list {
		result[i] = f(v)
	}
	return result
}

func main() {
	funcVar = incFn // Assigning a function to a variable
	fmt.Println(funcVar(1))

	square := func(x int) int { return x * x } // Anonymous function

	list := []int{1, 2, 3}
	fmt.Println(mapFn(incFn, list))
	fmt.Println(mapFn(square, list))
	fmt.Println(mapFn(func(x int) int { return x % 2 }, list))
}
