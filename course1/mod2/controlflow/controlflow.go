package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	x, _ := strconv.ParseInt(os.Args[1], 10, 32)
	if x > 10 {
		fmt.Println("x is greater than 10")
	} else {
		fmt.Println("x is less than or equal to 10")
	}

	//basic for loop
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	// while loop
	i := 0
	for i < 5 {
		fmt.Println(i)
		i++
	}

	for i := range 5 {
		fmt.Printf("Hello - %d\n", i)
	}

	doubled := Map([]int{1, 2, 3, 4, 5}, func(x int) float64 { return math.Pow(float64(x), 2) })
	fmt.Println(doubled)

	// switch statements with tag
	switch x {
	case 1:
		fmt.Println("x is 1")
	case 2:
		fmt.Println("x is 2")
	default:
		fmt.Println("x is not 1 or 2")
	}

	// switch statements without tag
	switch {
	case x < 10:
		fmt.Println("x is less than 10")
	case x == 10:
		fmt.Println("x is equal to 10")
	case x > 10:
		fmt.Println("x is greater than 10")
	default:
		fmt.Println("x is not a number")
	}

	// break and continue
	for i := 0; i < 5; i++ {
		if i == 3 {
			continue
		}
		fmt.Println(i)
		if i == 4 {
			break
		}
	}

	// scan for user input
	var userNum int
	fmt.Println("Enter a number: ")
	fmt.Scan(&userNum)
	fmt.Printf("You entered: %d\n", userNum)
}

func Map[T, V any](ts []T, fn func(T) V) []V {
	result := make([]V, len(ts))
	for i, t := range ts {
		result[i] = fn(t)
	}
	return result
}
