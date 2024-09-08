package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func GenDisplaceFn(a, v0, s0 float64) func(float64) float64 {
	fn := func(t float64) float64 {
		return 0.5*a*math.Pow(t, 2) + v0*t + s0
	}
	return fn
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var a, v0, s0, t float64
	var err error
	fmt.Print("Enter a0, v0 and s0 as space-separated floats: ")
	scanner.Scan()
	_, err = fmt.Sscanf(scanner.Text(), "%f %f %f ", &a, &v0, &s0)
	if err != nil {
		fmt.Println("Invalid input")
		return
	}
	fmt.Printf("Initial values: a = %.2f, v0 = %.2f, s0 = %.2f\n", a, v0, s0)
	fn := GenDisplaceFn(a, v0, s0)

	for {
		fmt.Print("\nEnter time (in seconds) or type 'x' to exit: ")
		scanner.Scan()
		if scanner.Text() == "x" {
			break
		}
		t, err = strconv.ParseFloat(scanner.Text(), 64)
		if err != nil {
			fmt.Println("Invalid input")
			continue
		}
		fmt.Printf("Displacement after %.2f seconds: %.2f\n", t, fn(t))
	}
}
