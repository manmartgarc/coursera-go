package main

import "fmt"

func Max(vals ...int) int {
	maxV := vals[0]
	for _, v := range vals {
		if v > maxV {
			maxV = v
		}
	}
	return maxV
}

func deferred() {
	defer fmt.Println("I'm done!")
	fmt.Println("I'm doing a bunch of work!")
}

func notHowYouThink() {
	i := 1
	// this gets evaluated immediately!
	defer fmt.Println(i + 1)
	i++
	fmt.Println("Hello!")
}

func main() {
	fmt.Println(Max(1, 2, 3, 4, 5))
	// or even better with a slice
	vals := []int{1, 2, 3, 4, 5}
	fmt.Println(Max(vals...))
	deferred()
	notHowYouThink()
}
