package main

import (
	"fmt"
	"math"
)

// this is a way to parametrize functions
// but why would you do this if you can just pass the arguments to a single function?
func MakeDistOrigin(o_x, o_y float64) func(float64, float64) float64 {
	fn := func(x, y float64) float64 {
		return math.Sqrt(math.Pow(x-o_x, 2) + math.Pow(y-o_y, 2))
	}
	return fn
}

func main() {
	dist1 := MakeDistOrigin(0, 0)
	dist2 := MakeDistOrigin(2, 2)
	fmt.Printf("%.2f\n", dist1(8, 8))
	fmt.Printf("%.2f\n", dist2(8, 8))
}
