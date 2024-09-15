package main

import (
	"fmt"
	"os"
	"strconv"
	"sync"
)

func p1(x *string) {
	*x = "p1"
}

func p2(x *string) {
	*x = "p2"
}

func runSim() string {
	x := ""
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		p1(&x)
		wg.Done()
	}()
	go func() {
		p2(&x)
		wg.Done()
	}()
	wg.Wait()
	return x
}

// The main function runs the simulation n times and prints the number of times p1 and p2 won.
// Since a race condition exists when the outcome depends on a non-deterministic order of execution,
// the number of times p1 and p2 win will vary.
// To run the simulation, use the following command:
//
//	go run racecondition.go 1000
func main() {
	var np1, np2 int
	args := os.Args[1:]
	if len(args) > 0 {
		n, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println("Error: ", err)
			return
		}
		for i := 0; i < n; i++ {
			if runSim() == "p1" {
				np1++
			} else {
				np2++
			}
		}
		fmt.Printf("In %d runs, p1 won %d times and p2 won %d times\n", n, np1, np2)
	}
}
